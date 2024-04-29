package commands

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/omerbeden/event-mate/backend/tatooine/modules/profile/app/domain/model"
	"github.com/omerbeden/event-mate/backend/tatooine/modules/profile/app/domain/ports/repositories"
	"github.com/omerbeden/event-mate/backend/tatooine/pkg/cache"
)

type GetCreatedActivitiesCommand struct {
	Repo   repositories.UserProfileRepository
	Cache  cache.Cache
	UserId int64
}

func (cmd GetCreatedActivitiesCommand) Handle(ctx context.Context) ([]model.Activity, error) {

	activities, err := cmd.getFromCache(ctx)
	if err != nil || len(activities) < 1 || activities == nil {
		activities, err := cmd.Repo.GetCreatedActivities(ctx, cmd.UserId)
		if err != nil {
			return nil, err
		}
		return activities, nil
	}

	return activities, nil
}

func (cmd GetCreatedActivitiesCommand) getFromCache(ctx context.Context) ([]model.Activity, error) {
	cacheKey := fmt.Sprintf("%s:%d", cache.CREATED_ACTIVITIES_CACHE_KEY, cmd.UserId)

	createdActivitiesStr, err := cmd.Cache.GetMembers(ctx, cacheKey)
	if err != nil {
		return nil, err
	}

	var createdActivities []model.Activity
	for _, activityStr := range createdActivitiesStr {
		var activity model.Activity
		err := json.Unmarshal([]byte(activityStr), &activity)
		if err != nil {
			return nil, fmt.Errorf("parsing erorr  %w", err)

		}

		createdActivities = append(createdActivities, activity)
	}

	return createdActivities, nil
}
