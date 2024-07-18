package main

import (
	"fmt"
	"log"
	"net/http"
	"renatonasc/ratelimit/configs"
	"renatonasc/ratelimit/internal/infra/database"

	"renatonasc/ratelimit/internal/infra/webserver"
	customMiddleware "renatonasc/ratelimit/internal/middleware"
)

func main() {

	configs, err := configs.LoadConfig(".")
	if err != nil {
		log.Fatalf("Error loading configs: %v", err)
		panic(err)
	}

	rdb := database.NewRedis(configs.DBHost, configs.DBPort)
	rl := customMiddleware.NewRateLimit(configs.MaxRequestToken, configs.MaxRequestIp, configs.BlockTimeIp, configs.BlockTimeToken, rdb)
	webserver := webserver.NewWebServer(configs.WebServerPort)

	webserver.AddHandler("GET", "/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello World"))
	})

	fmt.Println("Starting web server on port", configs.WebServerPort)
	webserver.Start(rl)

}
