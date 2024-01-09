package commands

import (
	"encoding/json"
	"fmt"

	"github.com/omerbeden/event-mate/backend/tatooine/modules/profile/app/adapters/cachedapter"
	"github.com/omerbeden/event-mate/backend/tatooine/modules/profile/app/domain/model"
	"github.com/omerbeden/event-mate/backend/tatooine/modules/profile/app/domain/ports/repositories"
)

type GetAttandedActivitiesCommand struct {
	Repo   repositories.UserProfileRepository
	Cache  cachedapter.Cache
	UserId int64
}

func (c *GetAttandedActivitiesCommand) Handle() ([]model.Activity, error) {

	attandedActivitiesFromCache, err := c.getFromCache(c.UserId)
	if err != nil || len(attandedActivitiesFromCache) < 1 || attandedActivitiesFromCache == nil {
		fmt.Printf("returning attended activities from db")
		c.Repo.GetAttandedActivities(c.UserId)
	}

	return attandedActivitiesFromCache, nil
}

func (c *GetAttandedActivitiesCommand) getFromCache(userId int64) ([]model.Activity, error) {
	attandedActivitiesStr, err := c.Cache.GetAttandedActivities(userId)
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
