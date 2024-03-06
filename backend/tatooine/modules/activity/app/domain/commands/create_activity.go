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
	"github.com/omerbeden/event-mate/backend/tatooine/pkg/db"

	repo "github.com/omerbeden/event-mate/backend/tatooine/modules/activity/app/domain/ports/repositories"
)

type CreateCommand struct {
	Activity          model.Activity
	ActivityRepo      repo.ActivityRepository
	ActivityRulesRepo repo.ActivityRulesRepository
	ActivityFlowRepo  repo.ActivityFlowRepository
	LocRepo           repo.LocationRepository
	Redis             cache.Cache
	Tx                db.TransactionManager
}

func (cmd *CreateCommand) Handle(ctx context.Context) (bool, error) {

	tx, err := cmd.Tx.Begin(ctx)
	if err != nil {
		return false, fmt.Errorf("failed to begin transaction: %w", err)
	}

	defer func() {
		if err != nil {
			tx.Rollback(ctx)
		}
	}()

	activity, errCreate := cmd.ActivityRepo.Create(ctx, tx, cmd.Activity)
	if errCreate != nil {
		return false, errCreate
	}

	err = cmd.ActivityRulesRepo.CreateActivityRules(ctx, tx, activity.ID, cmd.Activity.Rules)
	if err != nil {
		return false, err
	}

	err = cmd.ActivityFlowRepo.CreateActivityFlow(ctx, tx, activity.ID, cmd.Activity.Flow)
	if err != nil {
		return false, err
	}

	_, errLoc := cmd.LocRepo.Create(ctx, tx, &activity.Location)
	if errLoc != nil {
		return false, errLoc
	}

	err = tx.Commit(ctx)
	if err != nil {
		return false, fmt.Errorf("failed to commit transaction: %w", err)
	}

	jsonActivity, errMarshall := json.Marshal(activity)
	if errMarshall != nil {
		return false, errMarshall
	}

	activityId := strconv.FormatInt(activity.ID, 10)
	activityKey := fmt.Sprintf("%s:%s", cacheadapter.ACTIVITY_CACHE_KEY, activityId)

	err = cmd.Redis.Set(ctx, activityKey, jsonActivity)
	if err != nil {
		fmt.Printf("activity could not inserted to Redis %s\n", activityId)
	}

	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
		cmd.addCityToRedis(ctx, activity.Location.City, jsonActivity)
	}()

	wg.Wait()
	return true, nil

}

func (ccmd *CreateCommand) addCityToRedis(ctx context.Context, city string, valueJson []byte) error {
	cityKey := fmt.Sprintf("%s:%s", cacheadapter.CITY_CACHE_KEY, city)

	return ccmd.Redis.AddMember(ctx, cityKey, valueJson)
}
