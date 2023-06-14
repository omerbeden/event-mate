package adapters

import (
	"context"

	"github.com/omerbeden/event-mate/backend/tatooine/modules/event/app/domain/model"
	"gorm.io/gorm"
)

type EventRepository struct {
	DB *gorm.DB
}

func (r *EventRepository) CreateEvent(event model.Event) (bool, error) {

	if err := r.DB.Create(&event).Error; err != nil {
		r.DB.Logger.Error(context.TODO(), "Error occurred while Creating Event")
		return false, err
	}

	return true, nil
}
func (r *EventRepository) GetEventByID(id int32) (model.Event, error) {
	var event model.Event
	if err := r.DB.First(&event, id).Error; err != nil {
		r.DB.Logger.Error(context.TODO(), "Error occurred while Getting Event by ID")
		return event, err
	}
	return event, nil
}
func (r *EventRepository) GetEventByLocation(loc *model.Location) ([]model.Event, error) {
	var event []model.Event
	if err := r.DB.Find(&event, loc.City).Error; err != nil {
		r.DB.Logger.Error(context.TODO(), "Error occurred while Getting Event by Location")
		return event, err
	}
	return event, nil
}

func (r *EventRepository) UpdateEventByID(id int32, event model.Event) (bool, error) {

	var eventTobeUpdated model.Event

	if err := r.DB.First(&eventTobeUpdated, id).Error; err != nil {
		return false, err
	}

	eventTobeUpdated.Category = event.Category
	eventTobeUpdated.Title = event.Title

	if err := r.DB.Save(&eventTobeUpdated).Error; err != nil {
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
