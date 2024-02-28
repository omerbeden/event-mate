package repo_test

import (
	"context"
	"errors"
	"fmt"
	"testing"
	"time"

	"github.com/jackc/pgx/v5"
	"github.com/omerbeden/event-mate/backend/tatooine/modules/activity/app/adapters/repo"
	"github.com/omerbeden/event-mate/backend/tatooine/modules/activity/app/domain/model"
	"github.com/stretchr/testify/assert"
)

func TestCreateActivityFlow(t *testing.T) {
	flow := []string{"flow1", "flow2"}
	tests := []struct {
		name        string
		id          int
		activityId  int64
		flow        []string
		expectError bool
		setupMock   func(*MockDBExecuter)
	}{
		{
			name:        "should create activity flow",
			id:          1,
			activityId:  1,
			flow:        flow,
			expectError: false,
			setupMock: func(md *MockDBExecuter) {
				md.CopyFromFunc = func(ctx context.Context, tableName pgx.Identifier, columnNames []string, rowSrc pgx.CopyFromSource) (int64, error) {
					return 2, nil
				}
			},
		},
		{
			name:        "should return error",
			id:          2,
			activityId:  1,
			flow:        flow,
			expectError: true,
			setupMock: func(md *MockDBExecuter) {
				md.CopyFromFunc = func(ctx context.Context, tableName pgx.Identifier, columnNames []string, rowSrc pgx.CopyFromSource) (int64, error) {
					return 0, errors.New("database error")
				}
			},
		},
	}
	for _, tc := range tests {
		t.Run(fmt.Sprintf("%s,%d", tc.name, tc.id), func(t *testing.T) {
			mockDB := new(MockDBExecuter)
			repo := repo.NewActivityFlowRepo(mockDB)

			ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
			defer cancel()

			if tc.setupMock != nil {
				tc.setupMock(mockDB)
			}

			err := repo.CreateActivityFlow(ctx, tc.activityId, tc.flow)

			if tc.expectError {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}

func TestGetActivityFlow(t *testing.T) {

	flow := []string{"flow1", "flow2"}
	tests := []struct {
		name        string
		id          int
		activityId  int64
		expected    []string
		expectError bool
		setupMock   func(*MockDBExecuter)
	}{
		{
			name:        "should get activity flow",
			id:          1,
			activityId:  1,
			expected:    flow,
			expectError: false,
			setupMock: func(md *MockDBExecuter) {
				md.QueryFunc = func(ctx context.Context, sql string, args ...any) (pgx.Rows, error) {
					return &MockRows{
						Activities: []model.Activity{},
						Rules:      []string{},
						Flow:       flow,
						Current:    0,
					}, nil
				}
			},
		},
		{
			name:        "should return error",
			id:          2,
			activityId:  1,
			expected:    flow,
			expectError: true,
			setupMock: func(md *MockDBExecuter) {
				md.QueryFunc = func(ctx context.Context, sql string, args ...any) (pgx.Rows, error) {
					return nil, errors.New("database error")
				}
			},
		},
	}
	for _, tc := range tests {
		t.Run(fmt.Sprintf("%s,%d", tc.name, tc.id), func(t *testing.T) {
			mockDB := new(MockDBExecuter)
			repo := repo.NewActivityFlowRepo(mockDB)

			ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
			defer cancel()

			if tc.setupMock != nil {
				tc.setupMock(mockDB)
			}

			res, err := repo.GetActivityFlow(ctx, tc.activityId)

			if tc.expectError {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.NotEmpty(t, res)
				assert.Equal(t, len(tc.expected), len(res))
			}
		})
	}
}
