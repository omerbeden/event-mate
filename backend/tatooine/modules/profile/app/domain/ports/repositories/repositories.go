package repositories

import "github.com/omerbeden/event-mate/backend/tatooine/modules/profile/app/domain/model"

type UserProfileRepository interface {
	GetUsersByAddress(address model.UserProfileAdress) ([]model.UserProfile, error)
	InsertUser(user *model.UserProfile) (*model.UserProfile, error)
	UpdateProfileImage(userId int64, imageUrl string) (*model.UserProfile, error)
	DeleteUserById(id int64) error
	GetAttandedActivities(userId int64) ([]model.Activity, error)
}

type EventRepository interface {
	GetUserEvent(userId int) ([]model.Activity, error)
}

type UserStatRepository interface {
	GetUserStat(userId int) (model.UserProfileStat, error)
}
