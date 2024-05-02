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
				GiverId:    1,
				ReceiverId: 2,
				Points:     1,
				Comment:    "test comment",
			},
			setupMock: func(md *testutils.MockDBExecuter) {
				md.BeginFunc = func(ctx context.Context) (db.Tx, error) {
					return &testutils.MockTx{
						ExecFunc: func(ctx context.Context, sql string, arguments ...any) (commandTag db.CommandTag, err error) {
							return db.CommandTag{}, nil
						},
					}, nil
				}
			},
		},
		{
			name:    "should return error",
			wantErr: true,
			evaluation: model.UserEvaluation{
				GiverId:    1,
				ReceiverId: 2,
				Points:     1,
				Comment:    "test comment",
			},
			setupMock: func(md *testutils.MockDBExecuter) {
				md.BeginFunc = func(ctx context.Context) (db.Tx, error) {
					return &testutils.MockTx{
						ExecFunc: func(ctx context.Context, sql string, arguments ...any) (commandTag db.CommandTag, err error) {
							return db.CommandTag{}, errors.New("database error")
						},
					}, nil
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
