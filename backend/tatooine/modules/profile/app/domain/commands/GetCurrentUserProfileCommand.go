package commands

import (
	"context"
	"fmt"

	"github.com/omerbeden/event-mate/backend/tatooine/modules/profile/app/adapters/cachedapter"
	"github.com/omerbeden/event-mate/backend/tatooine/modules/profile/app/domain/model"
	"github.com/omerbeden/event-mate/backend/tatooine/modules/profile/app/domain/ports/repositories"
)

var errLogPrefixCurrentGetUserProfileCommand = "GetUserProfile"

type GetCurrentUserProfileCommand struct {
	Repo       repositories.UserProfileRepository
	Cache      cachedapter.Cache
	ExternalId string
}

func (cmd *GetCurrentUserProfileCommand) Handle(ctx context.Context) (*model.UserProfile, error) {
	user, err := cmd.getFromCache(cmd.ExternalId)
	if err != nil {
		fmt.Printf("%s: error while getting user profile %s from cache, returning from db", errLogPrefixCurrentGetUserProfileCommand, cmd.ExternalId)
		return cmd.Repo.GetCurrentUserProfile(ctx, cmd.ExternalId)
	}

	return user, nil
}

func (cmd *GetCurrentUserProfileCommand) getFromCache(externalId string) (*model.UserProfile, error) {
	cacheKey := fmt.Sprintf("%s:%s", userProfileCacheKey, externalId)
	return cmd.Cache.GetUserProfile(cacheKey)
}
