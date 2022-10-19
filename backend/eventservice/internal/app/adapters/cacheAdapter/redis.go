package cacheadapter

import (
	"context"

	"github.com/go-redis/cache/v8"
	"github.com/go-redis/redis/v8"
	"github.com/omerbeden/event-mate/backend/eventservice/internal/app/domain/model"
	"github.com/omerbeden/event-mate/backend/eventservice/internal/app/domain/ports/caching"
)

type RedisAdapter struct {
	config redisConfig
}

func Set(key string, value interface{}, cacher caching.Cache) error {
	return cacher.AddToCache(key, value)
}

func (redisA *RedisAdapter) AddToCache(key string, value interface{}) error {
	client := redis.NewClient(&redis.Options{
		Addr: redisA.config.resourceName,
	})

	mycache := cache.New(&cache.Options{
		Redis: client,
	})

	if err := mycache.Set(&cache.Item{Key: key, Value: value}); err != nil {
		return err
	}
	return nil
}

func (redisA *RedisAdapter) GetPostsFromCache(key string) ([]*model.Event, error) {
	client := redis.NewClient(&redis.Options{
		Addr: redisA.config.resourceName,
	})

	mycache := cache.New(&cache.Options{
		Redis: client,
	})

	var wanted []*model.Event
	if err := mycache.Get(context.TODO(), key, wanted); err != nil {
		return nil, err
	}

	return wanted, nil
}

//todo: event modelinin icinde location bilgisi eklenecek, redis ve testler refactor edilip  tamamlanacak
