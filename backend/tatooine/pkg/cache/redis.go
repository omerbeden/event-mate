package cache

import (
	"context"
	"errors"
	"time"

	"github.com/go-redis/redis/v8"
)

var ErrRedisMember = errors.New("members Nil or Blank")

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

func (client *RedisClient) Set(key string, value any) error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	_, err := client.redis.Set(ctx, key, value, client.options.ExpirationTime).Result()

	if err != nil {
		return err
	}
	return nil

}

func (client *RedisClient) Get(key string) (any, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	result, err := client.redis.Get(ctx, key).Result()
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (client *RedisClient) Delete(key string) error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	_, err := client.redis.Del(ctx, key).Result()
	if err != nil {
		return err
	}
	return nil
}

func (client *RedisClient) AddMember(key string, members ...any) error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	if members == nil || len(members) < 1 {
		return ErrRedisMember
	}
	_, err := client.redis.SAdd(ctx, key, members).Result()
	if err != nil {
		return err
	}

	return nil

}

func (client *RedisClient) GetMembers(key string) ([]string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()
	result, err := client.redis.SMembers(ctx, key).Result()
	if err != nil {
		return nil, err
	}

	return result, err
}
