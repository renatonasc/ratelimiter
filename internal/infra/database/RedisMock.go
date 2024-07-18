package database

import (
	"context"
	"time"

	"github.com/stretchr/testify/mock"
)

type RedisMock struct {
	count int
	mock.Mock
}

func NewRedisMock() *RedisMock {
	return new(RedisMock)
}

func (m *RedisMock) Expire(ctx context.Context, key string, expiration time.Duration) error {
	return nil
}

func (m *RedisMock) Incr(ctx context.Context, key string) (int64, error) {
	m.count++
	return int64(m.count), nil
}
