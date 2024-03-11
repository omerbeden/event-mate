package commands

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/omerbeden/event-mate/backend/tatooine/modules/profile/app/adapters/cachedapter"
	"github.com/omerbeden/event-mate/backend/tatooine/modules/profile/app/domain/model"
	"github.com/omerbeden/event-mate/backend/tatooine/modules/profile/app/domain/ports/repositories"
	"github.com/omerbeden/event-mate/backend/tatooine/pkg"
	"github.com/omerbeden/event-mate/backend/tatooine/pkg/cache"
	"github.com/omerbeden/event-mate/backend/tatooine/pkg/db"
	"go.uber.org/zap"
)

var errLogPrefixCreateCommand = "profile:createCommand"

type CreateProfileCommand struct {
	Profile     model.UserProfile
	UserRepo    repositories.UserProfileRepository
	AddressRepo repositories.UserProfileAddressRepository
	StatRepo    repositories.UserProfileStatRepository
	Cache       cache.Cache
	Tx          db.TransactionManager
}

func (cmd *CreateProfileCommand) Handle(ctx context.Context) error {
	logger, ok := ctx.Value(pkg.LoggerKey).(*zap.SugaredLogger)
	if !ok {
		return fmt.Errorf("failed to get logger for CreateCommand")
	}

	tx, err := cmd.Tx.Begin(ctx)
	if err != nil {
		return fmt.Errorf("failed to begin transaction: %w", err)
	}

	defer func() {
		if err != nil {
			tx.Rollback(ctx)
		}
	}()

	userProfile, err := cmd.UserRepo.Insert(ctx, tx, &cmd.Profile)
	if err != nil {
		return fmt.Errorf("error while inserting user profile %w", err)
	}

	userProfile.Adress.ProfileId = userProfile.Id
	userProfile.Stat.ProfileId = userProfile.Id

	err = cmd.AddressRepo.Insert(ctx, tx, cmd.Profile.Adress)
	if err != nil {
		return fmt.Errorf("error while inserting user profile address %w", err)
	}

	if err = tx.Commit(ctx); err != nil {
		return fmt.Errorf("failed to commit transaction: %w", err)
	}

	cacheResult := cmd.addUserProfileToCache(ctx, userProfile)
	if cacheResult != nil {
		logger.Info("error while inserting user profile to cache %s", cacheResult.Error())
	}

	return nil

}

func (ccmd *CreateProfileCommand) addUserProfileToCache(ctx context.Context, userProfile *model.UserProfile) error {
	jsonValue, err := json.Marshal(userProfile)
	if err != nil {
		return fmt.Errorf("%s could not marshal , %w ", errLogPrefixCreateCommand, err)
	}

	cacheKeyCurrentUser := fmt.Sprintf("%s:%s", cachedapter.USER_PROFILE_CACHE_KEY, userProfile.ExternalId)
	cacheKeyUserName := fmt.Sprintf("%s:%s", cachedapter.USER_PROFILE_CACHE_KEY, userProfile.UserName)

	if err := ccmd.Cache.Set(ctx, cacheKeyCurrentUser, jsonValue); err != nil {
		return err
	}
	if err := ccmd.Cache.Set(ctx, cacheKeyUserName, jsonValue); err != nil {
		return err
	}

	return nil

}
