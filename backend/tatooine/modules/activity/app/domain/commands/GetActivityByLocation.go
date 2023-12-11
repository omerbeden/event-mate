package commands

import (
	"encoding/json"
	"fmt"

	"github.com/omerbeden/event-mate/backend/tatooine/modules/activity/app/domain/model"
	"github.com/omerbeden/event-mate/backend/tatooine/modules/activity/app/domain/ports/caching"
	repo "github.com/omerbeden/event-mate/backend/tatooine/modules/activity/app/domain/ports/repositories"
)

type GetByLocationCommand struct {
	Location model.Location
	Repo     repo.ActivityRepository
	Redis    caching.Cache
}

func (gc *GetByLocationCommand) Handle() ([]model.Activity, error) {

	result, redisErr := gc.Redis.Get(gc.Location.City)
	if redisErr != nil {
		fmt.Printf("redis error %s \n returning from db", redisErr.Error()) // log error
		return gc.Repo.GetByLocation(&gc.Location)
	}

	activities := []model.Activity{}
	err := json.Unmarshal([]byte(result.(string)), &activities)
	if err != nil {
		fmt.Printf("parsing erorr returning from db %s", err.Error())
		return gc.Repo.GetByLocation(&gc.Location)
	}

	if result != nil && err == nil {
		fmt.Printf("returning activities from redis, l: %d\n", len(activities))
		return activities, err
	}

	return gc.Repo.GetByLocation(&gc.Location)
}
