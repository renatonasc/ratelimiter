package usecase

import (
	"context"
	"renatonasc/ratelimit/internal/infra/database"
	"time"
)

type RateLimitUseCase struct {
	MaxRequests int
	BlockTime   int
	rdb         database.DBClient
}

func NewRateLimitUseCase(maxRequests, blockTime int, rdb database.DBClient) *RateLimitUseCase {
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

	requests, err := rl.rdb.Incr(input.Context, input.Key)
	if err != nil {
		return false, err
	}

	if requests == 1 {
		rl.rdb.Expire(input.Context, input.Key, time.Duration(1)*time.Second)
	}

	if int(requests) > rl.MaxRequests {
		if int(requests) == rl.MaxRequests+1 {
			rl.rdb.Expire(input.Context, input.Key, time.Duration(rl.BlockTime)*time.Second)
		}
		return false, nil
	}

	return true, nil
}
