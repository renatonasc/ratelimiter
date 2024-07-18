package middleware

import (
	"log"
	"net/http"
	"renatonasc/ratelimit/internal/usecase"

	"github.com/redis/go-redis/v9"
)

type RateLimit struct {
	MaxRequestToken int
	MaxRequestIp    int
	BlockTimeIp     int
	BlockTimeToken  int
	rdb             *redis.Client
}

func NewRateLimit(maxRequestToken, maxRequestIp, blockTimeIp int, blockTimeToken int, rdb *redis.Client) *RateLimit {
	return &RateLimit{
		MaxRequestToken: maxRequestToken,
		MaxRequestIp:    maxRequestIp,
		BlockTimeIp:     blockTimeIp,
		BlockTimeToken:  blockTimeToken,
		rdb:             rdb,
	}
}

func (rl *RateLimit) RateLimitMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println("RateLimitMiddleware")
		ip := r.RemoteAddr
		requestToken := r.Header.Get("API_KEY")
		log.Println("RateLimitMiddleware IP -> " + ip)
		log.Println("RateLimitMiddleware Token -> " + requestToken)

		var dto usecase.RateLimitInputDTO
		var ratelimit *usecase.RateLimitUseCase
		if requestToken != "" {
			ratelimit = usecase.NewRateLimitUseCase(rl.MaxRequestToken, rl.BlockTimeToken, rl.rdb)
			dto = usecase.RateLimitInputDTO{Key: requestToken, Context: r.Context()}
		} else {
			log.Println("RateLimitMiddleware Token not found")
			dto = usecase.RateLimitInputDTO{Key: ip, Context: r.Context()}
			ratelimit = usecase.NewRateLimitUseCase(rl.MaxRequestIp, rl.BlockTimeIp, rl.rdb)

		}

		canAccess, _ := ratelimit.Execute(dto)

		if !canAccess {
			log.Println("RateLimitMiddleware You are blocked")
			http.Error(w, "You are blocked", http.StatusTooManyRequests)
			return
		}

		log.Println("RateLimitMiddleware You are not blocked")
		next.ServeHTTP(w, r)
	})
}
