package entrypoints

import (
	"fmt"

	"github.com/omerbeden/event-mate/backend/tatooine/modules/profile/app/adapters/cachedapter"
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

func (service *UserService) CreateUser(user *model.UserProfile) error {
	createCmd := &commands.CreateProfileCommand{
		Profile: *user,
		Repo:    service.userRepository,
		Cache:   *cachedapter.NewCache(&service.redisClient),
	}

	return createCmd.Handle()
}

func (service *UserService) DeleteUser(externalId, userName string) error {
	deleteCmd := &commands.DeleteProfileCommand{
		Repo:       service.userRepository,
		Cache:      *cachedapter.NewCache(&service.redisClient),
		ExternalId: externalId,
		UserName:   userName,
	}

	return deleteCmd.Handle()
}

func (service *UserService) GetAttandedActivities(userId int64) ([]model.Activity, error) {
	cmd := &commands.GetAttandedActivitiesCommand{
		Repo:   service.userRepository,
		Cache:  *cachedapter.NewCache(&service.redisClient),
		UserId: userId,
	}

	return cmd.Handle()
}

func (service *UserService) UpdateProfileImage(externalId string, imageUrl string) error {
	cmd := &commands.UpdateProfileImageCommand{
		Repo:       service.userRepository,
		Cache:      *cachedapter.NewCache(&service.redisClient),
		ImageUrl:   imageUrl,
		ExternalId: externalId,
	}

	return cmd.Handle()
}

func (service *UserService) GetUserProfileStats(userId int64) (*model.UserProfileStat, error) {
	cmd := &commands.GetUserProfileStatsCommand{
		Repo:   service.userRepository,
		Cache:  *cachedapter.NewCache(&service.redisClient),
		UserId: userId,
	}

	return cmd.Handle()
}

func (service *UserService) GetCurrentUserProfile(externalId string) (*model.UserProfile, error) {
	cmd := &commands.GetCurrentUserProfileCommand{
		Repo:       service.userRepository,
		Cache:      *cachedapter.NewCache(&service.redisClient),
		ExternalId: externalId,
	}

	user, err := cmd.Handle()
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

func (service *UserService) GetUserProfile(userName string) (*model.UserProfile, error) {
	cmd := &commands.GetUserProfileCommand{
		Repo:     service.userRepository,
		Cache:    *cachedapter.NewCache(&service.redisClient),
		UserName: userName,
	}

	user, err := cmd.Handle()
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

func (service *UserService) GivePointsToUser(receiverUserName string, point float32) error {
	cmd := &commands.GiveUserPointCommand{
		Repo:             service.userRepository,
		Cache:            *cachedapter.NewCache(&service.redisClient),
		ReceiverUserName: receiverUserName,
		Point:            point,
	}

	return cmd.Handle()
}
