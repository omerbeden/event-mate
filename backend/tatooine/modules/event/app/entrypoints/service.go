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
	EventRepository   repositories.EventRepository
	LocationReposiroy repositories.LocationRepository
	RedisClient       redis.Client
}

func (service EventService) CreateEvent(ctx context.Context, event model.Event) (bool, error) {

	createCmd := &commands.CreateCommand{
		EventRepo: service.EventRepository,
		LocRepo:   service.LocationReposiroy,
		Event:     event,
		Redis:     redisadapter.NewRedisAdapter(&service.RedisClient),
	}

	createCmdResult, err := createCmd.Handle()
	if err != nil {
		return false, err
	}

	return createCmdResult, nil

}

func (service EventService) GetEventById(ctx context.Context, eventId int64) (*model.Event, error) {
	getCommand := &commands.GetByIDCommand{
		Repo:    service.EventRepository,
		EventID: "eventId",
		Redis:   *redisadapter.NewRedisAdapter(&service.RedisClient),
	}

	commandResult, err := getCommand.Handle()
	if err != nil {
		return nil, err
	}

	return commandResult, nil
}

func (service EventService) GetEventsByLocation(ctx context.Context, loc model.Location) ([]model.Event, error) {
	getCommand := &commands.GetByLocationCommand{
		Location: loc,
		Repo:     service.EventRepository,
		Redis:    *redisadapter.NewRedisAdapter(&service.RedisClient),
	}

	commandResult, err := getCommand.Handle()
	if err != nil {
		return nil, err
	}

	return commandResult, nil
}
