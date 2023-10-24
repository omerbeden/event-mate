package redis_test

import "github.com/omerbeden/event-mate/backend/tatooine/modules/event/app/domain/model"

const Adress = "Localhost:6379"
const Password = ""
const DB = 0

var testKey = "integrationTest"
var pushObj = model.Event{Title: "Integration Test Event", Category: "Test2"}

/*
func TestPushIntegration(t *testing.T) {
	client := redis.NewClient(&redis.Options{Addr: Adress, Password: Password, DB: DB})
	defer client.Close()
	redis := cacheadapter.NewRedisAdapter(client)
	err := cacheadapter.Push(testKey, pushObj, redis)

	if err != nil {
		t.Errorf("got err: %q", err)
	}
}

func TestGetPostsIntegration(t *testing.T) {
	client := redis.NewClient(&redis.Options{Addr: Adress, Password: Password, DB: DB})
	defer client.Close()

	redis := cacheadapter.NewRedisAdapter(redis.NewClient(&redis.Options{Addr: Adress, Password: Password, DB: DB}))
	posts, err := cacheadapter.GetPosts(testKey, redis)
	t.Logf("%+v\n", posts)
	if err != nil {
		t.Errorf("got err: %q", err)
	}
	if posts == nil {
		t.Error("got nill")
	}

	if len(posts) < 1 {
		t.Error("got posts empty")
	}

}

func TestGetEventIntegration(t *testing.T) {
	client := redis.NewClient(&redis.Options{Addr: Adress, Password: Password, DB: DB})
	defer client.Close()

	redis := cacheadapter.NewRedisAdapter(redis.NewClient(&redis.Options{Addr: Adress, Password: Password, DB: DB}))

	event, err := cacheadapter.GetEvent(0, testKey, redis)
	if err != nil {
		t.Errorf("got err: %q", err)
	}
	if event.Category != pushObj.Category {
		t.Errorf("got %q but want %q", event.Category, pushObj.Category)
	}
	if event.ID != 0 {
		t.Errorf("got %q but want %q", event.ID, 0)
	}

}

func TestExistIntegration(t *testing.T) {
	client := redis.NewClient(&redis.Options{Addr: Adress, Password: Password, DB: DB})
	defer client.Close()

	redis := cacheadapter.NewRedisAdapter(redis.NewClient(&redis.Options{Addr: Adress, Password: Password, DB: DB}))
	result, err := cacheadapter.Exist(testKey, redis)
	if err != nil {
		t.Errorf("got err: %q", err)
	}
	if result == false {
		t.Errorf("got %t but want %t", result, true)
	}
}*/
