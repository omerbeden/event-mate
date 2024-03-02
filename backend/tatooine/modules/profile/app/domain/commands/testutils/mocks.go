package testutils

import (
	"context"

	"github.com/omerbeden/event-mate/backend/tatooine/modules/profile/app/domain/model"
	"github.com/omerbeden/event-mate/backend/tatooine/modules/profile/app/domain/ports/repositories"
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
	InsertFunc                func(ctx context.Context, profile *model.UserProfile) (*model.UserProfile, error)
	UpdateProfileImageFunc    func(ctx context.Context, externalId string, imageUrl string) error
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

func (m *MockUserRepository) Insert(ctx context.Context, profile *model.UserProfile) (*model.UserProfile, error) {
	if m.InsertFunc != nil {
		return m.InsertFunc(ctx, profile)
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
	InsertFunc func(ctx context.Context, address model.UserProfileAdress) error
}

func (m *MockAddressRepository) Insert(ctx context.Context, address model.UserProfileAdress) error {
	if m.InsertFunc != nil {
		return m.InsertFunc(ctx, address)
	}

	return nil
}

var _ repositories.UserProfileAddressRepository = (*MockAddressRepository)(nil)

type MockStatRepository struct {
	InsertFunc              func(ctx context.Context, stat model.UserProfileStat) error
	UpdatePRofilePointsFunc func(context.Context, string, float32) error
}

var _ repositories.UserProfileStatRepository = (*MockStatRepository)(nil)

func (m *MockStatRepository) Insert(ctx context.Context, stat model.UserProfileStat) error {
	if m.InsertFunc != nil {
		return m.InsertFunc(ctx, stat)
	}
	return nil
}

func (m *MockStatRepository) UpdateProfilePoints(ctx context.Context, receiverId string, point float32) error {
	if m.InsertFunc != nil {
		return m.UpdateProfilePoints(ctx, receiverId, point)
	}
	return nil
}
