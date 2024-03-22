package entrypoints

import (
	"context"

	"github.com/omerbeden/event-mate/backend/tatooine/modules/profile/app/domain/commands"
	"github.com/omerbeden/event-mate/backend/tatooine/modules/profile/app/domain/mmap"
	"github.com/omerbeden/event-mate/backend/tatooine/modules/profile/app/domain/model"
	"github.com/omerbeden/event-mate/backend/tatooine/modules/profile/app/domain/ports/repositories"
	"github.com/omerbeden/event-mate/backend/tatooine/pkg/cache"
	"github.com/omerbeden/event-mate/backend/tatooine/pkg/db"
)

type UserService struct {
	userRepository         repositories.UserProfileRepository
	userStatRepository     repositories.UserProfileStatRepository
	userAddressRepository  repositories.UserProfileAddressRepository
	profileBadgeRepository repositories.ProfileBadgeRepository
	redisClient            cache.RedisClient
	tx                     db.TransactionManager
}

func NewService(
	userRepository repositories.UserProfileRepository,
	userStatRepository repositories.UserProfileStatRepository,
	userAddressRepository repositories.UserProfileAddressRepository,
	profileBadgeRepository repositories.ProfileBadgeRepository,
	redisClient cache.RedisClient,
	tx db.TransactionManager,

) *UserService {
	return &UserService{
		userRepository:         userRepository,
		userStatRepository:     userStatRepository,
		userAddressRepository:  userAddressRepository,
		profileBadgeRepository: profileBadgeRepository,
		redisClient:            redisClient,
		tx:                     tx,
	}
}

func (service *UserService) CreateUser(ctx context.Context, request *model.CreateUserProfileRequest) error {

	profile := mmap.CreateUserRequestToProfile(request)

	createCmd := &commands.CreateProfileCommand{
		Profile:     *profile,
		UserRepo:    service.userRepository,
		AddressRepo: service.userAddressRepository,
		BadgeRepo:   service.profileBadgeRepository,
		Cache:       &service.redisClient,
		Tx:          service.tx,
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

	badges, err := service.profileBadgeRepository.GetBadges(ctx, user.Id)
	if err != nil {
		return nil, err
	}

	user.Badges = badges

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

	// user.AttandedActivities, err = service.GetAttandedActivities(user.Id)
	// if err != nil {
	// 	return nil, err
	// }

	return user, nil
}

func (service *UserService) EvaluateUser(ctx context.Context, evaluation model.UserEvaluation) error {
	cmd := &commands.EvaluateUserCommand{
		UserRepo:   service.userRepository,
		StatRepo:   service.userStatRepository,
		Cache:      &service.redisClient,
		Evaluation: evaluation,
	}

	user, err := cmd.Handle(ctx)
	if err != nil {
		return err
	}

	badge := service.badgeDecision(user)

	if badge != nil {
		service.profileBadgeRepository.Insert(ctx, nil, badge)
	}

	return nil
}
func (service *UserService) badgeDecision(user *model.UserProfile) *model.ProfileBadge {

	var badge *model.ProfileBadge

	if user.Stat.AttandedActivities >= 5 {
		_, ok := user.Badges[model.TrustworthyBadgeId]
		if !ok {
			badge = model.TrustworthyBadge()
			badge.ProfileId = user.Id
		}
	}

	if user.Stat.Point > 7 {
		_, ok := user.Badges[model.ActiveBadgeId]
		if !ok {
			badge = model.ActiveBadge()
			badge.ProfileId = user.Id
		}
	}

	return badge
}
