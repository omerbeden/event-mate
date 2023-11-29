package redisadapter

import (
	"context"
	"fmt"
	"time"

	"github.com/go-redis/redis/v8"
)

const expiration_time = 0
const err_prefix = "adapter:redis"

type RedisAdapter struct {
	client *redis.Client
}

func NewRedisAdapter(client *redis.Client) *RedisAdapter {
	return &RedisAdapter{
		client: client,
	}
}

func (adapter *RedisAdapter) Set(key string, value any) error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()
	_, err := adapter.client.Set(ctx, key, value, expiration_time).Result()

	if err != nil {
		return fmt.Errorf("%s could not set redis key-value %w", err_prefix, err)
	}
	return nil

}

func (adapter *RedisAdapter) Get(key string) (any, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	result, err := adapter.client.Get(ctx, key).Result()
	if err != nil {
		return nil, fmt.Errorf("%s could not set redis key-value", err_prefix)
	}
	return result, nil
}

func (adapter *RedisAdapter) Exist(key string) (bool, error) {
	return true, nil
}
