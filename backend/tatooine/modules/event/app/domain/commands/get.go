package commands

import (
	"encoding/json"
	"fmt"
	"strconv"

	"github.com/omerbeden/event-mate/backend/tatooine/modules/event/app/adapters/redisadapter"
	"github.com/omerbeden/event-mate/backend/tatooine/modules/event/app/domain/model"
	repo "github.com/omerbeden/event-mate/backend/tatooine/modules/event/app/domain/ports/repositories"
)

type GetByIDCommand struct {
	EventID int64
	Repo    repo.EventRepository
	Redis   redisadapter.RedisAdapter
}

func (gc *GetByIDCommand) Handle() (*model.Event, error) {

	eventId := strconv.FormatInt(gc.EventID, 10)
	result, redisErr := gc.Redis.Get(eventId)
	if redisErr != nil {
		fmt.Printf("redis error %s \n returning from db", redisErr.Error()) // log error
		return gc.Repo.GetByID(gc.EventID)
	}

	event := model.Event{}
	err := json.Unmarshal([]byte(result.(string)), &event)
	if err != nil {
		fmt.Printf("parsing erorr returning from db %s", err.Error())
		return gc.Repo.GetByID(gc.EventID)
	}

	if result != nil && err == nil {
		fmt.Printf("returning from redis %+v\n", event)
		return &event, err
	}

	return gc.Repo.GetByID(gc.EventID)
}

type GetByLocationCommand struct {
	Location model.Location
	Repo     repo.EventRepository
	Redis    redisadapter.RedisAdapter
}

func (gc *GetByLocationCommand) Handle() ([]model.Event, error) {

	result, redisErr := gc.Redis.Get(gc.Location.City)
	if redisErr != nil {
		fmt.Printf("redis error %s \n returning from db", redisErr.Error()) // log error
		return gc.Repo.GetByLocation(&gc.Location)
	}

	events := []model.Event{}
	err := json.Unmarshal([]byte(result.(string)), &events)
	if err != nil {
		fmt.Printf("parsing erorr returning from db %s", err.Error())
		return gc.Repo.GetByLocation(&gc.Location)
	}

	if result != nil && err == nil {
		fmt.Printf("returning events from redis, l: %d\n", len(events))
		return events, err
	}

	return gc.Repo.GetByLocation(&gc.Location)
}
