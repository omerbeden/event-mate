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
	"go.uber.org/zap"
)

var errLogPrefixGetUserProfileCommand = "GetUserProfile"

type GetUserProfileCommand struct {
	Repo     repositories.UserProfileRepository
	Cache    cache.Cache
	UserName string
}

func (cmd *GetUserProfileCommand) Handle(ctx context.Context) (*model.UserProfile, error) {
	logger, ok := ctx.Value(pkg.LoggerKey).(*zap.SugaredLogger)
	if !ok {
		return nil, fmt.Errorf("failed to get logger for CreateCommand")
	}

	user, err := cmd.getFromCache(ctx, cmd.UserName)
	if err != nil {
		logger.Infof("%s: error while getting user profile %s from cache, returning from db", errLogPrefixGetUserProfileCommand, cmd.UserName)
		return cmd.Repo.GetUserProfile(ctx, cmd.UserName)
	}

	return user, nil
}

func (cmd *GetUserProfileCommand) getFromCache(ctx context.Context, userName string) (*model.UserProfile, error) {
	profileKey := fmt.Sprintf("%s:%s", cachedapter.USER_PROFILE_CACHE_KEY, userName)
	cacheResult, err := cmd.Cache.Get(ctx, profileKey)
	if err != nil {
		return nil, fmt.Errorf("%s could not get user profile for key: %s ", errLogPrefixGetUserProfileCommand, profileKey)
	}

	var user model.UserProfile
	err = json.Unmarshal([]byte(cacheResult.(string)), &user)
	if err != nil {
		return nil, fmt.Errorf("%s could not unmarshal result of the cache key: %s ", errLogPrefixGetUserProfileCommand, profileKey)
	}

	return &user, nil
}
