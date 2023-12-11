package commands

import (
	"github.com/omerbeden/event-mate/backend/tatooine/modules/activity/app/domain/model"
	repo "github.com/omerbeden/event-mate/backend/tatooine/modules/activity/app/domain/ports/repositories"
)

type UpdateCommand struct {
	Event *model.Event
	Repo  repo.EventRepository
}

func (uc *UpdateCommand) Handle() (bool, error) {
	model := &model.Event{
		Title:    uc.Event.Title,
		Category: uc.Event.Category,
	}

	return uc.Repo.UpdateByID(int32(uc.Event.ID), *model)
}
