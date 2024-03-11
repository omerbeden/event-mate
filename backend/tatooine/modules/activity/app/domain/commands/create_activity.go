package commands

import (
	"context"
	"encoding/json"
	"fmt"
	"strconv"
	"sync"

	"github.com/omerbeden/event-mate/backend/tatooine/modules/activity/app/adapters/cacheadapter"
	"github.com/omerbeden/event-mate/backend/tatooine/modules/activity/app/domain/model"
	"github.com/omerbeden/event-mate/backend/tatooine/pkg"
	"github.com/omerbeden/event-mate/backend/tatooine/pkg/cache"
	"github.com/omerbeden/event-mate/backend/tatooine/pkg/db"
	"go.uber.org/zap"

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

	logger, ok := ctx.Value(pkg.LoggerKey).(*zap.SugaredLogger)
	if !ok {
		return false, fmt.Errorf("failed to get logger for CreateCommand")
	}
	tx, err := cmd.Tx.Begin(ctx)
	if err != nil {
		return false, fmt.Errorf("failed to begin transaction: %w", err)
	}

	defer func() {
		if err != nil {
			tx.Rollback(ctx)
		}
	}()

	activity, err := cmd.ActivityRepo.Create(ctx, tx, cmd.Activity)
	if err != nil {
		return false, err
	}

	err = cmd.ActivityRulesRepo.CreateActivityRules(ctx, tx, activity.ID, cmd.Activity.Rules)
	if err != nil {
		return false, err
	}

	err = cmd.ActivityFlowRepo.CreateActivityFlow(ctx, tx, activity.ID, cmd.Activity.Flow)
	if err != nil {
		return false, err
	}

	_, err = cmd.LocRepo.Create(ctx, tx, &activity.Location)
	if err != nil {
		return false, err
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
		logger.Infof("activity could not inserted to Redis %s\n", activityId)
	}

	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		ctx := context.Background()
		defer wg.Done()
		err := cmd.addCityToRedis(ctx, activity.Location.City, jsonActivity)
		if err != nil {
			logger.Info("failed to add city to Redis %s\n", activity)
		}
	}()

	wg.Wait()
	return true, nil

}

func (ccmd *CreateCommand) addCityToRedis(ctx context.Context, city string, valueJson []byte) error {
	cityKey := fmt.Sprintf("%s:%s", cacheadapter.CITY_CACHE_KEY, city)
	return ccmd.Redis.AddMember(ctx, cityKey, valueJson)
}
