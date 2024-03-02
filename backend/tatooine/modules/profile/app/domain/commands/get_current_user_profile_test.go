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

func TestGetCurrentUserProfile_Handle(t *testing.T) {
	tests := []struct {
		name                      string
		wantErr                   bool
		externalId                string
		profiles                  []model.UserProfile
		GetCurrentUserProfileFunc func(context.Context, string) (*model.UserProfile, error)
	}{
		{
			name:       "should get user profile successfully",
			wantErr:    false,
			externalId: "testuser",
			profiles: []model.UserProfile{
				{
					ExternalId: "testuser",
					Name:       "test",
					LastName:   "user",
				},
			},
		},
		{
			name:       "should get user profile successfully",
			wantErr:    false,
			externalId: "testuser",
			profiles: []model.UserProfile{
				{
					ExternalId: "testuser",
					Name:       "test",
					LastName:   "user",
				},
			},
			GetCurrentUserProfileFunc: func(ctx context.Context, s string) (*model.UserProfile, error) {
				return nil, errors.New("database error")
			},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {

			mr, mockRedisClient := testutils.SetupMiniredis(t)
			defer mr.Close()
			defer mockRedisClient.Close()

			cmd := commands.GetCurrentUserProfileCommand{
				Cache:      mockRedisClient,
				ExternalId: tc.externalId,
				Repo: &testutils.MockUserRepository{
					Profiles: tc.profiles,
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
				assert.NotNil(t, result)
			}

		})
	}
}
