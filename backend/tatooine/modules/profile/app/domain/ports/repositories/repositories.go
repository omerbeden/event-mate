package repositories

import "github.com/omerbeden/event-mate/backend/tatooine/modules/profile/app/domain/model"

type UserProfileRepository interface {
	GetUsersByAddress(address model.UserProfileAdress) ([]model.UserProfile, error)
	InsertUser(user *model.UserProfile) (*model.UserProfile, error)
	UpdateProfileImage(userId int64, imageUrl string) (*model.UserProfile, error)
	UpdateProfilePoints(userId int64, point float32) error
	DeleteUserById(id int64) error
	GetAttandedActivities(userId int64) ([]model.Activity, error)
	GetUserProfileStats(userId int64) (*model.UserProfileStat, error)
	GetUserProfile(userId int64) (*model.UserProfile, error)
}
