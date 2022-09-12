package adapters

import "github.com/omerbeden/event-mate/backend/eventservice/internal/app/domain/model"

type Repositories struct{}

func (r *Repositories) CreateEvent(event model.Event) error {
	return nil
}
func (r *Repositories) GetEvent() (model.Event, error) {
	return model.Event{}, nil
}
