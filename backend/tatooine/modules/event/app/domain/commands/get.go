package commands

import (
	"encoding/json"
	"fmt"
	"strconv"

	"github.com/omerbeden/event-mate/backend/tatooine/modules/event/app/adapters/redisadapter"
	"github.com/omerbeden/event-mate/backend/tatooine/modules/event/app/domain/model"
	repo "github.com/omerbeden/event-mate/backend/tatooine/modules/event/app/domain/ports/repositories"
)

type GetCommand struct {
	EventID   string
	EventCity string
	Repo      repo.EventRepository
	Redis     redisadapter.RedisAdapter
}

func (gc *GetCommand) Handle() (*model.Event, error) {

	intID, err := strconv.Atoi(gc.EventID)
	if err != nil {
		return nil, fmt.Errorf("get command: %w", err)
	}

	result, redisErr := gc.Redis.Get(gc.EventID)
	if redisErr != nil {
		fmt.Printf("redis error %s \n returning from db", redisErr.Error()) // log error
		return gc.Repo.GetByID(int64(intID))
	}

	event := model.Event{}
	err = json.Unmarshal([]byte(result.(string)), &event)
	if err != nil {
		fmt.Printf("parsing erorr returning from db %s", err.Error())
		return gc.Repo.GetByID(int64(intID))
	}

	if result != nil && err == nil {
		fmt.Printf("returning from redis %+v\n", event)
		return &event, err
	}

	return gc.Repo.GetByID(int64(intID))
}
