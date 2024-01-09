package cachedapter

import (
	"fmt"

	"github.com/omerbeden/event-mate/backend/tatooine/pkg/cache"
)

var ErrLogPrefix = "profile:cacheAdapter"
var attandedActivitiesRedisKey = "attandedActivities"

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

func (adapter *Cache) GetAttandedActivities(userId int64) ([]string, error) {
	key := fmt.Sprintf("%s:%d", attandedActivitiesRedisKey, userId)
	attandedActivities, err := adapter.client.GetMembers(key)
	if err != nil {
		return nil, fmt.Errorf("%s could not get attanded activities for key: %s ", ErrLogPrefix, key)
	}

	return attandedActivities, nil
}
