package cacheadapter

import (
	"fmt"

	"github.com/omerbeden/event-mate/backend/eventservice/internal/app/domain"
	"github.com/omerbeden/event-mate/backend/eventservice/internal/app/domain/model"
	"github.com/omerbeden/event-mate/backend/eventservice/internal/app/domain/ports/caching"
	"golang.org/x/exp/slices"
)

func Push(key string, values interface{}, cacher caching.Cache) error {

	err := cacher.Push(key, values)
	if err != nil {
		return fmt.Errorf("cache push: %w", err)
	}
	return nil
}

func GetPosts(key string, cacher caching.Cache) ([]model.Event, error) {
	posts, err := cacher.GetPosts(key)
	if err != nil {
		return nil, fmt.Errorf("cache get posts: %w", err)
	}
	return posts, err
}

func GetEvent(id int, city string, cacher caching.Cache) (model.Event, error) {
	posts, err := cacher.GetPosts(city)
	if err != nil {
		return model.Event{}, fmt.Errorf("cache get event: %w", err)
	}
	idx := slices.IndexFunc(posts, func(event model.Event) bool {
		return int(event.ID) == id
	})

	if idx == -1 {
		return model.Event{}, domain.ErrCacheItemNotFound // nil dönülecek

	}

	return posts[idx], nil
}
func Exist(key string, cacher caching.Cache) (bool, error) {

	isExist, err := cacher.Exist(key)
	if err != nil {
		return isExist, fmt.Errorf("cache exist: %w", err)
	}

	return isExist, nil
}
