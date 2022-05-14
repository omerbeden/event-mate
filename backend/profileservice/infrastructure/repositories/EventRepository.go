package repositories

import (
	"github.com/omerbeden/event-mate/backend/profileservice/core"
	"github.com/omerbeden/event-mate/backend/profileservice/infrastructure/database"
	"gorm.io/gorm"
)

type EventRepo struct{}

func NewEventRepo() *EventRepo {
	return &EventRepo{}

}

func (r *EventRepo) GetUserEvent(userId int) ([]core.Event, error) {
	db := database.NewConnPG()
	var events []core.Event
	if err := db.First(&events, userId).Error; err != nil {
		return events, gorm.ErrRecordNotFound
	}
	return events, nil
}
