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

func TestCreateActivityFlow(t *testing.T) {
	flow := []string{"flow1", "flow2"}
	tests := []struct {
		name        string
		id          int
		activityId  int64
		flow        []string
		expectError bool
		setupMock   func(*testutils.MockDBExecuter)
	}{
		{
			name:        "should create activity flow",
			id:          1,
			activityId:  1,
			flow:        flow,
			expectError: false,
			setupMock: func(md *testutils.MockDBExecuter) {
				md.BeginFunc = func(ctx context.Context) (db.Tx, error) {
					return &testutils.MockTx{
						CopyFromFunc: func(ctx context.Context, tableName db.Identifier, columnNames []string, rowSrc db.CopyFromSource) (int64, error) {
							return 2, nil

						},
					}, nil
				}
			},
		},
		{
			name:        "should return error",
			id:          2,
			activityId:  1,
			flow:        flow,
			expectError: true,
			setupMock: func(md *testutils.MockDBExecuter) {
				md.BeginFunc = func(ctx context.Context) (db.Tx, error) {
					return &testutils.MockTx{
						CopyFromFunc: func(ctx context.Context, tableName db.Identifier, columnNames []string, rowSrc db.CopyFromSource) (int64, error) {
							return 0, errors.New("database error")
						},
					}, nil
				}
			},
		},
	}
	for _, tc := range tests {
		t.Run(fmt.Sprintf("%s,%d", tc.name, tc.id), func(t *testing.T) {
			mockDB := new(testutils.MockDBExecuter)
			repo := postgresadapter.NewActivityFlowRepo(mockDB)

			ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
			defer cancel()

			if tc.setupMock != nil {
				tc.setupMock(mockDB)
			}

			tx, _ := mockDB.Begin(ctx)

			err := repo.CreateActivityFlow(ctx, tx, tc.activityId, tc.flow)

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
		setupMock   func(*testutils.MockDBExecuter)
	}{
		{
			name:        "should get activity flow",
			id:          1,
			activityId:  1,
			expected:    flow,
			expectError: false,
			setupMock: func(md *testutils.MockDBExecuter) {
				md.QueryFunc = func(ctx context.Context, sql string, args ...any) (db.Rows, error) {
					return &testutils.MockRows{
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
			setupMock: func(md *testutils.MockDBExecuter) {
				md.QueryFunc = func(ctx context.Context, sql string, args ...any) (db.Rows, error) {
					return nil, errors.New("database error")
				}
			},
		},
	}
	for _, tc := range tests {
		t.Run(fmt.Sprintf("%s,%d", tc.name, tc.id), func(t *testing.T) {
			mockDB := new(testutils.MockDBExecuter)
			repo := postgresadapter.NewActivityFlowRepo(mockDB)

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
