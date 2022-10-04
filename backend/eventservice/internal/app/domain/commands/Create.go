package commands

import (
	"github.com/omerbeden/event-mate/backend/eventservice/internal/app/domain/model"
	"github.com/omerbeden/event-mate/backend/eventservice/internal/app/domain/ports/repo"
	"github.com/omerbeden/event-mate/backend/eventservice/internal/infra/grpc/pb"
)

type CreateCommand struct {
	Event *pb.Event
	Repo  repo.Repository
}

func (CreateCommand *CreateCommand) Handle() (bool, error) {
	model := &model.Event{
		Title:    CreateCommand.Event.Title,
		Category: CreateCommand.Event.Category,
	}

	return CreateCommand.Repo.CreateEvent(*model)

}

type TEstCommand struct{}

func (CreateCommand *TEstCommand) Handle() (bool, error) {
	return true, nil

}
