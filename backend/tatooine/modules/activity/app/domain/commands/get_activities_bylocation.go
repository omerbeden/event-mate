package commands

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/omerbeden/event-mate/backend/tatooine/modules/activity/app/adapters/cacheadapter"
	"github.com/omerbeden/event-mate/backend/tatooine/modules/activity/app/domain/model"
	repo "github.com/omerbeden/event-mate/backend/tatooine/modules/activity/app/domain/ports/repositories"
	"github.com/omerbeden/event-mate/backend/tatooine/pkg/cache"
)

type GetByLocationCommand struct {
	Location model.Location
	Repo     repo.ActivityRepository
	Redis    cache.Cache
}

func (gc *GetByLocationCommand) Handle(ctx context.Context) ([]model.Activity, error) {
	cityKey := fmt.Sprintf("%s:%s", cacheadapter.CITY_CACHE_KEY, gc.Location.City)

	activities, redisErr := gc.Redis.GetMembers(ctx, cityKey)
	if redisErr != nil {
		fmt.Printf("redis error %s \n returning from db", redisErr.Error()) // log error
		return gc.Repo.GetByLocation(ctx, &gc.Location)
	}

	var activitiesResult []model.Activity
	for _, activity := range activities {

		var activityObject = model.Activity{}
		err := json.Unmarshal([]byte(activity), &activityObject)
		if err != nil {
			fmt.Printf("parsing erorr returning from db %s", err.Error())
			return gc.Repo.GetByLocation(ctx, &gc.Location)
		}

		if activityObject.Location.City == gc.Location.City {
			activitiesResult = append(activitiesResult, activityObject)
		}
	}

	if activitiesResult != nil {
		fmt.Printf("returning activities from redis, l: %d\n", len(activities))
		return activitiesResult, nil
	}

	return gc.Repo.GetByLocation(ctx, &gc.Location)
}
