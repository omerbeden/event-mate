package entrypoints

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"strconv"

	"github.com/omerbeden/event-mate/backend/tatooine/modules/profile/app/adapters/cachedapter"
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
		ImageUrl:   imageUrl,
		ExternalId: externalId,
	}

	err := cmd.Handle(ctx)
	if err != nil {
		return err
	}

	updatedUser, err := service.userRepository.GetCurrentUserProfile(ctx, externalId)
	if err != nil {
		return err
	}

	return updateCache(ctx, &service.redisClient, updatedUser)

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

	if user.Badges == nil {
		badges, err := service.profileBadgeRepository.GetBadges(ctx, user.Id)
		if err != nil {
			return nil, err
		}

		user.Badges = badges
	}

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
func (service *UserService) GetUserProfileById(ctx context.Context, id string) (*model.UserProfile, error) {

	//id , string to int convs and unmasking

	idInt, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		return nil, err
	}

	cmd := &commands.GetUserProfileByIdCommand{
		Repo:  service.userRepository,
		Cache: &service.redisClient,
		Id:    idInt,
	}

	user, err := cmd.Handle(ctx)
	if err != nil {
		return nil, err
	}

	// user.AttandedActivities, err = service.GetAttandedActivities(ctx, user.Id)
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

		createBadgeCommand := &commands.CreateBadgeCommand{
			BadgeRepo: service.profileBadgeRepository,
			Badge:     badge,
		}

		err = createBadgeCommand.Handle(ctx)

		if err != nil {
			return err
		}
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

func (service *UserService) GetProfileBadges(ctx context.Context, externalId string) (map[int64]*model.ProfileBadge, error) {
	cmd := commands.GetProfileBadgesCommand{
		BadgeRepo:   service.profileBadgeRepository,
		ProfileRepo: service.userRepository,
		ExternalId:  externalId,
	}

	return cmd.Handle(ctx)

}

func (service *UserService) UpdateVerification(ctx context.Context, isVerified bool, externalId string) error {

	if isVerified {
		cmd := &commands.UpdateVerificationCommand{
			Repo:       service.userRepository,
			IsVerified: isVerified,
			ExternalId: externalId,
		}

		err := cmd.Handle(ctx)
		if err != nil {
			return err
		}

		updatedUser, err := service.userRepository.GetCurrentUserProfile(ctx, externalId)
		if err != nil {
			return err
		}

		badge := model.VerifiedBadge()
		badge.ProfileId = updatedUser.Id
		_, ok := updatedUser.Badges[model.VerifiedBadgeId]
		if !ok {
			m := make(map[int64]*model.ProfileBadge)
			m[model.VerifiedBadgeId] = badge
			updatedUser.Badges = m
		}

		createBadgeCommand := &commands.CreateBadgeCommand{
			BadgeRepo: service.profileBadgeRepository,
			Badge:     badge,
		}

		err = createBadgeCommand.Handle(ctx)
		if err != nil {
			return err
		}

		return updateCache(ctx, &service.redisClient, updatedUser)

	}

	return errors.New("unverified account, verify account first")
}

func updateCache(ctx context.Context, cache cache.Cache, updatedUser *model.UserProfile) error {
	cacheKeyExternalId := fmt.Sprintf("%s:%s", cachedapter.USER_PROFILE_CACHE_KEY, updatedUser.ExternalId)
	cacheKeyUserName := fmt.Sprintf("%s:%s", cachedapter.USER_PROFILE_CACHE_KEY, updatedUser.Header.UserName)

	jsonValue, err := json.Marshal(updatedUser)
	if err != nil {
		return fmt.Errorf("parsing error while updating user profile on cache")
	}

	err = cache.Set(ctx, cacheKeyExternalId, jsonValue)
	if err != nil {
		return err
	}

	err = cache.Set(ctx, cacheKeyUserName, jsonValue)
	if err != nil {
		return err
	}

	return nil

}
