package commands

import (
	"context"

	"github.com/omerbeden/event-mate/backend/tatooine/modules/activity/app/domain/model"
	repo "github.com/omerbeden/event-mate/backend/tatooine/modules/activity/app/domain/ports/repositories"
)

type UpdateCommand struct {
	Activity *model.Activity
	Repo     repo.ActivityRepository
}

func (uc *UpdateCommand) Handle(ctx context.Context) (bool, error) {
	model := &model.Activity{
		Title:    uc.Activity.Title,
		Category: uc.Activity.Category,
	}

	return uc.Repo.UpdateByID(ctx, uc.Activity.ID, *model)
}
