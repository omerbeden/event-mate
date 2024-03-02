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

func TestGetUserProfile_Handle(t *testing.T) {
	tests := []struct {
		name               string
		wantErr            bool
		userName           string
		profiles           []model.UserProfile
		GetUserProfileFunc func(ctx context.Context, username string) (*model.UserProfile, error)
	}{
		{
			name:     "should get user profile successfully",
			wantErr:  false,
			userName: "testuser",
			profiles: []model.UserProfile{
				{
					ExternalId: "testuser",
					Name:       "test",
					LastName:   "user",
				},
			},
		},
		{
			name:     "should return error",
			wantErr:  true,
			userName: "testuser",
			profiles: []model.UserProfile{
				{
					ExternalId: "testuser",
					Name:       "test",
					LastName:   "user",
				},
			},
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

			cmd := commands.GetUserProfileCommand{
				Cache:    mockRedisClient,
				UserName: tc.userName,
				Repo: &testutils.MockUserRepository{
					Profiles:           tc.profiles,
					GetUserProfileFunc: tc.GetUserProfileFunc,
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
