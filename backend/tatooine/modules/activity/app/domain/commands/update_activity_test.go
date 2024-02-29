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

func TestUpdateActivity_Handle(t *testing.T) {

	mr, mockRedisClient := testutils.SetupMiniredis(t)
	defer mr.Close()
	defer mockRedisClient.Close()

	tests := []struct {
		name                string
		activityId          int64
		wantError           bool
		setupUpdateByIdFunc func(*testutils.MockActivityRepo)
	}{
		{
			name:       "should update activity successfully",
			activityId: 1,
			wantError:  false,
		},
		{
			name:       "should return error from db",
			activityId: 1,
			wantError:  true,
			setupUpdateByIdFunc: func(mar *testutils.MockActivityRepo) {
				mar.UpdateByIDFunc = func(ctx context.Context, activityId int64, activity model.Activity) (bool, error) {
					return false, fmt.Errorf("an error occurred when updating activity")
				}
			},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {

			cmd := commands.UpdateCommand{
				Activity: &model.Activity{
					ID: 2,
				},
				Repo: &testutils.MockActivityRepo{},
			}

			if tc.setupUpdateByIdFunc != nil {
				tc.setupUpdateByIdFunc(cmd.Repo.(*testutils.MockActivityRepo))
			}

			result, err := cmd.Handle(context.Background())

			if tc.wantError {
				assert.Error(t, err)
				assert.False(t, result)
			} else {
				assert.True(t, result)
				assert.NoError(t, err)
			}
		})
	}

}
