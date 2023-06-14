package caching

import "github.com/omerbeden/event-mate/backend/tatooine/modules/event/app/domain/model"

type Cache interface {
	Push(key string, values interface{}) error
	GetPosts(key string) ([]model.Event, error)
	Exist(key string) (bool, error)
}