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
	DeleteUser(ctx context.Context, externalId string) error
	GetAttandedActivities(ctx context.Context, activityId int64) ([]model.Activity, error)
	GetCurrentUserProfile(ctx context.Context, externalId string) (*model.UserProfile, error)
	GetUserProfile(ctx context.Context, username string) (*model.UserProfile, error)
}

type UserProfileAddressRepository interface {
	Insert(ctx context.Context, tx db.Tx, address model.UserProfileAdress) error
}
type UserProfileStatRepository interface {
	Insert(ctx context.Context, tx db.Tx, stat model.UserProfileStat) error
	EvaluateUser(ctx context.Context, eval model.UserEvaluation) error
}
