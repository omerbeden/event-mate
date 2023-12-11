package commands

import (
	"encoding/json"
	"fmt"
	"strconv"

	"github.com/omerbeden/event-mate/backend/tatooine/modules/activity/app/adapters/redisadapter"
	"github.com/omerbeden/event-mate/backend/tatooine/modules/activity/app/domain/model"
	repo "github.com/omerbeden/event-mate/backend/tatooine/modules/activity/app/domain/ports/repositories"
)

type GetByIDCommand struct {
	ActivityId int64
	Repo       repo.ActivityRepository
	Redis      redisadapter.RedisAdapter
}

func (gc *GetByIDCommand) Handle() (*model.Activity, error) {

	activityId := strconv.FormatInt(gc.ActivityId, 10)
	result, redisErr := gc.Redis.Get(activityId)
	if redisErr != nil {
		fmt.Printf("redis error %s \n returning from db", redisErr.Error()) // log error
		return gc.Repo.GetByID(gc.ActivityId)
	}

	activity := model.Activity{}
	err := json.Unmarshal([]byte(result.(string)), &activity)
	if err != nil {
		fmt.Printf("parsing erorr returning from db %s", err.Error())
		return gc.Repo.GetByID(gc.ActivityId)
	}

	if result != nil && err == nil {
		fmt.Printf("returning from redis %+v\n", activity)
		return &activity, err
	}

	return gc.Repo.GetByID(gc.ActivityId)
}

type GetByLocationCommand struct {
	Location model.Location
	Repo     repo.ActivityRepository
	Redis    redisadapter.RedisAdapter
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
