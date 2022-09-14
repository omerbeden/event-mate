package adapters

import (
	"github.com/omerbeden/event-mate/backend/eventservice/internal/app/adapters/database"
	"github.com/omerbeden/event-mate/backend/eventservice/internal/app/domain/model"
)

type EventRepository struct {
	db database.PostgresAdapter
}

func (r *EventRepository) CreateEvent(event model.Event) error {
	db := r.db.NewConnection()
	if err := db.Create(event).Error; err != nil {
		db.Logger.Error(nil, "Error occurred while Creating Event")
		return err
	}

	return nil
}
func (r *EventRepository) GetEvent(id int32) (model.Event, error) {
	db := r.db.NewConnection()
	var event model.Event
	if err := db.First(&event, id).Error; err != nil {
		db.Logger.Error(nil, "Error occurred while Getting Event")
		return event, err
	}
	return event, nil
}
