package redisadapter_test

import (
	"encoding/json"
	"testing"

	"github.com/go-redis/redis/v8"
	"github.com/omerbeden/event-mate/backend/tatooine/modules/event/app/adapters/redisadapter"
	"github.com/omerbeden/event-mate/backend/tatooine/modules/event/app/domain/model"
)

const Adress = "Localhost:6379"
const Password = ""
const DB = 0

var testKey = "1"
var testValue = model.Event{
	ID:           1,
	Title:        "Integration Test Event",
	Category:     "Test2",
	CreatedBy:    model.User{ID: 1},
	Location:     model.Location{EventId: 1, City: "Sakarya"},
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
