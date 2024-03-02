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

func TestUpdateProfileImage_Handle(t *testing.T) {
	tests := []struct {
		name                      string
		wantErr                   bool
		imageUrl                  string
		externalId                string
		userName                  string
		UpdateProfileImageFunc    func(ctx context.Context, externalId string, imageUrl string) error
		GetCurrentUserProfileFunc func(context.Context, string) (*model.UserProfile, error)
	}{
		{
			name:       "should update user profile image successfully",
			wantErr:    false,
			imageUrl:   "testimage",
			externalId: "testuser",
		},
		{
			name:       "should return error when UpdateProfileImage failes",
			wantErr:    true,
			imageUrl:   "testimage",
			externalId: "testuser",
			UpdateProfileImageFunc: func(ctx context.Context, externalId, imageUrl string) error {
				return errors.New("database error")
			},
		},
		{
			name:       "should return error when GetCurrentUserProfile fails",
			wantErr:    true,
			imageUrl:   "testimage",
			externalId: "testuser",
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

			cmd := commands.UpdateProfileImageCommand{
				Repo: &testutils.MockUserRepository{
					GetCurrentUserProfileFunc: tc.GetCurrentUserProfileFunc,
					UpdateProfileImageFunc:    tc.UpdateProfileImageFunc,
				},
				Cache:      mockRedisClient,
				ImageUrl:   tc.imageUrl,
				ExternalId: tc.externalId,
				Username:   tc.userName,
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
