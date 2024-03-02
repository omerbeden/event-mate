package commands_test

import (
	"context"
	"errors"
	"testing"
	"time"

	"github.com/omerbeden/event-mate/backend/tatooine/modules/profile/app/domain/commands"
	"github.com/omerbeden/event-mate/backend/tatooine/modules/profile/app/domain/commands/testutils"
	"github.com/stretchr/testify/assert"
)

func TestDeleteProflieCommand_Handle(t *testing.T) {
	tests := []struct {
		name           string
		externalId     string
		userName       string
		wantErr        bool
		DeleteUserFunc func(ctx context.Context, externalId string) error
	}{
		{
			name:       "should delete user successfully",
			wantErr:    false,
			userName:   "testuser",
			externalId: "testExternalId",
		},
		{
			name:       "should return error",
			wantErr:    true,
			userName:   "testuser",
			externalId: "testExternalId",
			DeleteUserFunc: func(ctx context.Context, externalId string) error {
				return errors.New("database error")
			},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {

			mr, mockRedisClient := testutils.SetupMiniredis(t)
			defer mr.Close()
			defer mockRedisClient.Close()

			cmd := commands.DeleteProfileCommand{
				Cache:      mockRedisClient,
				UserName:   tc.userName,
				ExternalId: tc.externalId,
				Repo: &testutils.MockUserRepository{
					DeleteUserFunc: tc.DeleteUserFunc,
				},
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
