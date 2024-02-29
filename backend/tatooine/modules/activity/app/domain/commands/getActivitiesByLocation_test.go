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

func setupRedisMockData(t *testing.T, rc *cache.RedisClient, city string) {
	t.Helper()

	cacheKey := fmt.Sprintf("city:%s", city)

	activities := []model.Activity{
		{
			ID:       1,
			Location: model.Location{City: city},
		},
		{
			ID:       2,
			Location: model.Location{City: city},
		},
	}

	jsonActivities, _ := json.Marshal(activities)
	rc.AddMember(context.Background(), cacheKey, jsonActivities)

}

func TestGetActivitiesByLocation_Handle(t *testing.T) {
	mr, mockRedisClient := testutils.SetupMiniredis(t)
	defer mr.Close()
	defer mockRedisClient.Close()

	defer mr.Close()
	defer mockRedisClient.Close()

	tests := []struct {
		name                   string
		location               model.Location
		addmockToRedis         bool
		wantError              bool
		setupGetByLocationFunc func(*testutils.MockActivityRepo)
	}{
		{
			name:           "should get activities successfully from redis",
			location:       model.Location{City: "London"},
			wantError:      false,
			addmockToRedis: true,
		},
		{
			name:           "should get activities successfully from db when redis key not found",
			location:       model.Location{City: "London"},
			addmockToRedis: false,
			wantError:      false,
		},
		{
			name:           "should return error while getting activities from db",
			location:       model.Location{City: "London"},
			addmockToRedis: false,
			wantError:      true,
			setupGetByLocationFunc: func(mar *testutils.MockActivityRepo) {
				mar.GetByLocationFunc = func(ctx context.Context, location *model.Location) ([]model.Activity, error) {
					return nil, fmt.Errorf("an error occurred when getting activities")
				}
			},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			ctx := context.Background()
			mockActivityRepo := &testutils.MockActivityRepo{
				Activities: []model.Activity{
					{
						ID:       1,
						Location: tc.location,
					},
					{
						ID:       2,
						Location: tc.location,
					},
				},
			}
			cmd := commands.GetByLocationCommand{
				Location: model.Location{},
				Redis:    mockRedisClient,
				Repo:     mockActivityRepo,
			}
			if tc.addmockToRedis {
				setupRedisMockData(t, mockRedisClient, tc.location.City)
			}
			if tc.setupGetByLocationFunc != nil {
				tc.setupGetByLocationFunc(mockActivityRepo)
			}

			result, err := cmd.Handle(ctx)

			if tc.wantError {
				assert.Error(t, err)
				assert.Nil(t, result)
			} else {
				assert.NotEmpty(t, result)
				assert.NoError(t, err)
				assert.Equal(t, len(mockActivityRepo.Activities), len(result))
			}

		})
	}

}
