package repo

import "github.com/omerbeden/event-mate/backend/eventservice/internal/app/domain/model"

type Repository interface {
	CreateEvent(event model.Event) (bool, error)
	GetEventByID(id int32) (model.Event, error)
}
