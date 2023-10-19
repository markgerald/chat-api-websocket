package rediscache

import (
	"context"
	"github.com/redis/go-redis/v9"
)

var ctx = context.Background()

func Client() *redis.Client {
	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "MgM23",
		DB:       0,
	})
	return rdb
}
