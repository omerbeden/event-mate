package commands

import (
	"fmt"

	"github.com/omerbeden/event-mate/backend/tatooine/modules/profile/app/adapters/cachedapter"
	"github.com/omerbeden/event-mate/backend/tatooine/modules/profile/app/domain/model"
	"github.com/omerbeden/event-mate/backend/tatooine/modules/profile/app/domain/ports/repositories"
)

var errLogPrefixGetUSerProfileCommand = "GetUserProfile"

type GetUserProfileCommand struct {
	Repo  repositories.UserProfileRepository
	Cache cachedapter.Cache
}

func (cmd *GetUserProfileCommand) Handle(userId int64) (*model.UserProfile, error) {
	user, err := cmd.getFromCache(userId)
	if err != nil {
		fmt.Printf("%s: error while getting user profile %d from cache, returning from db", errLogPrefixGetUSerProfileCommand, userId)
		return cmd.Repo.GetUserProfile(userId)
	}

	return user, nil
}

func (cmd *GetUserProfileCommand) getFromCache(userId int64) (*model.UserProfile, error) {
	cacheKey := fmt.Sprintf("%s:%d", userProfileCacheKey, userId)
	return cmd.Cache.GetUserProfile(cacheKey)
}
