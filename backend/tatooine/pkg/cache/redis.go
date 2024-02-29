package cache

import (
	"context"
	"errors"
	"time"

	"github.com/go-redis/redis/v8"
)

var ErrRedisMember = errors.New("members Nil or Blank")

const ATTANDED_ACTIVITIES_CACHE_KEY = "attandedActivities"

type RedisOption struct {
	*redis.Options
	ExpirationTime time.Duration
}
type RedisClient struct {
	redis   *redis.Client
	options RedisOption
}

func NewRedisClient(option RedisOption) *RedisClient {

	redis := redis.NewClient(&redis.Options{
		Addr:     option.Addr,
		Password: option.Password,
		DB:       option.DB,
	})
	return &RedisClient{
		redis:   redis,
		options: option,
	}
}

func (client *RedisClient) Close() error {
	err := client.redis.Close()
	if err != nil {
		return err
	}

	return err
}

func (client *RedisClient) Set(ctx context.Context, key string, value any) error {

	_, err := client.redis.Set(ctx, key, value, client.options.ExpirationTime).Result()

	if err != nil {
		return err
	}
	return nil

}

func (client *RedisClient) Get(ctx context.Context, key string) (any, error) {

	result, err := client.redis.Get(ctx, key).Result()
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (client *RedisClient) Delete(ctx context.Context, key string) error {
	_, err := client.redis.Del(ctx, key).Result()
	if err != nil {
		return err
	}
	return nil
}

func (client *RedisClient) AddMember(ctx context.Context, key string, members ...any) error {

	if members == nil || len(members) < 1 {
		return ErrRedisMember
	}
	_, err := client.redis.SAdd(ctx, key, members).Result()
	if err != nil {
		return err
	}

	return nil

}

func (client *RedisClient) GetMembers(ctx context.Context, key string) ([]string, error) {
	result, err := client.redis.SMembers(ctx, key).Result()
	if err != nil {
		return nil, err
	}

	return result, err
}
