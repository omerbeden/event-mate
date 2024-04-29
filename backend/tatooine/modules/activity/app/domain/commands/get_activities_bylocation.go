package commands

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/omerbeden/event-mate/backend/tatooine/modules/activity/app/adapters/cacheadapter"
	"github.com/omerbeden/event-mate/backend/tatooine/modules/activity/app/domain/model"
	repo "github.com/omerbeden/event-mate/backend/tatooine/modules/activity/app/domain/ports/repositories"
	"github.com/omerbeden/event-mate/backend/tatooine/pkg"
	"github.com/omerbeden/event-mate/backend/tatooine/pkg/cache"
	"go.uber.org/zap"
)

type GetByLocationCommand struct {
	Location model.Location
	Repo     repo.ActivityRepository
	Redis    cache.Cache
}

func (gc *GetByLocationCommand) Handle(ctx context.Context) ([]model.GetActivityCommandResult, error) {
	logger, ok := ctx.Value(pkg.LoggerKey).(*zap.SugaredLogger)
	if !ok {
		return nil, fmt.Errorf("failed to get logger for GetByLocationCommand")
	}

	cityKey := fmt.Sprintf("%s:%s", cacheadapter.CITY_CACHE_KEY, gc.Location.City)

	activities, redisErr := gc.Redis.GetMembers(ctx, cityKey)
	if redisErr != nil {
		logger.Infof("redis error %s \n returning from db", redisErr.Error())
		return gc.Repo.GetByLocation(ctx, &gc.Location)
	}

	var activitiesResult []model.GetActivityCommandResult
	for _, activity := range activities {

		var activityObject = model.GetActivityCommandResult{}
		err := json.Unmarshal([]byte(activity), &activityObject)
		if err != nil {
			logger.Infof("parsing erorr returning from db %s", err.Error())
			return gc.Repo.GetByLocation(ctx, &gc.Location)
		}

		if activityObject.Location.City == gc.Location.City {
			activitiesResult = append(activitiesResult, activityObject)
		}
	}

	if activitiesResult != nil {
		return activitiesResult, nil
	}

	return gc.Repo.GetByLocation(ctx, &gc.Location)
}
