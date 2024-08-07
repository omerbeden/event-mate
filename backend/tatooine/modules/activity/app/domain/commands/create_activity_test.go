package commands_test

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/omerbeden/event-mate/backend/tatooine/modules/activity/app/domain/commands"
	"github.com/omerbeden/event-mate/backend/tatooine/modules/activity/app/domain/commands/testutils"
	"github.com/omerbeden/event-mate/backend/tatooine/modules/activity/app/domain/model"
	"github.com/omerbeden/event-mate/backend/tatooine/pkg/db"
	"github.com/stretchr/testify/assert"
)

func TestCreateCommand_Handle(t *testing.T) {
	mr, mockRedisClient := testutils.SetupMiniredis(t)
	defer mr.Close()
	defer mockRedisClient.Close()

	expected := model.Activity{
		ID:        int64(1),
		Title:     "test",
		Category:  "test",
		CreatedBy: model.User{ID: 1},
		StartAt:   time.Now(),
		EndAt:     time.Now(),
		Content:   "test",
		Quota:     1,
		Location:  model.Location{ActivityId: int64(1), City: "London"},
		Participants: []model.User{
			{ID: 2},
			{ID: 3},
		},
		Rules: []string{"rule1", "rule2"},
		Flow:  []string{"flow1", "flow2"},
	}

	tests := []struct {
		name                    string
		activity                model.Activity
		wantError               bool
		setupCreateActivityFunc func(*testutils.MockActivityRepo)
		setupActivityRulesFunc  func(*testutils.MockActivityRulesRepo)
		setupActivityFlowFunc   func(*testutils.MockActivityFlowRepo)
		setupLocationFunc       func(*testutils.MockLocationRepo)
	}{
		{
			name:      "should create an activity",
			activity:  expected,
			wantError: false,
		},
		{
			name:      "should return an error when Create function returns an error",
			activity:  expected,
			wantError: true,
			setupCreateActivityFunc: func(ar *testutils.MockActivityRepo) {

				ar.CreateFunc = func(ctx context.Context, tx db.Tx, activity model.Activity) (*model.Activity, error) {
					return nil, fmt.Errorf("an error occurred when creating")
				}
			},
		},
		{
			name:      "should return an error when CreateActivityRules function returns an error",
			activity:  expected,
			wantError: true,
			setupActivityRulesFunc: func(ar *testutils.MockActivityRulesRepo) {
				ar.CreateActivityRulesFunc = func(ctx context.Context, tx db.Tx, activityId int64, rules []string) error {
					return fmt.Errorf("an error occurred when creating")
				}
			},
		},
		{
			name:      "should return an error when CreateActivityFlow function returns an error",
			activity:  expected,
			wantError: true,
			setupActivityFlowFunc: func(ar *testutils.MockActivityFlowRepo) {
				ar.CreateActivityFlowFunc = func(ctx context.Context, tx db.Tx, activityId int64, rules []string) error {
					return fmt.Errorf("an error occurred when creating")
				}
			},
		},
		{
			name:      "should return an error when Location Create function returns an error",
			activity:  expected,
			wantError: true,
			setupLocationFunc: func(ar *testutils.MockLocationRepo) {
				ar.CreateFunc = func(ctx context.Context, tx db.Tx, location *model.Location) (bool, error) {
					return false, fmt.Errorf("an error occurred when creating")
				}
			},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {

			ctx := context.Background()
			mockActivityRepo := &testutils.MockActivityRepo{
				Activity: expected,
			}
			mockActivityRulesRepo := &testutils.MockActivityRulesRepo{
				Rules: []string{"rule1", "rule2"},
			}
			mockActivityFlowRepo := &testutils.MockActivityFlowRepo{
				Flow: []string{"rule1", "rule2"},
			}
			mockLocationRepo := new(testutils.MockLocationRepo)

			if tc.setupCreateActivityFunc != nil {
				tc.setupCreateActivityFunc(mockActivityRepo)
			}
			if tc.setupActivityRulesFunc != nil {
				tc.setupActivityRulesFunc(mockActivityRulesRepo)
			}
			if tc.setupActivityFlowFunc != nil {
				tc.setupActivityFlowFunc(mockActivityFlowRepo)
			}
			if tc.setupLocationFunc != nil {
				tc.setupLocationFunc(mockLocationRepo)
			}

			ccmd := commands.CreateCommand{
				Activity:          tc.activity,
				ActivityRepo:      mockActivityRepo,
				ActivityRulesRepo: mockActivityRulesRepo,
				ActivityFlowRepo:  mockActivityFlowRepo,
				LocRepo:           mockLocationRepo,
				Redis:             mockRedisClient,
				Tx:                &testutils.MockTxnManager{},
			}

			res, err := ccmd.Handle(ctx)

			if tc.wantError {
				assert.Error(t, err)
				assert.False(t, res)
			} else {
				assert.True(t, res)
				assert.NoError(t, err)
			}
		})
	}
}
