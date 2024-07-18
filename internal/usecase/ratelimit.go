package usecase

import (
	"context"
	"strconv"
	"time"

	"github.com/redis/go-redis/v9"
)

type RateLimitUseCase struct {
	MaxRequests int
	BlockTime   int
	rdb         *redis.Client
}

func NewRateLimitUseCase(maxRequests, blockTime int, rdb *redis.Client) *RateLimitUseCase {
	return &RateLimitUseCase{
		MaxRequests: maxRequests,
		BlockTime:   blockTime,
		rdb:         rdb,
	}
}

type RateLimitInputDTO struct {
	Key     string
	Context context.Context
}

func (rl *RateLimitUseCase) Execute(input RateLimitInputDTO) (bool, error) {

	_, err := rl.rdb.Incr(input.Context, input.Key).Result()
	if err != nil {
		return false, err
	}

	requestsStr, err := rl.rdb.Get(input.Context, input.Key).Result()
	if err != nil {
		return false, err
	}

	requests, err := strconv.Atoi(requestsStr)
	if err != nil {
		return false, err
	}

	if requests == 1 {
		rl.rdb.Expire(input.Context, input.Key, time.Duration(1)*time.Second)
	}

	if requests > rl.MaxRequests {
		if requests == rl.MaxRequests+1 {
			rl.rdb.Expire(input.Context, input.Key, time.Duration(rl.BlockTime)*time.Second)
		}
		return false, nil
	}

	// _, err := rl.rdb.Incr(r.Context(), ip)
	// if err != nil {
	// 	log.Println("Error incrementing request count:", err)
	// 	http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	// 	return
	// }

	// requestsStr, err := rl.rdb.Get(r.Context(), ip) // Supondo que Get retorne uma string e um erro
	// if err != nil {
	// 	log.Println("Error getting request count:", err)
	// 	http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	// 	return
	// }

	// requests, err := strconv.Atoi(requestsStr) // Converte de string para int
	// if err != nil {
	// 	log.Println("Error converting request count to int:", err)
	// 	http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	// 	return
	// }

	// if requests == 1 {
	// 	log.Println("RateLimitMiddleware Expire -> 1")
	// 	rl.rdb.Expire(r.Context(), ip, time.Duration(1)*time.Second)
	// }

	// log.Println("RateLimitMiddleware requests -> ", requests)
	// log.Println("RateLimitMiddleware MaxRequestIp -> ", rl.MaxRequestIp)
	// if requests > rl.MaxRequestIp {
	// 	if requests == rl.MaxRequestIp+1 {
	// 		log.Println("RateLimitMiddleware Block -> ", rl.BlockTime)
	// 		rl.rdb.Set(r.Context(), ip, requests, time.Duration(rl.BlockTime)*time.Second)
	// 	}
	// 	log.Println("RateLimitMiddleware You are blocked")
	// 	http.Error(w, "You are blocked", http.StatusTooManyRequests)
	// 	return
	// }
	return true, nil
}
