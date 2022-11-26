package cacheadapter

import (
	"github.com/omerbeden/event-mate/backend/eventservice/internal/app/domain/model"
	"github.com/omerbeden/event-mate/backend/eventservice/internal/app/domain/ports/caching"
)

func Set(key string, value interface{}, cacher caching.Cache) error {
	return cacher.AddToCache(key, value)
}

func GetPosts(key string, cacher caching.Cache) ([]model.Event, error) {
	posts, err := cacher.GetFromCache(key)
	return posts.([]model.Event), err
}
func Exist(key string, cacher caching.Cache) (bool, error) {
	return cacher.Exist(key)
}
