package repositories

import (
	"context"

	"github.com/omerbeden/event-mate/backend/tatooine/modules/profile/app/domain/model"
)

type UserProfileRepository interface {
	GetUsersByAddress(context.Context, model.UserProfileAdress) ([]model.UserProfile, error)
	InsertUser(context.Context, *model.UserProfile) (*model.UserProfile, error)
	UpdateProfileImage(context.Context, string, string) error
	UpdateProfilePoints(context.Context, string, float32) error
	DeleteUser(context.Context, string) error
	GetAttandedActivities(context.Context, int64) ([]model.Activity, error)
	GetUserProfileStats(context.Context, int64) (*model.UserProfileStat, error)
	GetCurrentUserProfile(context.Context, string) (*model.UserProfile, error)
	GetUserProfile(context.Context, string) (*model.UserProfile, error)
}
