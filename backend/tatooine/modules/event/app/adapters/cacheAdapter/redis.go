package cacheadapter

import (
	"context"
	"encoding/json"
	"strings"

	"github.com/go-redis/redis/v8"
	"github.com/omerbeden/event-mate/backend/tatooine/modules/event/app/domain/model"
)

type RedisAdapter struct {
	client *redis.Client
}

func NewRedisAdapter(client *redis.Client) *RedisAdapter {
	return &RedisAdapter{
		client: client,
	}
}

func (redisA *RedisAdapter) Push(key string, values interface{}) error {
	valJson, jsonErr := json.Marshal(values)

	if jsonErr != nil {
		return jsonErr
	}
	_, err := redisA.client.LPush(context.TODO(), key, valJson).Result()
	return err
}

func (redisA *RedisAdapter) GetPosts(key string) ([]model.Event, error) {

	var events []model.Event
	res, err := redisA.client.LRange(context.TODO(), key, 0, -1).Result()
	if err != nil {
		return nil, err
	}

	separeted := strings.Join(res, ",")
	first := append([]string{"["}, separeted)
	result := append(first, "]")

	if err := json.Unmarshal([]byte(strings.Join(result, "")), &events); err != nil {
		return nil, err
	}

	return events, nil
}

func (redisA *RedisAdapter) Exist(key string) (bool, error) {

	if _, err := redisA.client.Exists(context.TODO(), key).Result(); err != nil {
		return false, err
	}
	return true, nil
}