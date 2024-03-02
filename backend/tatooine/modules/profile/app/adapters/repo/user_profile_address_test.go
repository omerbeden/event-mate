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

func TestUserProfileAddressRepo_Insert(t *testing.T) {
	test := []struct {
		name      string
		wantErr   bool
		address   model.UserProfileAdress
		setupMock func(*testutils.MockDBExecuter)
	}{
		{
			name:    "should insert a profile stat successfully",
			wantErr: false,
			address: model.UserProfileAdress{
				City:      "San Francisco",
				ProfileId: 1,
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
			address: model.UserProfileAdress{
				City:      "San Francisco",
				ProfileId: 1,
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

			repository := repo.NewUserProfileAddressRepo(mockDB)
			ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
			defer cancel()

			err := repository.Insert(ctx, tc.address)

			if tc.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}
