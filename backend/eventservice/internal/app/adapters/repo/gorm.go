package adapters

import "github.com/omerbeden/event-mate/backend/eventservice/internal/app/domain/model"

type EventRepository struct{}

func (r *EventRepository) CreateEvent(event model.Event) error {
	return nil
}
func (r *EventRepository) GetEvent() (model.Event, error) {
	return model.Event{}, nil
}
