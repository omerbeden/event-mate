package cacheadapter

import (
	"github.com/omerbeden/event-mate/backend/eventservice/internal/app/domain/model"
	"github.com/omerbeden/event-mate/backend/eventservice/internal/app/domain/ports/caching"
)

func Push(key string, values interface{}, cacher caching.Cache) error {
	return cacher.Push(key, values)
}

func GetPosts(key string, cacher caching.Cache) ([]model.Event, error) {
	posts, err := cacher.GetPosts(key)
	return posts, err
}
func Exist(key string, cacher caching.Cache) (bool, error) {
	return cacher.Exist(key)
}
