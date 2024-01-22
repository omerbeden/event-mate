package commands

import (
	"fmt"

	"github.com/omerbeden/event-mate/backend/tatooine/modules/profile/app/adapters/cachedapter"
	"github.com/omerbeden/event-mate/backend/tatooine/modules/profile/app/domain/model"
	"github.com/omerbeden/event-mate/backend/tatooine/modules/profile/app/domain/ports/repositories"
)

var errLogPrefixGetUserProfileCommand = "GetUserProfile"

type GetUserProfileCommand struct {
	Repo     repositories.UserProfileRepository
	Cache    cachedapter.Cache
	UserName string
}

func (cmd *GetUserProfileCommand) Handle() (*model.UserProfile, error) {
	user, err := cmd.getFromCache(cmd.UserName)
	if err != nil {
		fmt.Printf("%s: error while getting user profile %s from cache, returning from db", errLogPrefixGetUserProfileCommand, cmd.UserName)
		return cmd.Repo.GetUserProfile(cmd.UserName)
	}

	return user, nil
}

func (cmd *GetUserProfileCommand) getFromCache(userName string) (*model.UserProfile, error) {
	cacheKey := fmt.Sprintf("%s:%s", userProfileCacheKey, userName)
	return cmd.Cache.GetUserProfile(cacheKey)
}
