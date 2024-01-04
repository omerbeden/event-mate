package commands

import (
	"encoding/json"
	"fmt"
	"strconv"

	"github.com/omerbeden/event-mate/backend/tatooine/modules/profile/app/adapters/cachedapter"
	"github.com/omerbeden/event-mate/backend/tatooine/modules/profile/app/domain/model"
	"github.com/omerbeden/event-mate/backend/tatooine/modules/profile/app/domain/ports"
)

var userProfileCacheKey = "userProfile"
var errLogPrefixCreateCommand = "profile:createCommand"

type CreateProfileCommand struct {
	Profile model.UserProfile
	Repo    ports.UserProfileRepository
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

	userId := strconv.FormatInt(userProfile.Id, 10)
	cacheKey := fmt.Sprintf("%s:%s", userProfileCacheKey, userId)

	return ccmd.Cache.Set(cacheKey, jsonValue)

}
