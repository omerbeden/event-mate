package repositories

import (
	"context"

	"github.com/omerbeden/event-mate/backend/tatooine/modules/activity/app/domain/model"
)

type ActivityRepository interface {
	Create(context.Context, model.Activity) (*model.Activity, error)
	GetByID(context.Context, int64) (*model.Activity, error)
	GetByLocation(context.Context, *model.Location) ([]model.Activity, error)
	UpdateByID(context.Context, int64, model.Activity) (bool, error)
	DeleteByID(context.Context, int64) (bool, error)
	AddParticipants(context.Context, model.Activity) error
	AddParticipant(context.Context, int64, model.User) error
	GetParticipants(context.Context, int64) ([]model.User, error)
}

type ActivityRulesRepository interface {
	CreateActivityRules(context.Context, int64, []string) error
	GetActivityRules(context.Context, int64) ([]string, error)
}

type ActivityFlowRepository interface {
	CreateActivityFlow(context.Context, int64, []string) error
	GetActivityFlow(context.Context, int64) ([]string, error)
}

type LocationRepository interface {
	Create(context.Context, *model.Location) (bool, error)
	UpdateByID(context.Context, model.Location) (bool, error)
}
