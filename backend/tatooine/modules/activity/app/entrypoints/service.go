package entrypoints

import (
	"context"

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

func NewService(
	activityRepository repositories.ActivityRepository,
	locationRepository repositories.LocationRepository,
	redisClient redis.Client,
) *ActivityService {
	return &ActivityService{
		ActivityRepository: activityRepository,
		LocationReposiroy:  locationRepository,
		RedisClient:        redisClient,
	}
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

func (service ActivityService) AddParticipant(participant model.User, activityId int64) error {
	addParticipantCommand := &commands.AddParticipantCommand{
		ActivityRepository: service.ActivityRepository,
		Redis:              redisadapter.NewRedisAdapter(&service.RedisClient),
		Participant:        participant,
		ActivityId:         activityId,
	}

	return addParticipantCommand.Handle()
}

func (service ActivityService) GetParticipants(activityId int64) ([]model.User, error) {

	getParticipantsCommand := &commands.GetParticipantsCommand{
		ActivityRepository: service.ActivityRepository,
		Redis:              redisadapter.NewRedisAdapter(&service.RedisClient),
		ActivityId:         activityId,
	}

	return getParticipantsCommand.Handle()

}
func (service ActivityService) GetActivityById(ctx context.Context, activityId int64) (*model.Activity, error) {
	getCommand := &commands.GetByIDCommand{
		Repo:       service.ActivityRepository,
		ActivityId: activityId,
		Redis:      redisadapter.NewRedisAdapter(&service.RedisClient),
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
		Redis:    redisadapter.NewRedisAdapter(&service.RedisClient),
	}

	activities, err := getCommand.Handle()

	for i := range activities {
		activities[i].Participants, err = service.GetParticipants(activities[i].ID)
		if err != nil {
			return nil, err
		}
	}

	if err != nil {
		return nil, err
	}
	if activities == nil {
		return nil, err
	}

	return activities, nil
}
