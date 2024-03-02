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

func TestGiveUserPoint_Handle(t *testing.T) {
	tests := []struct {
		name                    string
		wantErr                 bool
		point                   float32
		externalId              string
		UpdatePRofilePointsFunc func(context.Context, string, float32) error
		GetUserProfileFunc      func(ctx context.Context, username string) (*model.UserProfile, error)
	}{
		{
			name:       "should give user point successfully",
			wantErr:    false,
			point:      1,
			externalId: "testuser",
		},
		{
			name:       "should return an error when updateProfilePoint fails",
			wantErr:    true,
			point:      1,
			externalId: "testuser",
			UpdatePRofilePointsFunc: func(ctx context.Context, s string, f float32) error {
				return errors.New("database error")
			},
		},
		{
			name:       "should return an error when getUserProfile fails",
			wantErr:    true,
			point:      1,
			externalId: "testuser",
			GetUserProfileFunc: func(ctx context.Context, username string) (*model.UserProfile, error) {
				return nil, errors.New("database error")
			},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			mr, mockRedisClient := testutils.SetupMiniredis(t)
			defer mr.Close()
			defer mockRedisClient.Close()

			cmd := commands.GiveUserPointCommand{
				UserRepo: &testutils.MockUserRepository{
					GetUserProfileFunc: tc.GetUserProfileFunc,
				},
				StatRepo: &testutils.MockStatRepository{
					UpdateProfilePointsFunc: tc.UpdatePRofilePointsFunc,
				},
				Cache:      mockRedisClient,
				Point:      tc.point,
				ExternalId: tc.externalId,
			}
			ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
			defer cancel()

			err := cmd.Handle(ctx)

			if tc.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}

		})
	}
}
