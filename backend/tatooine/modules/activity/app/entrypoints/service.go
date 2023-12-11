package entrypoints

import (
	"context"
	"fmt"

	"github.com/go-redis/redis/v8"
	"github.com/omerbeden/event-mate/backend/tatooine/modules/activity/app/adapters/redisadapter"
	"github.com/omerbeden/event-mate/backend/tatooine/modules/activity/app/domain/commands"
	"github.com/omerbeden/event-mate/backend/tatooine/modules/activity/app/domain/model"
	"github.com/omerbeden/event-mate/backend/tatooine/modules/activity/app/domain/ports/repositories"
)

type ActivityService struct {
	ActivityRepository repositories.ActivityRepository
	LocationReposiroy  repositories.LocationRepository
	RedisClient        redis.Client
}

func (service ActivityService) CreateActivity(ctx context.Context, activity model.Activity) (bool, error) {

	createCmd := &commands.CreateCommand{
		ActivityRepo: service.ActivityRepository,
		LocRepo:      service.LocationReposiroy,
		Activity:     activity,
		Redis:        redisadapter.NewRedisAdapter(&service.RedisClient),
	}

	createCmdResult, err := createCmd.Handle()
	if err != nil {
		return false, err
	}

	return createCmdResult, nil

}

func (service ActivityService) GetActivityById(ctx context.Context, activityId int64) (*model.Activity, error) {
	getCommand := &commands.GetByIDCommand{
		Repo:       service.ActivityRepository,
		ActivityId: activityId,
		Redis:      *redisadapter.NewRedisAdapter(&service.RedisClient),
	}

	commandResult, err := getCommand.Handle()
	if err != nil {
		return nil, err
	}

	return commandResult, nil
}

func (service ActivityService) GetActivitiesByLocation(ctx context.Context, loc model.Location) ([]model.Activity, error) {
	getCommand := &commands.GetByLocationCommand{
		Location: loc,
		Repo:     service.ActivityRepository,
		Redis:    *redisadapter.NewRedisAdapter(&service.RedisClient),
	}

	commandResult, err := getCommand.Handle()

	fmt.Printf("%+v , %+v", commandResult, nil)
	if err != nil {
		return nil, err
	}
	if commandResult == nil {
		return nil, err
	}

	return commandResult, nil
}
