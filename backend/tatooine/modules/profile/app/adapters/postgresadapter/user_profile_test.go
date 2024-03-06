package postgresadapter_test

import (
	"context"
	"errors"
	"testing"
	"time"

	"github.com/omerbeden/event-mate/backend/tatooine/modules/profile/app/adapters/postgresadapter"
	"github.com/omerbeden/event-mate/backend/tatooine/modules/profile/app/adapters/postgresadapter/testutils"
	"github.com/omerbeden/event-mate/backend/tatooine/modules/profile/app/domain/model"
	"github.com/omerbeden/event-mate/backend/tatooine/pkg/db"
	"github.com/stretchr/testify/assert"
)

func TestInsertUser(t *testing.T) {

	tests := []struct {
		name      string
		setupMock func(*testutils.MockDBExecuter)
		user      model.UserProfile
		wantErr   bool
	}{
		{
			name:    "should insert user successfully",
			wantErr: false,
			setupMock: func(md *testutils.MockDBExecuter) {
				md.QueryRowFunc = func(ctx context.Context, sql string, args ...interface{}) db.Row {
					return &testutils.MockRow{
						ScanFunc: func(dest ...any) error {
							*dest[0].(*int64) = int64(1)
							return nil
						},
					}
				}
			},
		},
		{
			name:    "should return an error while inserting user",
			wantErr: true,
			setupMock: func(md *testutils.MockDBExecuter) {
				md.QueryRowFunc = func(ctx context.Context, sql string, args ...interface{}) db.Row {
					return &testutils.MockRow{
						ScanFunc: func(dest ...any) error {
							return errors.New("database error")
						},
					}
				}
			},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			mockDB := new(testutils.MockDBExecuter)
			if tc.setupMock != nil {
				tc.setupMock(mockDB)
			}

			userRepository := postgresadapter.NewUserProfileRepo(mockDB)
			ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
			defer cancel()

			res, err := userRepository.Insert(ctx, &tc.user)

			if tc.wantErr {
				assert.Error(t, err)
				assert.Nil(t, res)
			} else {
				assert.NoError(t, err)
				assert.NotNil(t, res)
			}
		})
	}
}

func TestGetUsersByAddress(t *testing.T) {

	users := []model.UserProfile{
		{
			Id: int64(1),
			Adress: model.UserProfileAdress{
				ProfileId: int64(1),
				City:      "San Francisco",
			},
			About:           "about",
			Name:            "name",
			LastName:        "lastName",
			ProfileImageUrl: "profileImageUrl",
			Stat: model.UserProfileStat{
				AttandedActivities: 0,
				Point:              0,
			},
			ExternalId: "ex1",
			Email:      "test",
		},
		{
			Id: int64(1),
			Adress: model.UserProfileAdress{
				ProfileId: int64(2),
				City:      "San Francisco",
			},
			About:           "about2",
			Name:            "name2",
			LastName:        "lastName2",
			ProfileImageUrl: "profileImageUrl2",
			Stat: model.UserProfileStat{
				AttandedActivities: 0,
				Point:              0,
			},
			ExternalId: "ex2",
			Email:      "test",
		},
	}
	tests := []struct {
		name      string
		setupMock func(*testutils.MockDBExecuter)
		users     []model.UserProfile
		address   model.UserProfileAdress
		wantErr   bool
	}{
		{
			name:    "should get users successfully",
			wantErr: false,
			address: model.UserProfileAdress{
				City: "San Francisco",
			},
			setupMock: func(md *testutils.MockDBExecuter) {
				md.QueryFunc = func(ctx context.Context, sql string, args ...any) (db.Rows, error) {
					return &testutils.MockRows{
						Users:   users,
						Current: 0,
					}, nil
				}
			},
		},

		{
			name: "should return database error",
			address: model.UserProfileAdress{
				City: "San Francisco",
			},
			wantErr: true,
			setupMock: func(md *testutils.MockDBExecuter) {
				md.QueryFunc = func(ctx context.Context, sql string, args ...any) (db.Rows, error) {
					return nil, errors.New("database error")
				}
			},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			mockDB := new(testutils.MockDBExecuter)
			if tc.setupMock != nil {
				tc.setupMock(mockDB)
			}

			repository := postgresadapter.NewUserProfileRepo(mockDB)
			ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
			defer cancel()

			result, err := repository.GetUsersByAddress(ctx, tc.address)

			if tc.wantErr {
				assert.Error(t, err)
				assert.Nil(t, result)
			} else {
				assert.NoError(t, err)
				assert.NotEmpty(t, result)
			}
		})
	}
}

func TestUpdateProfileImage(t *testing.T) {

	test := []struct {
		name       string
		setupMock  func(*testutils.MockDBExecuter)
		externalId string
		imageUrl   string
		wantErr    bool
	}{
		{
			name:       "should update profile image successfully",
			wantErr:    false,
			imageUrl:   "new image url.png",
			externalId: "1",
			setupMock: func(md *testutils.MockDBExecuter) {
				md.ExecFunc = func(ctx context.Context, sql string, arguments ...any) (db.CommandTag, error) {
					return db.CommandTag{}, nil
				}
			},
		},
		{
			name:       "should return an error",
			wantErr:    true,
			imageUrl:   "new image url.png",
			externalId: "1",
			setupMock: func(md *testutils.MockDBExecuter) {
				md.ExecFunc = func(ctx context.Context, sql string, arguments ...any) (db.CommandTag, error) {
					return db.CommandTag{}, errors.New("database error")
				}
			},
		},
	}

	for _, tc := range test {
		t.Run(tc.name, func(t *testing.T) {
			mockDB := new(testutils.MockDBExecuter)
			if tc.setupMock != nil {
				tc.setupMock(mockDB)
			}

			userRepository := postgresadapter.NewUserProfileRepo(mockDB)
			ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
			defer cancel()

			err := userRepository.UpdateProfileImage(ctx, tc.externalId, tc.imageUrl)

			if tc.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}

func TestDeleteUserById(t *testing.T) {

	tests := []struct {
		name       string
		externalId string
		setupMock  func(*testutils.MockDBExecuter)
		wantErr    bool
	}{
		{
			name:       "should delete user successfully",
			externalId: "1",
			wantErr:    false,
			setupMock: func(md *testutils.MockDBExecuter) {
				md.ExecFunc = func(ctx context.Context, sql string, arguments ...any) (db.CommandTag, error) {
					return db.CommandTag{}, nil
				}
			},
		},
		{
			name:       "should return an error",
			externalId: "1",
			wantErr:    true,
			setupMock: func(md *testutils.MockDBExecuter) {
				md.ExecFunc = func(ctx context.Context, sql string, arguments ...any) (db.CommandTag, error) {
					return db.CommandTag{}, errors.New("database error")
				}
			},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			mockDB := new(testutils.MockDBExecuter)
			if tc.setupMock != nil {
				tc.setupMock(mockDB)
			}

			repository := postgresadapter.NewUserProfileRepo(mockDB)
			ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
			defer cancel()

			err := repository.DeleteUser(ctx, tc.externalId)

			if tc.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}
