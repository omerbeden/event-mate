package commands

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/omerbeden/event-mate/backend/tatooine/modules/profile/app/domain/model"
	"github.com/omerbeden/event-mate/backend/tatooine/modules/profile/app/domain/ports/repositories"
	"github.com/omerbeden/event-mate/backend/tatooine/pkg"
	"github.com/omerbeden/event-mate/backend/tatooine/pkg/cache"
	"go.uber.org/zap"
)

var errLogPrefixCurrentGetUserProfileCommand = "GetUserProfile"

type GetCurrentUserProfileCommand struct {
	Repo       repositories.UserProfileRepository
	Cache      cache.Cache
	ExternalId string
}

func (cmd *GetCurrentUserProfileCommand) Handle(ctx context.Context) (*model.UserProfile, error) {
	logger, ok := ctx.Value(pkg.LoggerKey).(*zap.SugaredLogger)
	if !ok {
		return nil, fmt.Errorf("failed to get logger for CreateCommand")
	}

	user, err := cmd.getFromCache(ctx, cmd.ExternalId)
	if err != nil {
		logger.Infof("%s: error while getting user profile %s from cache, returning from db", errLogPrefixCurrentGetUserProfileCommand, cmd.ExternalId)
		return cmd.Repo.GetCurrentUserProfile(ctx, cmd.ExternalId)
	}

	return user, nil
}

func (cmd *GetCurrentUserProfileCommand) getFromCache(ctx context.Context, externalId string) (*model.UserProfile, error) {
	profileKey := fmt.Sprintf("%s:%s", cache.USER_PROFILE_CACHE_KEY, externalId)
	cacheResult, err := cmd.Cache.Get(ctx, profileKey)
	if err != nil {
		return nil, fmt.Errorf("%s could not get user profile for key: %s ", errLogPrefixCurrentGetUserProfileCommand, profileKey)
	}

	var user model.UserProfile
	err = json.Unmarshal([]byte(cacheResult.(string)), &user)
	if err != nil {
		return nil, fmt.Errorf("%s could not unmarshal result of the cache key: %s ", errLogPrefixCurrentGetUserProfileCommand, profileKey)
	}

	return &user, nil
}
