package commands

import (
	"strconv"

	"github.com/omerbeden/event-mate/backend/tatooine/modules/event/app/domain/model"
	"github.com/omerbeden/event-mate/backend/tatooine/modules/event/app/domain/ports/repo"
	"github.com/omerbeden/event-mate/backend/tatooine/modules/event/infra/grpc/pb"
)

type UpdateCommand struct {
	Event *pb.Event
	Repo  repo.EventRepository
}

func (uc *UpdateCommand) Handle() (bool, error) {
	model := &model.Event{
		Title:    uc.Event.Title,
		Category: uc.Event.Category,
	}

	intID, _ := strconv.Atoi(uc.Event.GetId())
	return uc.Repo.UpdateByID(int32(intID), *model)
}
