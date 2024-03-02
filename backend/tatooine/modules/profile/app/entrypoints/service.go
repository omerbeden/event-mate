package entrypoints

import (
	"context"
	"fmt"

	"github.com/omerbeden/event-mate/backend/tatooine/modules/profile/app/domain/commands"
	"github.com/omerbeden/event-mate/backend/tatooine/modules/profile/app/domain/model"
	"github.com/omerbeden/event-mate/backend/tatooine/modules/profile/app/domain/ports/repositories"
	"github.com/omerbeden/event-mate/backend/tatooine/pkg/cache"
)

type UserService struct {
	userRepository repositories.UserProfileRepository
	redisClient    cache.RedisClient
}

func NewService(
	userRepository repositories.UserProfileRepository,
	redisClient cache.RedisClient,
) *UserService {
	return &UserService{
		userRepository: userRepository,
		redisClient:    redisClient,
	}
}

func (service *UserService) CreateUser(ctx context.Context, user *model.UserProfile) error {
	createCmd := &commands.CreateProfileCommand{
		Profile:  *user,
		UserRepo: service.userRepository,
		Cache:    &service.redisClient,
	}

	return createCmd.Handle(ctx)
}

func (service *UserService) DeleteUser(ctx context.Context, externalId, userName string) error {
	deleteCmd := &commands.DeleteProfileCommand{
		Repo:       service.userRepository,
		Cache:      &service.redisClient,
		ExternalId: externalId,
		UserName:   userName,
	}

	return deleteCmd.Handle(ctx)
}

func (service *UserService) GetAttandedActivities(ctx context.Context, userId int64) ([]model.Activity, error) {
	cmd := &commands.GetAttandedActivitiesCommand{
		Repo:   service.userRepository,
		Cache:  &service.redisClient,
		UserId: userId,
	}

	return cmd.Handle(ctx)
}

func (service *UserService) UpdateProfileImage(ctx context.Context, externalId string, imageUrl string) error {
	cmd := &commands.UpdateProfileImageCommand{
		Repo:       service.userRepository,
		Cache:      &service.redisClient,
		ImageUrl:   imageUrl,
		ExternalId: externalId,
	}

	return cmd.Handle(ctx)
}

func (service *UserService) GetUserProfileStats(ctx context.Context, userId int64) (*model.UserProfileStat, error) {
	cmd := &commands.GetUserProfileStatsCommand{
		Repo:   service.userRepository,
		Cache:  &service.redisClient,
		UserId: userId,
	}

	return cmd.Handle(ctx)
}

func (service *UserService) GetCurrentUserProfile(ctx context.Context, externalId string) (*model.UserProfile, error) {
	cmd := &commands.GetCurrentUserProfileCommand{
		Repo:       service.userRepository,
		Cache:      &service.redisClient,
		ExternalId: externalId,
	}

	user, err := cmd.Handle(ctx)
	if err != nil {
		return nil, err
	}

	fmt.Printf("profile: %+v\n", user)

	// user.AttandedActivities, err = service.GetAttandedActivities(user.Id)
	// if err != nil {
	// 	return nil, err
	// }

	return user, nil
}

func (service *UserService) GetUserProfile(ctx context.Context, userName string) (*model.UserProfile, error) {
	cmd := &commands.GetUserProfileCommand{
		Repo:     service.userRepository,
		Cache:    &service.redisClient,
		UserName: userName,
	}

	user, err := cmd.Handle(ctx)
	if err != nil {
		return nil, err
	}

	fmt.Printf("profile: %+v\n", user)

	// user.AttandedActivities, err = service.GetAttandedActivities(user.Id)
	// if err != nil {
	// 	return nil, err
	// }

	return user, nil
}

func (service *UserService) GivePointsToUser(ctx context.Context, receiverUserName string, point float32) error {
	cmd := &commands.GiveUserPointCommand{
		Repo:             service.userRepository,
		Cache:            &service.redisClient,
		ReceiverUserName: receiverUserName,
		Point:            point,
	}

	return cmd.Handle(ctx)
}
