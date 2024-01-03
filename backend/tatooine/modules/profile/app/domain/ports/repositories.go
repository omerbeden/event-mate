package ports

import "github.com/omerbeden/event-mate/backend/tatooine/modules/profile/app/domain/model"

type UserProfileRepository interface {
	GetUsersByAddress(address model.UserProfileAdress) ([]model.UserProfile, error)
	InsertUser(user *model.UserProfile) (bool, error)
	UpdateProfileImage(imageUrl string) error
	DeleteUserById(id uint) error
}

type EventRepository interface {
	GetUserEvent(userId int) ([]model.Activity, error)
}

type UserStatRepository interface {
	GetUserStat(userId int) (model.UserProfileStat, error)
}
