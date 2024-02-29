package commands

import (
	"context"
	"encoding/json"
	"fmt"
	"strconv"

	"github.com/omerbeden/event-mate/backend/tatooine/modules/activity/app/adapters/cacheadapter"
	"github.com/omerbeden/event-mate/backend/tatooine/modules/activity/app/domain/model"
	repo "github.com/omerbeden/event-mate/backend/tatooine/modules/activity/app/domain/ports/repositories"
	"github.com/omerbeden/event-mate/backend/tatooine/pkg/cache"
)

type GetByIDCommand struct {
	ActivityId        int64
	Repo              repo.ActivityRepository
	ActivityRulesRepo repo.ActivityRulesRepository
	ActivityFlowRepo  repo.ActivityFlowRepository
	Redis             cache.Cache
}

func (gc *GetByIDCommand) Handle(ctx context.Context) (*model.Activity, error) {

	activityId := strconv.FormatInt(gc.ActivityId, 10)
	activityKey := fmt.Sprintf("%s:%s", cacheadapter.ACTIVITY_CACHE_KEY, activityId)

	result, redisErr := gc.Redis.Get(ctx, activityKey)
	if redisErr != nil {
		fmt.Printf("redis error %s \n returning from db", redisErr.Error()) // log error
		return gc.getActivityFromDb(ctx)
	}

	activity := model.Activity{}
	err := json.Unmarshal([]byte(result.(string)), &activity)
	if err != nil {
		fmt.Printf("parsing erorr returning from db %s", err.Error())
		return gc.getActivityFromDb(ctx)
	}

	if result != nil {
		fmt.Printf("returning from redis %+v\n", activity)
		return &activity, err
	}

	return gc.getActivityFromDb(ctx)
}

func (gc *GetByIDCommand) getActivityFromDb(ctx context.Context) (*model.Activity, error) {
	activity, err := gc.Repo.GetByID(ctx, gc.ActivityId)
	if err != nil {
		return nil, err
	}

	rules, err := gc.ActivityRulesRepo.GetActivityRules(ctx, gc.ActivityId)
	if err != nil {
		return nil, err
	}

	activity.Rules = rules

	flow, err := gc.ActivityFlowRepo.GetActivityFlow(ctx, gc.ActivityId)
	if err != nil {
		return nil, err
	}

	activity.Flow = flow

	return activity, nil

}
