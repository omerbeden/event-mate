package cachedapter

import (
	"encoding/json"
	"fmt"

	"github.com/omerbeden/event-mate/backend/tatooine/modules/profile/app/domain/model"
	"github.com/omerbeden/event-mate/backend/tatooine/pkg/cache"
)

var ErrLogPrefix = "profile:cacheAdapter"

const ATTANDED_ACTIVITIES_REDIS_KEY = "attandedActivities"

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
	err := adapter.client.Delete(key)
	if err != nil {
		return fmt.Errorf("%s could not delete key: %s ", ErrLogPrefix, key)
	}

	return nil
}

func (adapter *Cache) GetAttandedActivities(userId int64) ([]string, error) {
	key := fmt.Sprintf("%s:%d", ATTANDED_ACTIVITIES_REDIS_KEY, userId)
	attandedActivities, err := adapter.client.GetMembers(key)
	if err != nil {
		return nil, fmt.Errorf("%s could not get attanded activities for key: %s ", ErrLogPrefix, key)
	}

	return attandedActivities, nil
}

func (adapter *Cache) GetUserProfile(profileKey string) (*model.UserProfile, error) {

	cacheResult, err := adapter.client.Get(profileKey)
	if err != nil {
		return nil, fmt.Errorf("%s could not get user profile for key: %s ", ErrLogPrefix, profileKey)
	}

	var user model.UserProfile
	err = json.Unmarshal([]byte(cacheResult.(string)), &user)
	if err != nil {
		return nil, fmt.Errorf("%s could not unmarshal result of the cache key: %s ", ErrLogPrefix, profileKey)
	}

	return &user, nil

}
