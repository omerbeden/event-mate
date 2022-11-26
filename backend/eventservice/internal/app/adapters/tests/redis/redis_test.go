package redis_test

import (
	"fmt"

	"github.com/omerbeden/event-mate/backend/eventservice/internal/app/domain/model"
)

type MockRedis struct {
}

func (*MockRedis) Exist(key string) (bool, error) {
	return true, nil
}

func (*MockRedis) GetFromCache(key string) (interface{}, error) {
	fmt.Println("Getting from cache")
	posts := []*model.Event{
		{Title: "test1", Category: "category1"},
		{Title: "test2", Category: "category2"},
		{Title: "test3", Category: "category3"},
	}
	return posts, nil
}

func (m *MockRedis) AddToCache(key string, value interface{}) error {
	fmt.Println("item added to cache")
	return nil
}

// func TestSet(t *testing.T) {

// 	mockRedis := &MockRedis{}
// 	input := []model.Event{
// 		{Title: "test1", Category: "category1"},
// 		{Title: "test2", Category: "category2"},
// 		{Title: "test3", Category: "category3"},
// 	}

// 	sut := cacheadapter.Set("Sakarya", input, mockRedis)

// 	assert.NoError(t, sut)
// }

// func TestGet(t *testing.T) {
// 	mockRedis := &MockRedis{}

// 	sut, err := cacheadapter.GetPosts("Sakarya", mockRedis)

// 	assert.NotEmpty(t, sut)
// 	assert.NotNil(t, sut)
// 	assert.NoError(t, err)
// 	assert.Greater(t, len(sut), 0)

// }

// func TestExist(t *testing.T) {
// 	MockRedis := &MockRedis{}

// 	sut, _ := cacheadapter.Exist("Sakarya", MockRedis)

// 	assert.True(t, sut)
// }