package commands

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/omerbeden/event-mate/backend/tatooine/modules/profile/app/domain/model"
	"github.com/omerbeden/event-mate/backend/tatooine/modules/profile/app/domain/ports/repositories"
	"github.com/omerbeden/event-mate/backend/tatooine/pkg/cache"
)

type GetAttandedActivitiesCommand struct {
	Repo   repositories.UserProfileRepository
	Cache  cache.Cache
	UserId int64
}

func (c *GetAttandedActivitiesCommand) Handle(ctx context.Context) ([]model.Activity, error) {

	attandedActivitiesFromCache, err := c.getFromCache(ctx, c.UserId)
	if err != nil || len(attandedActivitiesFromCache) < 1 || attandedActivitiesFromCache == nil {
		fmt.Printf("returning attended activities from db userid : %d\n", c.UserId)
		return c.Repo.GetAttandedActivities(ctx, c.UserId)
	}

	return attandedActivitiesFromCache, nil
}

func (c *GetAttandedActivitiesCommand) getFromCache(ctx context.Context, userId int64) ([]model.Activity, error) {
	cacheKey := fmt.Sprintf("%s:%d", cache.ATTANDED_ACTIVITIES_CACHE_KEY, userId)

	attandedActivitiesStr, err := c.Cache.GetMembers(ctx, cacheKey)
	if err != nil {
		return nil, err
	}

	var attandedActivities []model.Activity
	for _, activityStr := range attandedActivitiesStr {
		var activity model.Activity
		err := json.Unmarshal([]byte(activityStr), &activity)
		if err != nil {
			return nil, fmt.Errorf("parsing erorr  %w", err)

		}

		attandedActivities = append(attandedActivities, activity)
	}

	return attandedActivities, nil
}
