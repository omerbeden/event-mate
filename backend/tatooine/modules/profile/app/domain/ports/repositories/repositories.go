package repositories

import (
	"context"

	"github.com/omerbeden/event-mate/backend/tatooine/modules/profile/app/domain/model"
)

type UserProfileRepository interface {
	GetUsersByAddress(context.Context, model.UserProfileAdress) ([]model.UserProfile, error)
	Insert(context.Context, *model.UserProfile) (*model.UserProfile, error)
	UpdateProfileImage(context.Context, string, string) error
	UpdateProfilePoints(context.Context, string, float32) error
	DeleteUser(context.Context, string) error
	GetAttandedActivities(context.Context, int64) ([]model.Activity, error)
	GetCurrentUserProfile(context.Context, string) (*model.UserProfile, error)
	GetUserProfile(context.Context, string) (*model.UserProfile, error)
}

type UserProfileAddressRepository interface {
	Insert(context.Context, int64, model.UserProfileAdress) error
}
type UserProfileStatRepository interface {
	Insert(ctx context.Context, user model.UserProfile) error
}
