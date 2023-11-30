package entrypoints

import (
	"context"

	"github.com/go-redis/redis/v8"
	"github.com/omerbeden/event-mate/backend/tatooine/modules/event/app/adapters/redisadapter"
	"github.com/omerbeden/event-mate/backend/tatooine/modules/event/app/domain/commands"
	"github.com/omerbeden/event-mate/backend/tatooine/modules/event/app/domain/model"
	"github.com/omerbeden/event-mate/backend/tatooine/modules/event/app/domain/ports/repositories"
	"github.com/omerbeden/event-mate/backend/tatooine/pkg/command"
)

type EventService struct {
	eventRepository   repositories.EventRepository
	locationReposiroy repositories.LocationRepository
	redisClient       redis.Client
}

type CreateEventRequest struct {
}

type CreateEventResponse struct {
	status bool
}

type GetEventResponse struct {
	model.Event
}

func (service EventService) CreateEvent(ctx context.Context, req CreateEventRequest) (*CreateEventResponse, error) {

	event := model.Event{}

	createCmd := &commands.CreateCommand{
		EventRepo: service.eventRepository,
		LocRepo:   service.locationReposiroy,
		Event:     event,
		Redis:     redisadapter.NewRedisAdapter(&service.redisClient),
	}

	createCmdResult, err := command.HandleCommand[bool](createCmd)
	if err != nil {
		return nil, err
	}

	return &CreateEventResponse{status: createCmdResult}, nil

}

func (service EventService) GetEventById(ctx context.Context, eventId int64) (*GetEventResponse, error) {
	//TODO: refactor here ,split get command by getbyid and getbylocation
	getCommand := &commands.GetCommand{
		Repo:      service.eventRepository,
		EventID:   "eventId",
		EventCity: "refactor",
		Redis:     *redisadapter.NewRedisAdapter(&service.redisClient),
	}

	commandResult, err := command.HandleCommand[*model.Event](getCommand)
	if err != nil {
		return nil, err
	}

	return &GetEventResponse{*commandResult}, nil
}
