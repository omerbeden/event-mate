package postgresadapter_test

import (
	"context"
	"errors"
	"fmt"
	"testing"
	"time"

	"github.com/omerbeden/event-mate/backend/tatooine/modules/activity/app/adapters/postgresadapter"
	"github.com/omerbeden/event-mate/backend/tatooine/modules/activity/app/adapters/postgresadapter/testutils"
	"github.com/omerbeden/event-mate/backend/tatooine/modules/activity/app/domain/model"
	"github.com/omerbeden/event-mate/backend/tatooine/pkg/db"
	"github.com/stretchr/testify/assert"
)

func TestCreateLocation(t *testing.T) {
	tests := []struct {
		name        string
		id          int
		location    model.Location
		setupMock   func(*testutils.MockDBExecuter)
		expectError bool
	}{
		{
			name: "should create location",
			id:   1,
			location: model.Location{
				City: "London",
			},
			expectError: false,
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
			name: "should return error",
			id:   2,
			location: model.Location{
				City: "London",
			},
			expectError: true,
			setupMock: func(md *testutils.MockDBExecuter) {
				md.BeginFunc = func(ctx context.Context) (db.Tx, error) {
					return &testutils.MockTx{
						ExecFunc: func(ctx context.Context, sql string, arguments ...any) (commandTag db.CommandTag, err error) {
							return db.CommandTag{}, fmt.Errorf("database error")
						},
					}, nil
				}
			},
		},
	}

	for _, tc := range tests {
		t.Run(fmt.Sprintf("%s,%d", tc.name, tc.id), func(t *testing.T) {
			mockDB := new(testutils.MockDBExecuter)
			repo := postgresadapter.NewLocationRepo(mockDB)

			if tc.setupMock != nil {
				tc.setupMock(mockDB)
			}
			ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
			defer cancel()

			tx, _ := mockDB.Begin(ctx)
			res, err := repo.Create(ctx, tx, &tc.location)

			if tc.expectError {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.True(t, res)
			}
		})
	}
}

func TestUpdateLocation(t *testing.T) {
	tests := []struct {
		name        string
		id          int
		location    model.Location
		setupMock   func(*testutils.MockDBExecuter)
		expectError bool
	}{
		{
			name: "should update location",
			id:   1,
			location: model.Location{
				City: "London",
			},
			expectError: false,
			setupMock: func(md *testutils.MockDBExecuter) {
				md.ExecFunc = func(ctx context.Context, sql string, arguments ...any) (db.CommandTag, error) {
					return db.CommandTag{}, nil
				}
			},
		},
		{
			name: "should return error",
			id:   2,
			location: model.Location{
				City: "NEW York",
			},
			expectError: true,
			setupMock: func(md *testutils.MockDBExecuter) {
				md.ExecFunc = func(ctx context.Context, sql string, arguments ...any) (db.CommandTag, error) {
					return db.CommandTag{}, errors.New("database error")
				}
			},
		},
	}

	for _, tc := range tests {
		t.Run(fmt.Sprintf("%s,%d", tc.name, tc.id), func(t *testing.T) {
			mockDB := new(testutils.MockDBExecuter)
			repo := postgresadapter.NewLocationRepo(mockDB)

			if tc.setupMock != nil {
				tc.setupMock(mockDB)
			}
			ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
			defer cancel()

			res, err := repo.UpdateByID(ctx, tc.location)

			if tc.expectError {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.True(t, res)
			}
		})
	}
}
