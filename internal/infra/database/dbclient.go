package database

import (
	"context"
	"time"
)

type DBClient interface {
	Expire(ctx context.Context, key string, expiration time.Duration) error
	Incr(ctx context.Context, key string) (int64, error)
}
