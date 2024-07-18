package database

import (
	"fmt"

	"github.com/redis/go-redis/v9"
)

func NewRedis(host string, port string) *redis.Client {
	return redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%s", host, port),
		Password: "", // no password set
		DB:       0,  // use default DB
	})
}
