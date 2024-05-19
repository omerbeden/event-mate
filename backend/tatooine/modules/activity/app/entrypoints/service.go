package entrypoints

import (
	"context"

	"github.com/omerbeden/event-mate/backend/tatooine/modules/activity/app/domain/commands"
	"github.com/omerbeden/event-mate/backend/tatooine/modules/activity/app/domain/model"
	"github.com/omerbeden/event-mate/backend/tatooine/modules/activity/app/domain/ports/repositories"
	notifierModel "github.com/omerbeden/event-mate/backend/tatooine/modules/notifier/app/domain/model"
	"github.com/omerbeden/event-mate/backend/tatooine/modules/notifier/app/entrypoints"
	"github.com/omerbeden/event-mate/backend/tatooine/pkg/cache"
	"github.com/omerbeden/event-mate/backend/tatooine/pkg/db"
)

type ActivityService struct {
	activityRepository      repositories.ActivityRepository
	activityRulesRepository repositories.ActivityRulesRepository
	activityFlowRepository  repositories.ActivityFlowRepository
	locationReposiroy       repositories.LocationRepository
	redisClient             cache.RedisClient
	tx                      db.TransactionManager
}

func NewService(
	activityRepository repositories.ActivityRepository,
	activityRulesRepository repositories.ActivityRulesRepository,
	activityFlowRepository repositories.ActivityFlowRepository,
	locationRepository repositories.LocationRepository,
	redisClient cache.RedisClient,
	tx db.TransactionManager,

) *ActivityService {
	return &ActivityService{
		activityRepository:      activityRepository,
		activityRulesRepository: activityRulesRepository,
		activityFlowRepository:  activityFlowRepository,
		locationReposiroy:       locationRepository,
		redisClient:             redisClient,
		tx:                      tx,
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

func (service ActivityService) SendAttendRequest(ctx context.Context, request model.AttendRequest) error {
	// validate , record request in case of FCM error

	cmd := commands.StoreAttendRequest{
		ActivityRepository: service.activityRepository,
		Request:            request,
	}

	err := cmd.Handle(ctx)
	if err != nil {
		return err
	}

	// send notification to user B
	notifierS := entrypoints.NewNotifierService(request.ReceiverId)
	_, err = notifierS.Send(&notifierModel.PushMessage{
		Title: "test first message",
		Body:  "test first message body",
	})

	if err != nil {
		return err
	}

	return nil

}

func (service ActivityService) GetParticipants(ctx context.Context, activityId int64) ([]model.User, error) {

	getParticipantsCommand := &commands.GetParticipantsCommand{
		ActivityRepository: service.activityRepository,
		Redis:              &service.redisClient,
		ActivityId:         activityId,
	}

	return getParticipantsCommand.Handle(ctx)

}
func (service ActivityService) GetActivityDetail(ctx context.Context, activityId int64) (*model.ActivityDetail, error) {
	getCommand := &commands.GetByIDCommand{
		Repo:              service.activityRepository,
		ActivityRulesRepo: service.activityRulesRepository,
		ActivityFlowRepo:  service.activityFlowRepository,
		ActivityId:        activityId,
		Redis:             &service.redisClient,
	}

	activity, err := getCommand.Handle(ctx)
	if err != nil {
		return nil, err
	}

	participants, err := service.GetParticipants(ctx, activityId)
	activity.Participants = participants

	if err != nil {
		return nil, err
	}

	return activity, nil
}

func (service ActivityService) GetActivitiesByLocation(ctx context.Context, loc model.Location) ([]model.GetActivityCommandResult, error) {
	getCommand := &commands.GetByLocationCommand{
		Location: loc,
		Repo:     service.activityRepository,
		Redis:    &service.redisClient,
	}

	activities, err := getCommand.Handle(ctx)

	if err != nil {
		return nil, err
	}

	return activities, nil
}
