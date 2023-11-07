package ports

import "github.com/omerbeden/event-mate/backend/tatooine/modules/profile/domain/model"

type UserProfileRepository interface {
	GetUsers() ([]model.UserProfile, error)
	GetUserById(id uint) (*model.UserProfile, error)
	InsertUser(user *model.UserProfile) (bool, error)
	UpdateUser(user *model.UserProfile) error
	DeleteUserById(id uint) (bool, error)
}

type EventRepository interface {
	GetUserEvent(userId int) ([]model.Event, error)
}

type UserStatRepository interface {
	GetUserStat(userId int) (model.UserProfileStat, error)
}
