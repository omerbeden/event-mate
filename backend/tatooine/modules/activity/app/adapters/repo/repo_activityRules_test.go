package repo_test

import (
	"context"
	"errors"
	"fmt"
	"testing"
	"time"

	"github.com/jackc/pgx/v5"
	"github.com/omerbeden/event-mate/backend/tatooine/modules/activity/app/adapters/repo"
	"github.com/omerbeden/event-mate/backend/tatooine/modules/activity/app/adapters/repo/testutils"
	"github.com/omerbeden/event-mate/backend/tatooine/modules/activity/app/domain/model"
	"github.com/stretchr/testify/assert"
)

func TestCreateActivityRules(t *testing.T) {
	tests := []struct {
		name        string
		id          int
		activityId  int64
		rules       []string
		expectError bool
		setupMock   func(*testutils.MockDBExecuter)
	}{
		{
			name:        "should create activity rules",
			id:          1,
			activityId:  1,
			rules:       []string{"rule1", "rule2"},
			expectError: false,
			setupMock: func(md *testutils.MockDBExecuter) {
				md.CopyFromFunc = func(ctx context.Context, tableName pgx.Identifier, columnNames []string, rowSrc pgx.CopyFromSource) (int64, error) {
					return 2, nil
				}
			},
		},
		{
			name:        "should return error",
			id:          2,
			activityId:  1,
			rules:       []string{"rule1", "rule2"},
			expectError: true,
			setupMock: func(md *testutils.MockDBExecuter) {
				md.CopyFromFunc = func(ctx context.Context, tableName pgx.Identifier, columnNames []string, rowSrc pgx.CopyFromSource) (int64, error) {
					return 0, errors.New("database error")
				}
			},
		},
	}
	for _, tc := range tests {
		t.Run(fmt.Sprintf("%s,%d", tc.name, tc.id), func(t *testing.T) {
			mockDB := new(testutils.MockDBExecuter)
			repo := repo.NewActivityRulesRepo(mockDB)

			ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
			defer cancel()

			if tc.setupMock != nil {
				tc.setupMock(mockDB)
			}

			err := repo.CreateActivityRules(ctx, tc.activityId, tc.rules)

			if tc.expectError {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}

func TestGetActivityRules(t *testing.T) {

	rules := []string{"rule1", "rule2"}
	tests := []struct {
		name        string
		id          int
		activityId  int64
		expected    []string
		expectError bool
		setupMock   func(*testutils.MockDBExecuter)
	}{
		{
			name:        "should get activity rules",
			id:          1,
			activityId:  1,
			expected:    rules,
			expectError: false,
			setupMock: func(md *testutils.MockDBExecuter) {
				md.QueryFunc = func(ctx context.Context, sql string, args ...any) (pgx.Rows, error) {
					return &testutils.MockRows{
						Activities: []model.Activity{},
						Rules:      rules,
						Current:    0,
					}, nil
				}
			},
		},
		{
			name:        "should return error",
			id:          2,
			activityId:  1,
			expected:    rules,
			expectError: true,
			setupMock: func(md *testutils.MockDBExecuter) {
				md.QueryFunc = func(ctx context.Context, sql string, args ...any) (pgx.Rows, error) {
					return nil, errors.New("database error")
				}
			},
		},
	}
	for _, tc := range tests {
		t.Run(fmt.Sprintf("%s,%d", tc.name, tc.id), func(t *testing.T) {
			mockDB := new(testutils.MockDBExecuter)
			repo := repo.NewActivityRulesRepo(mockDB)

			ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
			defer cancel()

			if tc.setupMock != nil {
				tc.setupMock(mockDB)
			}

			res, err := repo.GetActivityRules(ctx, tc.activityId)

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
