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

func TestCreateProfileCommand_Handle(t *testing.T) {

	tests := []struct {
		name              string
		wantErr           bool
		profile           model.UserProfile
		ProfileInsertFunc func(ctx context.Context, profile *model.UserProfile) (*model.UserProfile, error)
		AddressInsertFunc func(ctx context.Context, address model.UserProfileAdress) error
		StatInsertFunc    func(ctx context.Context, stat model.UserProfileStat) error
	}{
		{
			name:    "should insert user profile successfully",
			wantErr: false,
			profile: model.UserProfile{
				Id:       1,
				Name:     "name",
				LastName: "last name",
				About:    "about",
				Adress: model.UserProfileAdress{
					ProfileId: 1,
					City:      "San Francisco",
				},
				Stat: model.UserProfileStat{
					AttandedActivities: 0,
					Point:              1,
				},
			},
		},
		{
			name:    "should return error when userepo.Insert  fails",
			wantErr: true,
			profile: model.UserProfile{
				Id:       1,
				Name:     "name",
				LastName: "last name",
				About:    "about",
				Adress: model.UserProfileAdress{
					ProfileId: 1,
					City:      "San Francisco",
				},
				Stat: model.UserProfileStat{
					AttandedActivities: 0,
					Point:              1,
				},
			},
			ProfileInsertFunc: func(ctx context.Context, profile *model.UserProfile) (*model.UserProfile, error) {
				return nil, errors.New("database error")
			},
		},
		{
			name:    "should return error when addressRepo.Insert fails",
			wantErr: true,
			profile: model.UserProfile{
				Id:       1,
				Name:     "name",
				LastName: "last name",
				About:    "about",
				Adress: model.UserProfileAdress{
					ProfileId: 1,
					City:      "San Francisco",
				},
				Stat: model.UserProfileStat{
					AttandedActivities: 0,
					Point:              1,
				},
			},
			AddressInsertFunc: func(ctx context.Context, address model.UserProfileAdress) error {
				return errors.New("database error")
			},
		},
		{
			name:    "should return error when statRepo.Insert fails",
			wantErr: true,
			profile: model.UserProfile{
				Id:       1,
				Name:     "name",
				LastName: "last name",
				About:    "about",
				Adress: model.UserProfileAdress{
					ProfileId: 1,
					City:      "San Francisco",
				},
				Stat: model.UserProfileStat{
					AttandedActivities: 0,
					Point:              1,
				},
			},
			StatInsertFunc: func(ctx context.Context, stat model.UserProfileStat) error {
				return errors.New("database error")
			},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {

			mr, mockRedisClient := testutils.SetupMiniredis(t)
			defer mr.Close()
			defer mockRedisClient.Close()

			userRepo := &testutils.MockUserRepository{
				Profile:    tc.profile,
				InsertFunc: tc.ProfileInsertFunc,
			}
			statRepo := &testutils.MockStatRepository{
				InsertFunc: tc.StatInsertFunc,
			}
			addressRepo := &testutils.MockAddressRepository{
				InsertFunc: tc.AddressInsertFunc,
			}

			cmd := commands.CreateProfileCommand{
				UserRepo:    userRepo,
				AddressRepo: addressRepo,
				StatRepo:    statRepo,
				Cache:       mockRedisClient,
				Profile:     tc.profile,
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

/*
func TestCreateProfileCommand_Handle(t *testing.T) {

	tests := []struct {
		name    string
		wantErr bool
	}{}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {

			cmd := commands.CreateProfileCommand{}
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
*/
