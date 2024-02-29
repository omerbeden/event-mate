package entrypoints

import (
	"context"

	"github.com/omerbeden/event-mate/backend/tatooine/modules/activity/app/domain/commands"
	"github.com/omerbeden/event-mate/backend/tatooine/modules/activity/app/domain/model"
	"github.com/omerbeden/event-mate/backend/tatooine/modules/activity/app/domain/ports/repositories"
	"github.com/omerbeden/event-mate/backend/tatooine/pkg/cache"
)

type ActivityService struct {
	ActivityRepository      repositories.ActivityRepository
	ActivityRulesRepository repositories.ActivityRulesRepository
	ActivityFlowRepository  repositories.ActivityFlowRepository
	LocationReposiroy       repositories.LocationRepository
	RedisClient             cache.RedisClient
}

func NewService(
	activityRepository repositories.ActivityRepository,
	activityRulesRepository repositories.ActivityRulesRepository,
	activityFlowRepository repositories.ActivityFlowRepository,
	locationRepository repositories.LocationRepository,
	redisClient cache.RedisClient,
) *ActivityService {
	return &ActivityService{
		ActivityRepository:      activityRepository,
		ActivityRulesRepository: activityRulesRepository,
		ActivityFlowRepository:  activityFlowRepository,
		LocationReposiroy:       locationRepository,
		RedisClient:             redisClient,
	}
}

func (service ActivityService) CreateActivity(ctx context.Context, activity model.Activity) (bool, error) {

	createCmd := &commands.CreateCommand{
		ActivityRepo:      service.ActivityRepository,
		LocRepo:           service.LocationReposiroy,
		ActivityRulesRepo: service.ActivityRulesRepository,
		ActivityFlowRepo:  service.ActivityFlowRepository,
		Activity:          activity,
		Redis:             &service.RedisClient,
	}

	createCmdResult, err := createCmd.Handle(ctx)
	if err != nil {
		return false, err
	}

	return createCmdResult, nil

}

func (service ActivityService) AddParticipant(ctx context.Context, participant model.User, activityId int64) error {
	addParticipantCommand := &commands.AddParticipantCommand{
		ActivityRepository: service.ActivityRepository,
		Redis:              &service.RedisClient,
		Participant:        participant,
		ActivityId:         activityId,
	}

	return addParticipantCommand.Handle(ctx)
}

func (service ActivityService) GetParticipants(ctx context.Context, activityId int64) ([]model.User, error) {

	getParticipantsCommand := &commands.GetParticipantsCommand{
		ActivityRepository: service.ActivityRepository,
		Redis:              &service.RedisClient,
		ActivityId:         activityId,
	}

	return getParticipantsCommand.Handle(ctx)

}
func (service ActivityService) GetActivityById(ctx context.Context, activityId int64) (*model.Activity, error) {
	getCommand := &commands.GetByIDCommand{
		Repo:              service.ActivityRepository,
		ActivityRulesRepo: service.ActivityRulesRepository,
		ActivityFlowRepo:  service.ActivityFlowRepository,
		ActivityId:        activityId,
		Redis:             &service.RedisClient,
	}

	commandResult, err := getCommand.Handle(ctx)

	if err != nil {
		return nil, err
	}

	return commandResult, nil
}

func (service ActivityService) GetActivitiesByLocation(ctx context.Context, loc model.Location) ([]model.Activity, error) {
	getCommand := &commands.GetByLocationCommand{
		Location: loc,
		Repo:     service.ActivityRepository,
		Redis:    &service.RedisClient,
	}

	activities, err := getCommand.Handle(ctx)

	for i := range activities {
		activities[i].Participants, err = service.GetParticipants(ctx, activities[i].ID)
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
