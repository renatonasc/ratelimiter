package webserver

import (
	"log"
	"log/slog"
	"net/http"
	"os"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/httplog/v2"

	customMiddleware "renatonasc/ratelimit/internal/middleware"
)

type WebServer struct {
	Router            chi.Router
	Handlers          []WebServerHandler
	ProtectedHandlers []WebServerHandler
	WebServerPort     string
}

type WebServerHandler struct {
	Method  string
	Path    string
	Handler http.HandlerFunc
}

func NewWebServer(serverPort string) *WebServer {
	return &WebServer{
		Router:            chi.NewRouter(),
		Handlers:          make([]WebServerHandler, 0),
		ProtectedHandlers: make([]WebServerHandler, 0),
		WebServerPort:     serverPort,
	}
}

func (s *WebServer) AddHandler(method string, path string, handler http.HandlerFunc) {
	s.Handlers = append(s.Handlers, WebServerHandler{
		Method:  method,
		Path:    path,
		Handler: handler,
	})
}

func (s *WebServer) AddProtectedHandler(method string, path string, handler http.HandlerFunc) {
	s.ProtectedHandlers = append(s.ProtectedHandlers, WebServerHandler{
		Method:  method,
		Path:    path,
		Handler: handler,
	})
}

// loop through the handlers and add them to the router
// register middeleware logger
// start the server
func (s *WebServer) Start(rl *customMiddleware.RateLimit) {
	// Open the log file
	logFile, err := os.OpenFile("log.txt", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatal(err)
	}
	defer logFile.Close()

	// Logger
	logger := httplog.NewLogger("httplog", httplog.Options{
		JSON:             true,
		LogLevel:         slog.LevelDebug,
		Concise:          true,
		RequestHeaders:   true,
		ResponseHeaders:  true,
		MessageFieldName: "message",
		// TimeFieldFormat: time.RFC850,
		Tags: map[string]string{
			"version": "v1.0-81aa4244d9fc8076a",
			"env":     "dev",
		},
		QuietDownRoutes: []string{
			"/",
			"/ping",
		},
		QuietDownPeriod: 10 * time.Second,
		// SourceFieldName: "source",
		Writer: logFile,
	})

	s.Router.Use(rl.RateLimitMiddleware)
	s.Router.Use(middleware.Logger)
	s.Router.Use(httplog.RequestLogger(logger))
	s.Router.Use(middleware.Recoverer)
	s.Router.Use(middleware.Heartbeat("/ping"))
	//OpenAPI
	for _, handler := range s.Handlers {
		log.Println("Adding handler", handler.Path)
		s.Router.Method(handler.Method, handler.Path, handler.Handler)
	}
	// Protected routes
	s.Router.Group(func(r chi.Router) {
		for _, handler := range s.ProtectedHandlers {
			r.Method(handler.Method, handler.Path, handler.Handler)
		}
	})

	http.ListenAndServe(s.WebServerPort, s.Router)
}
