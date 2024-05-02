package testutils

import (
	"context"

	"github.com/omerbeden/event-mate/backend/tatooine/modules/profile/app/adapters/postgresadapter/testutils"
	"github.com/omerbeden/event-mate/backend/tatooine/modules/profile/app/domain/model"
	"github.com/omerbeden/event-mate/backend/tatooine/modules/profile/app/domain/ports/repositories"
	"github.com/omerbeden/event-mate/backend/tatooine/pkg/db"
)

type MockUserRepository struct {
	AttandedActivities        []model.Activity
	Profile                   model.UserProfile
	Profiles                  []model.UserProfile
	DeleteUserFunc            func(context.Context, string) error
	GetAttandedActivitiesFunc func(context.Context, int64) ([]model.Activity, error)
	GetCurrentUserProfileFunc func(context.Context, string) (*model.UserProfile, error)
	GetUserProfileFunc        func(ctx context.Context, username string) (*model.UserProfile, error)
	GetUsersByAddressFunc     func(context.Context, model.UserProfileAdress) ([]model.UserProfile, error)
	InsertFunc                func(ctx context.Context, tx db.Tx, profile *model.UserProfile) (*model.UserProfile, error)
	UpdateProfileImageFunc    func(ctx context.Context, externalId string, imageUrl string) error
	UpdateVerificationFunc    func(ctx context.Context, externalId string, isVerified bool) error
	GetIdFunc                 func(ctx context.Context, externalId string) (int64, error)
}

// GetCreatedActivities implements repositories.UserProfileRepository.
func (m *MockUserRepository) GetCreatedActivities(ctx context.Context, userId int64) ([]model.Activity, error) {
	panic("unimplemented")
}

// GetUserProfileById implements repositories.UserProfileRepository.
func (m *MockUserRepository) GetUserProfileById(ctx context.Context, id int64) (*model.UserProfile, error) {
	panic("unimplemented")
}

func (m *MockUserRepository) GetId(ctx context.Context, externalId string) (int64, error) {
	if m.GetIdFunc != nil {
		return m.GetIdFunc(ctx, externalId)
	}
	return 0, nil
}

func (m *MockUserRepository) UpdateVerification(ctx context.Context, externalId string, isVerified bool) error {
	if m.UpdateVerificationFunc != nil {
		return m.UpdateVerificationFunc(ctx, externalId, isVerified)
	}

	return nil
}

func (m *MockUserRepository) DeleteUser(ctx context.Context, externalId string) error {
	if m.DeleteUserFunc != nil {
		return m.DeleteUserFunc(ctx, externalId)
	}

	return nil
}

func (m *MockUserRepository) GetAttandedActivities(ctx context.Context, activityId int64) ([]model.Activity, error) {
	if m.GetAttandedActivitiesFunc != nil {
		return m.GetAttandedActivitiesFunc(ctx, activityId)
	}
	return m.AttandedActivities, nil
}

func (m *MockUserRepository) GetCurrentUserProfile(ctx context.Context, externalId string) (*model.UserProfile, error) {
	if m.GetCurrentUserProfileFunc != nil {
		return m.GetCurrentUserProfileFunc(ctx, externalId)
	}
	return &m.Profile, nil
}

func (m *MockUserRepository) GetUserProfile(ctx context.Context, username string) (*model.UserProfile, error) {
	if m.GetUserProfileFunc != nil {
		return m.GetUserProfileFunc(ctx, username)
	}

	return &m.Profile, nil
}

func (m *MockUserRepository) GetUsersByAddress(ctx context.Context, address model.UserProfileAdress) ([]model.UserProfile, error) {
	if m.GetUsersByAddressFunc != nil {
		return m.GetUsersByAddressFunc(ctx, address)
	}

	return m.Profiles, nil
}

func (m *MockUserRepository) Insert(ctx context.Context, tx db.Tx, profile *model.UserProfile) (*model.UserProfile, error) {
	if m.InsertFunc != nil {
		return m.InsertFunc(ctx, tx, profile)
	}

	return &m.Profile, nil
}

func (m *MockUserRepository) UpdateProfileImage(ctx context.Context, externalId string, imageURL string) error {
	if m.UpdateProfileImageFunc != nil {
		return m.UpdateProfileImageFunc(ctx, externalId, imageURL)
	}
	return nil
}

var _ repositories.UserProfileRepository = (*MockUserRepository)(nil)

type MockAddressRepository struct {
	InsertFunc func(ctx context.Context, tx db.Tx, address model.UserProfileAdress) error
}

func (m *MockAddressRepository) Insert(ctx context.Context, tx db.Tx, address model.UserProfileAdress) error {
	if m.InsertFunc != nil {
		return m.InsertFunc(ctx, tx, address)
	}

	return nil
}

var _ repositories.UserProfileAddressRepository = (*MockAddressRepository)(nil)

type MockStatRepository struct {
	InsertFunc       func(ctx context.Context, tx db.Tx, stat model.UserProfileStat) error
	EvaluateUserFunc func(ctx context.Context, eval model.UserEvaluation) error
}

// GetEvaluations implements repositories.UserProfileStatRepository.
func (m *MockStatRepository) GetEvaluations(ctx context.Context, userId int64) ([]model.GetUserEvaluations, error) {
	panic("unimplemented")
}

var _ repositories.UserProfileStatRepository = (*MockStatRepository)(nil)

func (m *MockStatRepository) Insert(ctx context.Context, tx db.Tx, stat model.UserProfileStat) error {
	if m.InsertFunc != nil {
		return m.InsertFunc(ctx, tx, stat)
	}
	return nil
}

func (m *MockStatRepository) EvaluateUser(ctx context.Context, eval model.UserEvaluation) error {
	if m.EvaluateUserFunc != nil {
		return m.EvaluateUserFunc(ctx, eval)
	}
	return nil
}

type MockTxnManager struct{}

func (m *MockTxnManager) Begin(ctx context.Context) (db.Tx, error) {
	return &testutils.MockTx{}, nil
}
