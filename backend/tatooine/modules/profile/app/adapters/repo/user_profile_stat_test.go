package repo_test

import (
	"context"
	"errors"
	"testing"
	"time"

	"github.com/jackc/pgx/v5/pgconn"
	"github.com/omerbeden/event-mate/backend/tatooine/modules/profile/app/adapters/repo"
	"github.com/omerbeden/event-mate/backend/tatooine/modules/profile/app/adapters/repo/testutils"
	"github.com/omerbeden/event-mate/backend/tatooine/modules/profile/app/domain/model"
	"github.com/stretchr/testify/assert"
)

func TestUserProfileStatRepo_Insert(t *testing.T) {
	test := []struct {
		name      string
		wantErr   bool
		stat      model.UserProfileStat
		setupMock func(*testutils.MockDBExecuter)
	}{
		{
			name:    "should insert a profile stat successfully",
			wantErr: false,
			stat: model.UserProfileStat{
				ProfileId:          1,
				Point:              10,
				AttandedActivities: 3,
			},
			setupMock: func(md *testutils.MockDBExecuter) {
				md.ExecFunc = func(ctx context.Context, sql string, arguments ...any) (pgconn.CommandTag, error) {
					return pgconn.NewCommandTag(""), nil
				}
			},
		},
		{
			name:    "should return error",
			wantErr: true,
			stat: model.UserProfileStat{
				ProfileId:          1,
				Point:              10,
				AttandedActivities: 3,
			},
			setupMock: func(md *testutils.MockDBExecuter) {
				md.ExecFunc = func(ctx context.Context, sql string, arguments ...any) (pgconn.CommandTag, error) {
					return pgconn.NewCommandTag(""), errors.New("database error")
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

			repository := repo.NewUserProfileStatRepo(mockDB)
			ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
			defer cancel()

			err := repository.Insert(ctx, tc.stat)

			if tc.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}

func TestUserProfileStatRepo_UpdateProfilePoints(t *testing.T) {
	test := []struct {
		name             string
		wantErr          bool
		receiverUserName string
		point            float32
		setupMock        func(*testutils.MockDBExecuter)
	}{
		{
			name:             "should insert a profile stat successfully",
			wantErr:          false,
			receiverUserName: "testuser",
			setupMock: func(md *testutils.MockDBExecuter) {
				md.ExecFunc = func(ctx context.Context, sql string, arguments ...any) (pgconn.CommandTag, error) {
					return pgconn.NewCommandTag(""), nil
				}
			},
		},
		{
			name:             "should return error",
			wantErr:          true,
			receiverUserName: "testuser",
			setupMock: func(md *testutils.MockDBExecuter) {
				md.ExecFunc = func(ctx context.Context, sql string, arguments ...any) (pgconn.CommandTag, error) {
					return pgconn.NewCommandTag(""), errors.New("database error")
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

			ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
			defer cancel()

			repository := repo.NewUserProfileStatRepo(mockDB)

			err := repository.UpdateProfilePoints(ctx, tc.receiverUserName, 3)

			if tc.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}
