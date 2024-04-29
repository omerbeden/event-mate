package repositories

import (
	"context"

	"github.com/omerbeden/event-mate/backend/tatooine/modules/activity/app/domain/model"
	"github.com/omerbeden/event-mate/backend/tatooine/pkg/db"
)

type ActivityRepository interface {
	Create(ctx context.Context, tx db.Tx, activity model.Activity) (*model.Activity, error)
	GetByID(ctx context.Context, activityId int64) (*model.Activity, error)
	GetByLocation(ctx context.Context, location *model.Location) ([]model.GetActivityCommandResult, error)
	UpdateByID(ctx context.Context, activityId int64, activity model.Activity) (bool, error)
	DeleteByID(ctx context.Context, activityId int64) (bool, error)
	AddParticipants(ctx context.Context, activityId int64, participants []model.User) error
	AddParticipant(ctx context.Context, activityId int64, participant model.User) error
	GetParticipants(ctx context.Context, activityId int64) ([]model.User, error)
}

type ActivityRulesRepository interface {
	CreateActivityRules(ctx context.Context, tx db.Tx, activityId int64, rules []string) error
	GetActivityRules(ctx context.Context, activityId int64) ([]string, error)
}

type ActivityFlowRepository interface {
	CreateActivityFlow(ctx context.Context, tx db.Tx, activityId int64, flows []string) error
	GetActivityFlow(ctx context.Context, activityId int64) ([]string, error)
}

type LocationRepository interface {
	Create(ctx context.Context, tx db.Tx, location *model.Location) (bool, error)
	UpdateByID(ctx context.Context, location model.Location) (bool, error)
}
