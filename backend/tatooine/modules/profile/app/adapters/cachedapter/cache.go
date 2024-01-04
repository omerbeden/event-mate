package cachedapter

import (
	"fmt"

	"github.com/omerbeden/event-mate/backend/tatooine/pkg/cache"
)

var ErrLogPrefix = "profile:cacheAdapter"

type Cache struct {
	client cache.Cache
}

func NewCache(client cache.Cache) *Cache {
	return &Cache{
		client: client,
	}
}

func (adapter *Cache) Set(key string, jsonValue []byte) error {
	err := adapter.client.Set(key, jsonValue)
	if err != nil {
		return fmt.Errorf("%s , could not set cache key  for: %s  , err:  %w", ErrLogPrefix, key, err)
	}

	return nil

}

func (adapter *Cache) Delete(key string) error {
	return adapter.client.Delete(key)
}
