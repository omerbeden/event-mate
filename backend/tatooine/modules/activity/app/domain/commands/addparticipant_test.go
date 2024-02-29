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

func TestAddParticipant_Handle(t *testing.T) {
	mr, mockRedisClient := testutils.SetupMiniredis(t)
	defer mr.Close()
	defer mockRedisClient.Close()

	tests := []struct {
		name                    string
		participant             model.User
		activityId              int64
		wantError               bool
		setupAddParticipantFunc func(*testutils.MockActivityRepo)
	}{
		{
			name:        "should add participant successfully",
			participant: model.User{ID: 2},
			activityId:  int64(1),
			wantError:   false,
		},
		{
			name:        "should add participant successfully",
			participant: model.User{ID: 2},
			activityId:  int64(1),
			wantError:   true,
			setupAddParticipantFunc: func(mar *testutils.MockActivityRepo) {
				mar.AddParticipantFunc = func(ctx context.Context, activityId int64, participant model.User) error {
					return fmt.Errorf("an error occurred while adding participant")
				}
			},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			ctx := context.Background()

			mockActivityRepo := &testutils.MockActivityRepo{}
			cmd := commands.AddParticipantCommand{
				Participant:        tc.participant,
				ActivityId:         tc.activityId,
				Redis:              mockRedisClient,
				ActivityRepository: mockActivityRepo,
			}
			if tc.setupAddParticipantFunc != nil {
				tc.setupAddParticipantFunc(mockActivityRepo)
			}

			err := cmd.Handle(ctx)

			if tc.wantError {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}
