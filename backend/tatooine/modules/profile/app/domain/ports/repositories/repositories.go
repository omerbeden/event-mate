package repositories

import (
	"context"

	"github.com/omerbeden/event-mate/backend/tatooine/modules/profile/app/domain/model"
	"github.com/omerbeden/event-mate/backend/tatooine/pkg/db"
)

type UserProfileRepository interface {
	GetUsersByAddress(ctx context.Context, address model.UserProfileAdress) ([]model.UserProfile, error)
	Insert(ctx context.Context, tx db.Tx, profile *model.UserProfile) (*model.UserProfile, error)
	UpdateProfileImage(ctx context.Context, externalId string, imageUrl string) error
	UpdateVerification(ctx context.Context, externalId string, isVerified bool) error
	DeleteUser(ctx context.Context, externalId string) error
	GetAttandedActivities(ctx context.Context, activityId int64) ([]model.Activity, error)
	GetCreatedActivities(ctx context.Context, userId int64) ([]model.Activity, error)
	GetCurrentUserProfile(ctx context.Context, externalId string) (*model.UserProfile, error)
	GetUserProfile(ctx context.Context, username string) (*model.UserProfile, error)
	GetUserProfileById(ctx context.Context, id int64) (*model.UserProfile, error)
	GetId(ctx context.Context, externalId string) (int64, error)
}

type UserProfileAddressRepository interface {
	Insert(ctx context.Context, tx db.Tx, address model.UserProfileAdress) error
}
type UserProfileStatRepository interface {
	EvaluateUser(ctx context.Context, eval model.UserEvaluation) error
	GetEvaluations(ctx context.Context, userId int64) ([]model.GetUserEvaluations, error)
}

type ProfileBadgeRepository interface {
	Insert(ctx context.Context, tx db.Tx, badge *model.ProfileBadge) error
	GetBadges(ctx context.Context, profileId int64) (map[int64]*model.ProfileBadge, error)
}
