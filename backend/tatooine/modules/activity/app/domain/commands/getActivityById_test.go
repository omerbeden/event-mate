package commands_test

import (
	"context"
	"fmt"
	"testing"

	"github.com/omerbeden/event-mate/backend/tatooine/modules/activity/app/domain/commands"
	"github.com/omerbeden/event-mate/backend/tatooine/modules/activity/app/domain/commands/testutils"
	"github.com/omerbeden/event-mate/backend/tatooine/modules/activity/app/domain/model"
	"github.com/stretchr/testify/assert"
)

func TestGetActivityById_Handle(t *testing.T) {
	mr, mockRedisClient := testutils.SetupMiniredis(t)
	defer mr.Close()
	defer mockRedisClient.Close()

	tests := []struct {
		name                     string
		activityId               int64
		wantError                bool
		addmockToRedis           bool
		setupGetByIdFunc         func(*testutils.MockActivityRepo)
		setupGetActivityRuleFunc func(*testutils.MockActivityRulesRepo)
		setupGetActivityFlowFunc func(*testutils.MockActivityFlowRepo)
	}{
		{
			name:           "should get activity successfully from redis",
			activityId:     1,
			wantError:      false,
			addmockToRedis: true,
		},
		{
			name:           "should get activity successfully from db",
			activityId:     1,
			wantError:      false,
			addmockToRedis: false,
		},
		{
			name:           "should return error when GetById failed",
			activityId:     1,
			wantError:      true,
			addmockToRedis: false,
			setupGetByIdFunc: func(mar *testutils.MockActivityRepo) {
				mar.GetByIDFunc = func(ctx context.Context, id int64) (*model.Activity, error) {
					return nil, fmt.Errorf("an error occurred when getting activity")
				}
			},
		},
		{
			name:           "should return error when GetActivityRules failed",
			activityId:     1,
			wantError:      true,
			addmockToRedis: false,
			setupGetActivityRuleFunc: func(mar *testutils.MockActivityRulesRepo) {
				mar.GetActivityRulesFunc = func(ctx context.Context, activityId int64) ([]string, error) {
					return nil, fmt.Errorf("an error occurred when getting activity rules")
				}
			},
		},

		{
			name:           "should return error when GetActivityFlow failed",
			activityId:     1,
			wantError:      true,
			addmockToRedis: false,
			setupGetActivityFlowFunc: func(mar *testutils.MockActivityFlowRepo) {
				mar.GetActivityFlowFunc = func(ctx context.Context, i int64) ([]string, error) {
					return nil, fmt.Errorf("an error occurred when getting activity flow")
				}
			},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			ctx := context.Background()

			mockActivityRepo := &testutils.MockActivityRepo{
				Activity: model.Activity{
					ID: 1,
				},
			}
			mockActivityRulesRepo := &testutils.MockActivityRulesRepo{}
			mockActivityFlowRepo := &testutils.MockActivityFlowRepo{}

			cmd := commands.GetByIDCommand{
				ActivityId:        tc.activityId,
				Redis:             mockRedisClient,
				Repo:              mockActivityRepo,
				ActivityFlowRepo:  mockActivityFlowRepo,
				ActivityRulesRepo: mockActivityRulesRepo,
			}

			if tc.setupGetByIdFunc != nil {
				tc.setupGetByIdFunc(mockActivityRepo)
			}
			if tc.setupGetActivityRuleFunc != nil {
				tc.setupGetActivityRuleFunc(mockActivityRulesRepo)
			}
			if tc.setupGetActivityFlowFunc != nil {
				tc.setupGetActivityFlowFunc(mockActivityFlowRepo)
			}

			result, err := cmd.Handle(ctx)

			if tc.wantError {
				assert.Error(t, err)
				assert.Nil(t, result)
			} else {
				assert.NotNil(t, result)
				assert.NoError(t, err)
			}

		})
	}
}
