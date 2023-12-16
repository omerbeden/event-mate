package redisadapter_test

import (
	"encoding/json"
	"testing"

	"github.com/go-redis/redis/v8"
	"github.com/omerbeden/event-mate/backend/tatooine/modules/activity/app/adapters/redisadapter"
	"github.com/omerbeden/event-mate/backend/tatooine/modules/activity/app/domain/model"
	"github.com/stretchr/testify/assert"
)

const Adress = "Localhost:6379"
const Password = ""
const DB = 0

var testKey = "1"
var testValue = model.Activity{
	ID:           1,
	Title:        "Integration Test Activity",
	Category:     "Test2",
	CreatedBy:    model.User{ID: 1},
	Location:     model.Location{ActivityId: 1, City: "Sakarya"},
	Participants: []model.User{{ID: 1}, {ID: 2}, {ID: 3}},
}

func TestSetRedisValue(t *testing.T) {
	client := redis.NewClient(&redis.Options{Addr: Adress, Password: Password, DB: DB})
	defer client.Close()

	redis_adapter := redisadapter.NewRedisAdapter(client)
	jsonValue, errJson := json.Marshal(testValue)
	if errJson != nil {
		t.Errorf("got err: %q", errJson)

	}
	err := redis_adapter.Set(testKey, jsonValue)

	if err != nil {
		t.Errorf("got err: %q", err)
	}
}

func TestGetRedisValue(t *testing.T) {
	client := redis.NewClient(&redis.Options{Addr: Adress, Password: Password, DB: DB})
	defer client.Close()

	redis_adapter := redisadapter.NewRedisAdapter(client)
	value, err := redis_adapter.Get(testKey)

	t.Logf("%+v\n", value)
	if err != nil {
		t.Errorf("got err: %q", err)
	}
	if value == nil {
		t.Error("got nill")
	}
}

func TestAddMembersAndGetMembers(t *testing.T) {
	client := redis.NewClient(&redis.Options{Addr: Adress, Password: Password, DB: DB})
	defer client.Close()

	tests := []struct {
		name       string
		activity   model.Activity
		activities []model.Activity
		verifyAdd  func(err error)
		verifyGet  func(result []string, err error)
	}{
		{
			name: "Add multiple activity at once",
			activities: []model.Activity{
				{
					ID:           11,
					Title:        "Integration Test Activity",
					Category:     "Test2",
					CreatedBy:    model.User{ID: 1},
					Location:     model.Location{ActivityId: 11, City: "Sakarya"},
					Participants: []model.User{{ID: 1}, {ID: 2}, {ID: 3}},
				},
				{
					ID:           22,
					Title:        "Integration Test Activity",
					Category:     "Test2",
					CreatedBy:    model.User{ID: 1},
					Location:     model.Location{ActivityId: 22, City: "Sakarya"},
					Participants: []model.User{{ID: 1}, {ID: 2}, {ID: 3}},
				},
				{
					ID:           33,
					Title:        "Integration Test Activity",
					Category:     "Test2",
					CreatedBy:    model.User{ID: 1},
					Location:     model.Location{ActivityId: 33, City: "Sakarya"},
					Participants: []model.User{{ID: 1}, {ID: 2}, {ID: 3}},
				},
			},
			verifyAdd: func(err error) {
				assert.NoError(t, err)
			},
			verifyGet: func(result []string, err error) {
				if assert.NoError(t, err) {
					assert.NotNil(t, result)
					assert.NotEmpty(t, result)
				}
			},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			redisAapter := redisadapter.NewRedisAdapter(client)
			jsonValue, errJson := json.Marshal(&tc.activities)
			if errJson != nil {
				t.Errorf("got err: %q", errJson)
			}

			err := redisAapter.AddMember("city:Sakarya", jsonValue)
			tc.verifyAdd(err)

			result, err := redisAapter.GetMembers("city:Sakarya")
			if tc.verifyGet != nil {
				tc.verifyGet(result, err)
			}
		})
	}
}
