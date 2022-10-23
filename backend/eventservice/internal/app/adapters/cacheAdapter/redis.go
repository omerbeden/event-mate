package cacheadapter

import (
	"context"

	"github.com/go-redis/cache/v8"
	"github.com/go-redis/redis/v8"
	"github.com/omerbeden/event-mate/backend/eventservice/internal/app/domain/model"
)

type RedisAdapter struct {
	cache *cache.Cache
}

func InitRedis(redisConfig redisConfig) *cache.Cache {
	client := redis.NewClient(&redis.Options{
		Addr: redisConfig.resourceName,
	})

	return cache.New(&cache.Options{
		Redis: client,
	})
}

func (redisA *RedisAdapter) AddToCache(key string, value interface{}) error {

	if err := redisA.cache.Set(&cache.Item{Key: key, Value: value}); err != nil {
		return err
	}
	return nil
}

func (redisA *RedisAdapter) GetFromCache(key string) (interface{}, error) {

	var wanted []model.Event
	if err := redisA.cache.Get(context.TODO(), key, wanted); err != nil {
		return nil, err
	}

	return wanted, nil
}

func (redisA *RedisAdapter) UpdateCache(key string, value interface{}) error {

	if err := redisA.cache.Delete(context.TODO(), key); err != nil {
		return err
	}
	if err := redisA.cache.Set(&cache.Item{Key: key, Value: value}); err != nil {
		return err
	}

	return nil
}

func (redisA *RedisAdapter) Exist(key string) bool {

	return redisA.cache.Exists(context.TODO(), key)
}
