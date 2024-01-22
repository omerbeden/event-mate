package repositories

import "github.com/omerbeden/event-mate/backend/tatooine/modules/profile/app/domain/model"

type UserProfileRepository interface {
	GetUsersByAddress(address model.UserProfileAdress) ([]model.UserProfile, error)
	InsertUser(user *model.UserProfile) (*model.UserProfile, error)
	UpdateProfileImage(externalId string, imageUrl string) error
	UpdateProfilePoints(receiverUserName string, point float32) error
	DeleteUser(external_id string) error
	GetAttandedActivities(userId int64) ([]model.Activity, error)
	GetUserProfileStats(userId int64) (*model.UserProfileStat, error)
	GetCurrentUserProfile(externalId string) (*model.UserProfile, error)
	GetUserProfile(username string) (*model.UserProfile, error)
}
