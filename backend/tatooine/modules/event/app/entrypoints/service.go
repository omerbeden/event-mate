package entrypoints

import (
	"context"

	"github.com/go-redis/redis/v8"
	"github.com/omerbeden/event-mate/backend/tatooine/modules/event/app/adapters/redisadapter"
	"github.com/omerbeden/event-mate/backend/tatooine/modules/event/app/domain/commands"
	"github.com/omerbeden/event-mate/backend/tatooine/modules/event/app/domain/model"
	"github.com/omerbeden/event-mate/backend/tatooine/modules/event/app/domain/ports/repositories"
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

	createCmdResult, err := createCmd.Handle()
	if err != nil {
		return nil, err
	}

	return &CreateEventResponse{status: createCmdResult}, nil

}

func (service EventService) GetEventById(ctx context.Context, eventId int64) (*GetEventResponse, error) {
	getCommand := &commands.GetByIDCommand{
		Repo:    service.eventRepository,
		EventID: "eventId",
		Redis:   *redisadapter.NewRedisAdapter(&service.redisClient),
	}

	commandResult, err := getCommand.Handle()
	if err != nil {
		return nil, err
	}

	return &GetEventResponse{*commandResult}, nil
}

func (service EventService) GetEventsByLocation(ctx context.Context, loc model.Location) ([]model.Event, error) {
	getCommand := &commands.GetByLocationCommand{
		Location: loc,
		Repo:     service.eventRepository,
		Redis:    *redisadapter.NewRedisAdapter(&service.redisClient),
	}

	commandResult, err := getCommand.Handle()
	if err != nil {
		return nil, err
	}

	return commandResult, nil
}
