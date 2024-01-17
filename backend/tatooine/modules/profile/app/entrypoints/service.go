package entrypoints

import (
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

func (service *UserService) DeleteUser(userId int64) error {
	deleteCmd := &commands.DeleteProfileCommand{
		Repo:   service.userRepository,
		Cache:  *cachedapter.NewCache(&service.redisClient),
		UserId: userId,
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

func (service *UserService) UpdateProfileImage(userId int64, imageUrl string) error {
	cmd := &commands.UpdateProfileImageCommand{
		Repo:     service.userRepository,
		Cache:    *cachedapter.NewCache(&service.redisClient),
		ImageUrl: imageUrl,
		UserId:   userId,
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

func (service *UserService) GetUserProfile(userId int64) (*model.UserProfile, error) {
	cmd := &commands.GetUserProfileCommand{
		Repo:  service.userRepository,
		Cache: *cachedapter.NewCache(&service.redisClient),
	}

	user, err := cmd.Handle(userId)
	if err != nil {
		return nil, err
	}

	user.AttandedActivities, err = service.GetAttandedActivities(userId)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (service *UserService) AddPointsToUser(receiverId int64, point float32) error {
	cmd := &commands.GiveUserPointCommand{
		Repo:       service.userRepository,
		Cache:      *cachedapter.NewCache(&service.redisClient),
		ReceiverId: receiverId,
		Point:      point,
	}

	return cmd.Handle()
}
