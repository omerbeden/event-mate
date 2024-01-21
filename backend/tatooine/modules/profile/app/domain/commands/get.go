package commands

import (
	"fmt"

	"github.com/omerbeden/event-mate/backend/tatooine/modules/profile/app/adapters/cachedapter"
	"github.com/omerbeden/event-mate/backend/tatooine/modules/profile/app/domain/model"
	"github.com/omerbeden/event-mate/backend/tatooine/modules/profile/app/domain/ports/repositories"
)

var errLogPrefixGetUSerProfileCommand = "GetUserProfile"

type GetUserProfileCommand struct {
	Repo       repositories.UserProfileRepository
	Cache      cachedapter.Cache
	ExternalId string
}

func (cmd *GetUserProfileCommand) Handle() (*model.UserProfile, error) {
	user, err := cmd.getFromCache(cmd.ExternalId)
	if err != nil {
		fmt.Printf("%s: error while getting user profile %s from cache, returning from db", errLogPrefixGetUSerProfileCommand, cmd.ExternalId)
		return cmd.Repo.GetUserProfile(cmd.ExternalId)
	}

	return user, nil
}

func (cmd *GetUserProfileCommand) getFromCache(externalId string) (*model.UserProfile, error) {
	cacheKey := fmt.Sprintf("%s:%s", userProfileCacheKey, externalId)
	return cmd.Cache.GetUserProfile(cacheKey)
}
