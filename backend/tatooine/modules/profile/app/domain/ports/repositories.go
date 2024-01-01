package ports

import "github.com/omerbeden/event-mate/backend/tatooine/modules/profile/app/domain/model"

type UserProfileRepository interface {
	GetUsersByAddress(address model.UserProfileAdress) ([]model.UserProfile, error)
	InsertUser(user *model.UserProfile) (bool, error)
	UpdateUser(user *model.UserProfile) error
	DeleteUserById(id uint) (bool, error)
}

type EventRepository interface {
	GetUserEvent(userId int) ([]model.Activity, error)
}

type UserStatRepository interface {
	GetUserStat(userId int) (model.UserProfileStat, error)
}
