package cacheadapter

import (
	"context"
	"encoding/json"

	"github.com/go-redis/redis/v8"
	"github.com/omerbeden/event-mate/backend/eventservice/internal/app/domain/model"
)

type RedisAdapter struct {
	client *redis.Client
}

func NewRedisAdapter(options *redis.Options) *RedisAdapter {
	client := redis.NewClient(options)
	return &RedisAdapter{
		client: client,
	}
}

func (redisA *RedisAdapter) AddToCache(key string, values interface{}) error {
	valJson, jsonErr := json.Marshal(values)
	if jsonErr != nil {
		return jsonErr
	}
	_, err := redisA.client.LPush(context.TODO(), key, valJson).Result()
	return err
}

func (redisA *RedisAdapter) GetFromCache(key string) (interface{}, error) {

	var events []model.Event
	res, err := redisA.client.LRange(context.TODO(), key, 0, -1).Result()
	if err != nil {
		return nil, err
	}

	if err := json.Unmarshal([]byte(res[0]), &events); err != nil {
		return nil, err
	}

	return events, nil
}

func (redisA *RedisAdapter) UpdateCache(key string, value interface{}, newValue interface{}) error {

	if _, err := redisA.client.LRem(context.TODO(), key, 1, value).Result(); err != nil {
		return err
	}
	if _, err := redisA.client.LPush(context.TODO(), key, newValue).Result(); err != nil {
		return err
	}
	return nil
}

func (redisA *RedisAdapter) Exist(key string) (bool, error) {

	if _, err := redisA.client.Exists(context.TODO(), key).Result(); err != nil {
		return false, err
	}
	return true, nil
}
