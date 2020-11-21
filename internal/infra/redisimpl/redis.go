package redisimpl

import (
	"context"

	"github.com/go-redis/redis"
)

var redisCtx = context.Background()

type RedisImpl struct {
	client *redis.Client
}

func NewRedisImpl() RedisImpl {
	return RedisImpl{
		redis.NewClient(&redis.Options{
			Addr:     "localhost:6379",
			Password: "",
			DB:       0,
		}),
	}
}
