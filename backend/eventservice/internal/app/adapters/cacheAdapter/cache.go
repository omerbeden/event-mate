package cacheadapter

import (
	"github.com/omerbeden/event-mate/backend/eventservice/internal/app/domain/model"
	"github.com/omerbeden/event-mate/backend/eventservice/internal/app/domain/ports/caching"
	"golang.org/x/exp/slices"
)

func Push(key string, values interface{}, cacher caching.Cache) error {
	return cacher.Push(key, values)
}

func GetPosts(key string, cacher caching.Cache) ([]model.Event, error) {
	posts, err := cacher.GetPosts(key)
	return posts, err
}

func GetEvent(id int, city string, cacher caching.Cache) (model.Event, error) {
	posts, _ := cacher.GetPosts(city)
	idx := slices.IndexFunc(posts, func(event model.Event) bool {
		return int(event.ID) == id
	})
	if idx != -1 {
		return posts[idx], nil
	}

	return model.Event{}, nil // bura error dönücek
}
func Exist(key string, cacher caching.Cache) (bool, error) {
	return cacher.Exist(key)
}
