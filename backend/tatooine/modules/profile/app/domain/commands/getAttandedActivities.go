package commands

import (
	"github.com/omerbeden/event-mate/backend/tatooine/modules/profile/app/domain/model"
	"github.com/omerbeden/event-mate/backend/tatooine/modules/profile/app/domain/ports"
)

type GetAttandedActivitiesCommand struct {
	Repo   ports.UserProfileRepository
	UserId int64
}

func (c *GetAttandedActivitiesCommand) Handle() ([]model.Activity, error) {
	return c.Repo.GetAttandedActivities(c.UserId)
}
