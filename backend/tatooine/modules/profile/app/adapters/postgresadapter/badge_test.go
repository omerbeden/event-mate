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

func TestProfileBadge_Insert(t *testing.T) {
	test := []struct {
		name      string
		wantErr   bool
		badge     model.ProfileBadge
		setupMock func(*testutils.MockDBExecuter)
	}{
		{
			name:    "should insert a profile badge",
			wantErr: false,
			badge: model.ProfileBadge{
				ProfileId: 1,
				BadgeId:   1,
				ImageUrl:  "trustworthy.png",
				Text:      "trustworthy",
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
			badge: model.ProfileBadge{
				ProfileId: 1,
				BadgeId:   1,
				ImageUrl:  "trustworthy.png",
				Text:      "trustworthy",
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

			repo := postgresadapter.NewBadgeRepo(mockDB)

			err := repo.Insert(ctx, nil, &tc.badge)

			if tc.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}
