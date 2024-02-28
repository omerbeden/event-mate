package repo_test

import (
	"context"
	"errors"
	"fmt"
	"testing"
	"time"

	"github.com/jackc/pgx/v5/pgconn"
	"github.com/omerbeden/event-mate/backend/tatooine/modules/activity/app/adapters/repo"
	"github.com/omerbeden/event-mate/backend/tatooine/modules/activity/app/domain/model"
	"github.com/stretchr/testify/assert"
)

func TestCreateLocation(t *testing.T) {
	tests := []struct {
		name        string
		id          int
		location    model.Location
		setupMock   func(*MockDBExecuter)
		expectError bool
	}{
		{
			name: "should create location",
			id:   1,
			location: model.Location{
				City: "London",
			},
			expectError: false,
			setupMock: func(md *MockDBExecuter) {
				md.ExecFunc = func(ctx context.Context, sql string, arguments ...any) (pgconn.CommandTag, error) {
					return pgconn.NewCommandTag(""), nil
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
			setupMock: func(md *MockDBExecuter) {
				md.ExecFunc = func(ctx context.Context, sql string, arguments ...any) (pgconn.CommandTag, error) {
					return pgconn.NewCommandTag(""), errors.New("database error")
				}
			},
		},
	}

	for _, tc := range tests {
		t.Run(fmt.Sprintf("%s,%d", tc.name, tc.id), func(t *testing.T) {
			mockDB := new(MockDBExecuter)
			repo := repo.NewLocationRepo(mockDB)

			if tc.setupMock != nil {
				tc.setupMock(mockDB)
			}
			ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
			defer cancel()

			res, err := repo.Create(ctx, &tc.location)

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
		setupMock   func(*MockDBExecuter)
		expectError bool
	}{
		{
			name: "should update location",
			id:   1,
			location: model.Location{
				City: "London",
			},
			expectError: false,
			setupMock: func(md *MockDBExecuter) {
				md.ExecFunc = func(ctx context.Context, sql string, arguments ...any) (pgconn.CommandTag, error) {
					return pgconn.NewCommandTag(""), nil
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
			setupMock: func(md *MockDBExecuter) {
				md.ExecFunc = func(ctx context.Context, sql string, arguments ...any) (pgconn.CommandTag, error) {
					return pgconn.NewCommandTag(""), errors.New("database error")
				}
			},
		},
	}

	for _, tc := range tests {
		t.Run(fmt.Sprintf("%s,%d", tc.name, tc.id), func(t *testing.T) {
			mockDB := new(MockDBExecuter)
			repo := repo.NewLocationRepo(mockDB)

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
