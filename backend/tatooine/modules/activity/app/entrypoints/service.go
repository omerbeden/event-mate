package entrypoints

import (
	"context"

	"github.com/omerbeden/event-mate/backend/tatooine/modules/activity/app/domain/commands"
	"github.com/omerbeden/event-mate/backend/tatooine/modules/activity/app/domain/model"
	"github.com/omerbeden/event-mate/backend/tatooine/modules/activity/app/domain/ports/repositories"
	"github.com/omerbeden/event-mate/backend/tatooine/pkg/cache"
	"github.com/omerbeden/event-mate/backend/tatooine/pkg/db"
	"go.uber.org/zap"
)

type ActivityService struct {
	activityRepository      repositories.ActivityRepository
	activityRulesRepository repositories.ActivityRulesRepository
	activityFlowRepository  repositories.ActivityFlowRepository
	locationReposiroy       repositories.LocationRepository
	redisClient             cache.RedisClient
	tx                      db.TransactionManager
	Logger                  *zap.SugaredLogger
}

func NewService(
	activityRepository repositories.ActivityRepository,
	activityRulesRepository repositories.ActivityRulesRepository,
	activityFlowRepository repositories.ActivityFlowRepository,
	locationRepository repositories.LocationRepository,
	redisClient cache.RedisClient,
	tx db.TransactionManager,
	logger *zap.SugaredLogger,
) *ActivityService {
	return &ActivityService{
		activityRepository:      activityRepository,
		activityRulesRepository: activityRulesRepository,
		activityFlowRepository:  activityFlowRepository,
		locationReposiroy:       locationRepository,
		redisClient:             redisClient,
		tx:                      tx,
		Logger:                  logger,
	}
}

func (service ActivityService) CreateActivity(ctx context.Context, activity model.Activity) (bool, error) {

	createCmd := &commands.CreateCommand{
		ActivityRepo:      service.activityRepository,
		LocRepo:           service.locationReposiroy,
		ActivityRulesRepo: service.activityRulesRepository,
		ActivityFlowRepo:  service.activityFlowRepository,
		Activity:          activity,
		Redis:             &service.redisClient,
		Tx:                service.tx,
	}

	createCmdResult, err := createCmd.Handle(ctx)
	if err != nil {
		return false, err
	}

	return createCmdResult, nil

}

func (service ActivityService) AddParticipant(ctx context.Context, participant model.User, activityId int64) error {
	addParticipantCommand := &commands.AddParticipantCommand{
		ActivityRepository: service.activityRepository,
		Redis:              &service.redisClient,
		Participant:        participant,
		ActivityId:         activityId,
	}

	return addParticipantCommand.Handle(ctx)
}

func (service ActivityService) GetParticipants(ctx context.Context, activityId int64) ([]model.User, error) {

	getParticipantsCommand := &commands.GetParticipantsCommand{
		ActivityRepository: service.activityRepository,
		Redis:              &service.redisClient,
		ActivityId:         activityId,
	}

	return getParticipantsCommand.Handle(ctx)

}
func (service ActivityService) GetActivityById(ctx context.Context, activityId int64) (*model.Activity, error) {
	getCommand := &commands.GetByIDCommand{
		Repo:              service.activityRepository,
		ActivityRulesRepo: service.activityRulesRepository,
		ActivityFlowRepo:  service.activityFlowRepository,
		ActivityId:        activityId,
		Redis:             &service.redisClient,
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
		Repo:     service.activityRepository,
		Redis:    &service.redisClient,
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
