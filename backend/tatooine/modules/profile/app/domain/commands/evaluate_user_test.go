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

func TestEvaluateUser_Handle(t *testing.T) {
	tests := []struct {
		name                      string
		wantErr                   bool
		evaluation                model.UserEvaluation
		EvaluateUserFunc          func(ctx context.Context, eval model.UserEvaluation) error
		GetUserCurrentProfileFunc func(ctx context.Context, username string) (*model.UserProfile, error)
	}{
		{
			name:    "should evaluate user point successfully",
			wantErr: false,
			evaluation: model.UserEvaluation{
				GiverId:    "1",
				ReceiverId: "2",
				Points:     1,
				Comment:    "test comment",
			},
		},
		{
			name:    "should return an error when evaluateUser fails",
			wantErr: true,
			evaluation: model.UserEvaluation{
				GiverId:    "1",
				ReceiverId: "2",
				Points:     1,
				Comment:    "test comment",
			},
			EvaluateUserFunc: func(ctx context.Context, eval model.UserEvaluation) error {
				return errors.New("database error")
			},
		},
		{
			name:    "should return an error when getUserProfile fails",
			wantErr: true,
			evaluation: model.UserEvaluation{
				GiverId:    "1",
				ReceiverId: "2",
				Points:     1,
				Comment:    "test comment",
			},
			GetUserCurrentProfileFunc: func(ctx context.Context, username string) (*model.UserProfile, error) {
				return nil, errors.New("database error")
			},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			mr, mockRedisClient := testutils.SetupMiniredis(t)
			defer mr.Close()
			defer mockRedisClient.Close()

			cmd := commands.EvaluateUserCommand{
				UserRepo: &testutils.MockUserRepository{
					GetCurrentUserProfileFunc: tc.GetUserCurrentProfileFunc,
				},
				StatRepo: &testutils.MockStatRepository{
					EvaluateUserFunc: tc.EvaluateUserFunc,
				},
				Cache:      mockRedisClient,
				Evaluation: tc.evaluation,
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
