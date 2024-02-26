package repositories

import "github.com/omerbeden/event-mate/backend/tatooine/modules/activity/app/domain/model"

type ActivityRepository interface {
	Create(activity model.Activity) (*model.Activity, error)
	GetByID(id int64) (*model.Activity, error)
	GetByLocation(loc *model.Location) ([]model.Activity, error)
	UpdateByID(id int32, activity model.Activity) (bool, error)
	DeleteByID(id int32) (bool, error)
	AddParticipants(activity model.Activity) error
	AddParticipant(activityId int64, user model.User) error
	GetParticipants(activityId int64) ([]model.User, error)
}

type ActivityRulesRepository interface {
	CreateActivityRules(int64, []string) error
	GetActivityRules(int64) ([]string, error)
}

type ActivityFlowRepository interface {
	CreateActivityFlow(int64, []string) error
	GetActivityFlow(int64) ([]string, error)
}

type LocationRepository interface {
	Create(loc *model.Location) (bool, error)
	UpdateByID(loc model.Location) (bool, error)
}
