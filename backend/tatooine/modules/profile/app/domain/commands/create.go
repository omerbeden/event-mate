package commands

import (
	"encoding/json"
	"fmt"

	"github.com/omerbeden/event-mate/backend/tatooine/modules/profile/app/adapters/cachedapter"
	"github.com/omerbeden/event-mate/backend/tatooine/modules/profile/app/domain/model"
	"github.com/omerbeden/event-mate/backend/tatooine/modules/profile/app/domain/ports/repositories"
)

var userProfileCacheKey = "userProfile"
var errLogPrefixCreateCommand = "profile:createCommand"

type CreateProfileCommand struct {
	Profile model.UserProfile
	Repo    repositories.UserProfileRepository
	Cache   cachedapter.Cache
}

func (ccmd *CreateProfileCommand) Handle() error {
	userProfile, err := ccmd.Repo.InsertUser(&ccmd.Profile)
	if err != nil {
		return err
	}

	return ccmd.addUserProfileToCache(userProfile)

}

func (ccmd *CreateProfileCommand) addUserProfileToCache(userProfile *model.UserProfile) error {
	jsonValue, err := json.Marshal(userProfile)
	if err != nil {
		return fmt.Errorf("%s could not marshal , %w ", errLogPrefixCreateCommand, err)
	}

	cacheKeyCurrentUser := fmt.Sprintf("%s:%s", userProfileCacheKey, userProfile.ExternalId)
	cacheKeyUserName := fmt.Sprintf("%s:%s", userProfileCacheKey, userProfile.UserName)

	if err := ccmd.Cache.Set(cacheKeyCurrentUser, jsonValue); err != nil {
		return err
	}
	if err := ccmd.Cache.Set(cacheKeyUserName, jsonValue); err != nil {
		return err
	}

	return nil

}
