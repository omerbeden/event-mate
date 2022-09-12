package repo

import "github.com/omerbeden/event-mate/backend/eventservice/domain/model"

type Repository interface {
	CreateEvent(event model.Event) error
	GetEvent() (error, event model.Event)
}
