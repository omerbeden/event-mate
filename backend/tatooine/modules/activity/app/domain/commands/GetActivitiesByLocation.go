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
	cityKey := fmt.Sprintf("%s:%s", CITY_KEY, gc.Location.City)

	activities, redisErr := gc.Redis.GetMembers(cityKey)
	if redisErr != nil {
		fmt.Printf("redis error %s \n returning from db", redisErr.Error()) // log error
		return gc.Repo.GetByLocation(&gc.Location)
	}

	var activitiesResult []model.Activity
	for _, activity := range activities {

		var activityObject = model.Activity{}
		err := json.Unmarshal([]byte(activity), &activityObject)
		if err != nil {
			fmt.Printf("parsing erorr returning from db %s", err.Error())
			return gc.Repo.GetByLocation(&gc.Location)
		}

		if activityObject.Location.City == gc.Location.City {
			activitiesResult = append(activitiesResult, activityObject)
		}
	}

	if activitiesResult != nil {
		fmt.Printf("returning activities from redis, l: %d\n", len(activities))
		return activitiesResult, nil
	}

	return gc.Repo.GetByLocation(&gc.Location)
}
