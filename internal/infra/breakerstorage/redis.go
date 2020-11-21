package breakerstorage

import (
	"context"
	"fmt"

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

func (ri RedisImpl) AddToList(list_key, value string) error {
	_, err := ri.client.LPush(list_key, 1).Result()

	if err != nil {
		if err == redis.Nil {
			return fmt.Errorf("key does not exists")
		}

		panic(err)
	}

	return nil
}

func (ri RedisImpl) GetList(list_key string, size int64) ([]string, error) {
	val, err := ri.client.LRange(list_key, 0, size).Result()

	if err != nil {
		if err == redis.Nil {
			return []string{}, fmt.Errorf("key does not exists")
		}

		panic(err)
	}

	return val, err
}

func (ri RedisImpl) EraseList(list_key string) error {
	_, err := ri.client.Del(list_key).Result()

	if err != nil {
		if err == redis.Nil {
			return fmt.Errorf("key does not exists")
		}

		panic(err)
	}

	return nil
}

func (ri RedisImpl) Put(key string, value interface{}) error {
	_, err := ri.client.Set(key, value, 0).Result()

	if err != nil {
		if err == redis.Nil {
			return fmt.Errorf("key does not exists")
		}

		panic(err)
	}

	return nil
}

func (ri RedisImpl) Get(key string) string {
	result, err := ri.client.Get(key).Result()

	if err != nil {
		if err == redis.Nil {
			fmt.Println("key does not exists")
		}

		panic(err)
	}

	return result
}
