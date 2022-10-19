package caching

import "github.com/omerbeden/event-mate/backend/eventservice/internal/app/domain/model"

type Cache interface {
	AddToCache(key string, value interface{}) error
	GetPostsFromCache(key string) ([]*model.Event, error)
}
