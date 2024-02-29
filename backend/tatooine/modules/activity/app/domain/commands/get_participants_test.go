package commands_test

import (
	"context"
	"encoding/json"
	"fmt"
	"testing"

	"github.com/omerbeden/event-mate/backend/tatooine/modules/activity/app/domain/commands"
	"github.com/omerbeden/event-mate/backend/tatooine/modules/activity/app/domain/commands/testutils"
	"github.com/omerbeden/event-mate/backend/tatooine/modules/activity/app/domain/model"
	"github.com/omerbeden/event-mate/backend/tatooine/pkg/cache"
	"github.com/stretchr/testify/assert"
)

func addRedisMockParticipants(t *testing.T, rc *cache.RedisClient, activityId int64) {
	t.Helper()

	cacheKey := fmt.Sprintf("participant:%d", activityId)

	participants := []model.User{
		{ID: 2},
		{ID: 3},
	}

	jsonActivities, _ := json.Marshal(participants)
	rc.AddMember(context.Background(), cacheKey, jsonActivities)

}

func TestGetParticiapnts_Handle(t *testing.T) {
	mr, mockRedisClient := testutils.SetupMiniredis(t)
	defer mr.Close()
	defer mockRedisClient.Close()

	defer mr.Close()
	defer mockRedisClient.Close()

	tests := []struct {
		name                     string
		activityId               int64
		wantRedis                bool
		wantError                bool
		setupGetParticipantsFunc func(*testutils.MockActivityRepo)
	}{
		{
			name:       "should get participants successfully from db",
			activityId: 1,
			wantRedis:  false,
			wantError:  false,
		},
		{
			name:       "should get participants successfully from redis",
			activityId: 1,
			wantRedis:  true,
			wantError:  false,
		},
		{
			name:       "should return error from db",
			activityId: 1,
			wantRedis:  false,
			wantError:  true,
			setupGetParticipantsFunc: func(mar *testutils.MockActivityRepo) {
				mar.GetPartipantsFunc = func(ctx context.Context, activityId int64) ([]model.User, error) {
					return nil, fmt.Errorf("an error occurred when getting participants")
				}
			},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			mockActivityRepo := testutils.MockActivityRepo{
				Activity: model.Activity{
					Participants: []model.User{
						{ID: 2},
						{ID: 3},
						{ID: 4},
					},
				},
			}

			cmd := commands.GetParticipantsCommand{
				ActivityRepository: &mockActivityRepo,
				ActivityId:         tc.activityId,
				Redis:              mockRedisClient,
			}

			if tc.wantRedis {
				addRedisMockParticipants(t, mockRedisClient, tc.activityId)
			}

			if tc.setupGetParticipantsFunc != nil {
				tc.setupGetParticipantsFunc(&mockActivityRepo)
			}

			result, err := cmd.Handle(context.Background())

			if tc.wantError {
				assert.Error(t, err)
				assert.Nil(t, result)
			} else {
				assert.NotEmpty(t, result)
				assert.NoError(t, err)
				assert.Equal(t, len(mockActivityRepo.Activity.Participants), len(result))
			}

		})
	}

}
