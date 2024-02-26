package commands

import (
	"encoding/json"
	"fmt"
	"strconv"

	"github.com/omerbeden/event-mate/backend/tatooine/modules/activity/app/domain/model"
	"github.com/omerbeden/event-mate/backend/tatooine/modules/activity/app/domain/ports/caching"
	repo "github.com/omerbeden/event-mate/backend/tatooine/modules/activity/app/domain/ports/repositories"
)

type GetByIDCommand struct {
	ActivityId        int64
	Repo              repo.ActivityRepository
	ActivityRulesRepo repo.ActivityRulesRepository
	Redis             caching.Cache
}

func (gc *GetByIDCommand) Handle() (*model.Activity, error) {

	activityId := strconv.FormatInt(gc.ActivityId, 10)
	activityKey := fmt.Sprintf("%s:%s", ACTIVITY_KEY, activityId)

	result, redisErr := gc.Redis.Get(activityKey)
	if redisErr != nil {
		fmt.Printf("redis error %s \n returning from db", redisErr.Error()) // log error
		return gc.getActivityFromDb()
	}

	activity := model.Activity{}
	err := json.Unmarshal([]byte(result.(string)), &activity)
	if err != nil {
		fmt.Printf("parsing erorr returning from db %s", err.Error())
		return gc.getActivityFromDb()
	}

	if result != nil && err == nil {
		fmt.Printf("returning from redis %+v\n", activity)
		return &activity, err
	}

	return gc.Repo.GetByID(gc.ActivityId)
}

func (gc *GetByIDCommand) getActivityFromDb() (*model.Activity, error) {
	activity, err := gc.Repo.GetByID(gc.ActivityId)
	if err != nil {
		return nil, err
	}

	rules, err := gc.ActivityRulesRepo.GetActivityRules(gc.ActivityId)
	if err != nil {
		return nil, err
	}

	activity.Rules = rules

	return activity, nil

}
