package repositories

import "github.com/omerbeden/event-mate/backend/tatooine/modules/event/app/domain/model"

type EventRepository interface {
	Create(event model.Event) (*model.Event, error)
	GetByID(id int64) (*model.Event, error)
	GetByLocation(loc *model.Location) ([]model.Event, error)
	UpdateByID(id int32, event model.Event) (bool, error)
	DeleteByID(id int32) (bool, error)
	AddParticipants(event model.Event) error
	AddParticipant(eventId int64, user model.User) error
}

type LocationRepository interface {
	Create(loc *model.Location) (bool, error)
	UpdateByID(loc model.Location) (bool, error)
}
