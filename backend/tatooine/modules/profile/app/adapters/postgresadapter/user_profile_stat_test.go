package postgresadapter_test

import (
	"context"
	"errors"
	"testing"
	"time"

	repo "github.com/omerbeden/event-mate/backend/tatooine/modules/profile/app/adapters/postgresadapter"
	"github.com/omerbeden/event-mate/backend/tatooine/modules/profile/app/adapters/postgresadapter/testutils"
	"github.com/omerbeden/event-mate/backend/tatooine/modules/profile/app/domain/model"
	"github.com/omerbeden/event-mate/backend/tatooine/pkg/db"
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
				md.ExecFunc = func(ctx context.Context, sql string, arguments ...any) (db.CommandTag, error) {
					return db.CommandTag{}, nil
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

			repository := repo.NewUserProfileStatRepo(mockDB)
			ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
			defer cancel()

			tx, _ := mockDB.Begin(ctx)

			err := repository.Insert(ctx, tx, tc.stat)

			if tc.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}

func TestUserProfileStatRepo_EvaluateUser(t *testing.T) {
	test := []struct {
		name       string
		wantErr    bool
		evaluation model.UserEvaluation
		setupMock  func(*testutils.MockDBExecuter)
	}{
		{
			name:    "should insert a profile stat successfully",
			wantErr: false,
			evaluation: model.UserEvaluation{
				GiverId:    "1",
				ReceiverId: "2",
				Points:     1,
				Comment:    "test comment",
			},
			setupMock: func(md *testutils.MockDBExecuter) {
				md.ExecFunc = func(ctx context.Context, sql string, arguments ...any) (db.CommandTag, error) {
					return db.CommandTag{}, nil
				}
			},
		},
		{
			name:    "should return error",
			wantErr: true,
			evaluation: model.UserEvaluation{
				GiverId:    "1",
				ReceiverId: "2",
				Points:     1,
				Comment:    "test comment",
			},
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

			ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
			defer cancel()

			repository := repo.NewUserProfileStatRepo(mockDB)

			err := repository.EvaluateUser(ctx, tc.evaluation)

			if tc.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}
