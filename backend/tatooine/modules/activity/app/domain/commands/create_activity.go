package commands

import (
	"context"
	"encoding/json"
	"fmt"
	"strconv"
	"sync"

	"github.com/omerbeden/event-mate/backend/tatooine/modules/activity/app/adapters/cacheadapter"
	"github.com/omerbeden/event-mate/backend/tatooine/modules/activity/app/domain/model"
	"github.com/omerbeden/event-mate/backend/tatooine/pkg/cache"

	repo "github.com/omerbeden/event-mate/backend/tatooine/modules/activity/app/domain/ports/repositories"
)

type CreateCommand struct {
	Activity          model.Activity
	ActivityRepo      repo.ActivityRepository
	ActivityRulesRepo repo.ActivityRulesRepository
	ActivityFlowRepo  repo.ActivityFlowRepository
	LocRepo           repo.LocationRepository
	Redis             cache.Cache
}

func (ccmd *CreateCommand) Handle(ctx context.Context) (bool, error) {

	activity, errCreate := ccmd.ActivityRepo.Create(ctx, ccmd.Activity)
	if errCreate != nil {
		return false, errCreate
	}

	err := ccmd.ActivityRulesRepo.CreateActivityRules(ctx, activity.ID, ccmd.Activity.Rules)
	if err != nil {
		return false, err
	}

	err = ccmd.ActivityFlowRepo.CreateActivityFlow(ctx, activity.ID, ccmd.Activity.Flow)
	if err != nil {
		return false, err
	}

	_, errLoc := ccmd.LocRepo.Create(ctx, &activity.Location)
	if errLoc != nil {
		return false, errLoc
	}

	jsonActivity, errMarshall := json.Marshal(activity)
	if errMarshall != nil {
		return false, errMarshall
	}

	activityId := strconv.FormatInt(activity.ID, 10)
	activityKey := fmt.Sprintf("%s:%s", cacheadapter.ACTIVITY_CACHE_KEY, activityId)

	err = ccmd.Redis.Set(ctx, activityKey, jsonActivity)
	if err != nil {
		fmt.Printf("activity could not inserted to Redis %s\n", activityId)
	}

	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
		ccmd.addCityToRedis(ctx, activity.Location.City, jsonActivity)
	}()

	wg.Wait()
	return true, nil

}

func (ccmd *CreateCommand) addCityToRedis(ctx context.Context, city string, valueJson []byte) error {
	cityKey := fmt.Sprintf("%s:%s", cacheadapter.CITY_CACHE_KEY, city)

	return ccmd.Redis.AddMember(ctx, cityKey, valueJson)
}