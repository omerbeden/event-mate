package commands

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/omerbeden/event-mate/backend/tatooine/modules/profile/app/adapters/cachedapter"
	"github.com/omerbeden/event-mate/backend/tatooine/modules/profile/app/domain/model"
	"github.com/omerbeden/event-mate/backend/tatooine/modules/profile/app/domain/ports/repositories"
	"github.com/omerbeden/event-mate/backend/tatooine/pkg/cache"
)

var errLogPrefixCreateCommand = "profile:createCommand"

type CreateProfileCommand struct {
	Profile model.UserProfile
	Repo    repositories.UserProfileRepository
	Cache   cache.Cache
}

func (ccmd *CreateProfileCommand) Handle(ctx context.Context) error {
	userProfile, err := ccmd.Repo.InsertUser(ctx, &ccmd.Profile)
	if err != nil {
		return err
	}

	return ccmd.addUserProfileToCache(ctx, userProfile)

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
