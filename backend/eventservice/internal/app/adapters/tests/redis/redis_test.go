package redis_test

import (
	"fmt"
	"testing"

	cacheadapter "github.com/omerbeden/event-mate/backend/eventservice/internal/app/adapters/cacheAdapter"
	"github.com/omerbeden/event-mate/backend/eventservice/internal/app/domain/model"
	"github.com/stretchr/testify/assert"
)

type MockRedis struct {
}

// GetPostsFromCache implements caching.Cache
func (*MockRedis) GetPostsFromCache(key string) ([]*model.Event, error) {
	panic("unimplemented")
}

func (m *MockRedis) AddToCache(key string, value interface{}) error {
	fmt.Println("item added to cache")
	return nil
}

func TestSet(t *testing.T) {

	mockRedis := &MockRedis{}
	input := []model.Event{
		{Title: "test1", Category: "category1"},
		{Title: "test2", Category: "category2"},
		{Title: "test3", Category: "category3"},
	}

	sut := cacheadapter.Set("Sakarya", input, mockRedis)

	assert.NoError(t, sut)
}
