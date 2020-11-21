package breakerstorage

import (
	"context"
	"fmt"
	"time"

	"github.com/go-redis/redis"
)

var redisCtx = context.Background()

var baseKey = "breakerstorage/redis/"

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

func (ri RedisImpl) OpenCircuitList(svc string) ([]string, error) {
	key := fmt.Sprint(baseKey, svc)

	val, err := ri.client.LRange(key, 0, 100).Result()

	if err != nil {
		if err == redis.Nil {
			return []int{}, fmt.Errorf("key does not exists")
		}

		panic(err)
	}

	return val
}

func (ri RedisImpl) ClearOpenCircuitList() {}

func (ri RedisImpl) AddSuccess(svc string) error {
	key := fmt.Sprint(baseKey, svc)

	_, err := ri.client.LPush(key, 1).Result()

	if err != nil {
		if err == redis.Nil {
			return fmt.Errorf("key does not exists")
		}

		panic(err)
	}

	return nil
}

func (ri RedisImpl) AddError() {}

func (ri RedisImpl) SetLastErrorOcurredAt() {

}

func (ri RedisImpl) LastErrorOcurredAt() time.Time {}
