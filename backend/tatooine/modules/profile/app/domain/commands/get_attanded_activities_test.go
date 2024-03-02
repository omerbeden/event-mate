package commands_test

import (
	"context"
	"errors"
	"testing"
	"time"

	"github.com/omerbeden/event-mate/backend/tatooine/modules/profile/app/domain/commands"
	"github.com/omerbeden/event-mate/backend/tatooine/modules/profile/app/domain/commands/testutils"
	"github.com/omerbeden/event-mate/backend/tatooine/modules/profile/app/domain/model"
	"github.com/stretchr/testify/assert"
)

func TestGetAttandedActivities_Handle(t *testing.T) {
	tests := []struct {
		name                      string
		userId                    int64
		wantErr                   bool
		attandedActivities        []model.Activity
		getAttandedActivitiesFunc func(context.Context, int64) ([]model.Activity, error)
	}{
		{
			name:    "should get attanded activities successfully",
			wantErr: false,
			userId:  int64(1),
			attandedActivities: []model.Activity{
				{
					ID: int64(1),
					Participants: []model.UserProfile{
						{
							Id: int64(1),
						},
						{
							Id: int64(2),
						},
					},
				},
				{
					ID: int64(2),
					Participants: []model.UserProfile{
						{
							Id: int64(1),
						},
						{
							Id: int64(3),
						},
					},
				},
			},
		},
		{
			name:    "should return an error",
			wantErr: true,
			userId:  int64(1),
			getAttandedActivitiesFunc: func(ctx context.Context, i int64) ([]model.Activity, error) {
				return nil, errors.New("database error")
			},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {

			mr, mockRedisClient := testutils.SetupMiniredis(t)
			defer mr.Close()
			defer mockRedisClient.Close()

			cmd := commands.GetAttandedActivitiesCommand{
				UserId: tc.userId,
				Cache:  mockRedisClient,
				Repo: &testutils.MockUserRepository{
					AttandedActivities:        tc.attandedActivities,
					GetAttandedActivitiesFunc: tc.getAttandedActivitiesFunc,
				},
			}
			ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
			defer cancel()

			result, err := cmd.Handle(ctx)

			if tc.wantErr {
				assert.Error(t, err)
				assert.Nil(t, result)
			} else {
				assert.NoError(t, err)
				assert.NotEmpty(t, result)
				assert.Equal(t, len(tc.attandedActivities), len(tc.attandedActivities))
			}

		})
	}
}
