package adapters

import (
	"context"

	"github.com/omerbeden/event-mate/backend/eventservice/internal/app/domain/model"
	"gorm.io/gorm"
)

type EventRepository struct {
	DB *gorm.DB
}

func (r *EventRepository) CreateEvent(event model.Event) (bool, error) {

	if err := r.DB.Create(event).Error; err != nil {
		r.DB.Logger.Error(context.TODO(), "Error occurred while Creating Event")
		return false, err
	}

	return true, nil
}
func (r *EventRepository) GetEventByID(id int32) (model.Event, error) {
	var event model.Event
	if err := r.DB.First(&event, id).Error; err != nil {
		r.DB.Logger.Error(context.TODO(), "Error occurred while Getting Event")
		return event, err
	}
	return event, nil
}

func (r *EventRepository) UpdateEvent(event model.Event) (bool, error) {
	if err := r.DB.Save(&event).Error; err != nil {
		r.DB.Logger.Error(context.TODO(), "Error occurred while updating Event")
		return false, err
	}
	return true, nil
}

func (r *EventRepository) DeleteEventByID(id int32) (bool, error) {
	var event model.Event
	if err := r.DB.First(&event, id).Error; err != nil {
		r.DB.Logger.Error(context.TODO(), "Error occurred while deleting Event, event didn't find")
		return false, err
	}
	if err := r.DB.Delete(&event).Error; err != nil {
		r.DB.Logger.Error(context.TODO(), "Error occurred while deleting Event")
		return false, err
	}
	return true, nil
}
