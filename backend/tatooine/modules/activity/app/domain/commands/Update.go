package commands

import (
	"github.com/omerbeden/event-mate/backend/tatooine/modules/activity/app/domain/model"
	repo "github.com/omerbeden/event-mate/backend/tatooine/modules/activity/app/domain/ports/repositories"
)

type UpdateCommand struct {
	Activity *model.Activity
	Repo     repo.ActivityRepository
}

func (uc *UpdateCommand) Handle() (bool, error) {
	model := &model.Activity{
		Title:    uc.Activity.Title,
		Category: uc.Activity.Category,
	}

	return uc.Repo.UpdateByID(int32(uc.Activity.ID), *model)
}
