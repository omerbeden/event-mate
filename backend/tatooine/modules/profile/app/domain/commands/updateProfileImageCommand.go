package commands

import (
	"encoding/json"
	"fmt"

	"github.com/omerbeden/event-mate/backend/tatooine/modules/profile/app/adapters/cachedapter"
	"github.com/omerbeden/event-mate/backend/tatooine/modules/profile/app/domain/model"
	"github.com/omerbeden/event-mate/backend/tatooine/modules/profile/app/domain/ports/repositories"
)

type UpdateProfileImageCommand struct {
	Repo       repositories.UserProfileRepository
	Cache      cachedapter.Cache
	ImageUrl   string
	ExternalId string
	Username   string
}

func (c *UpdateProfileImageCommand) Handle() error {
	err := c.Repo.UpdateProfileImage(c.ExternalId, c.ImageUrl)
	if err != nil {
		return err
	}

	updatedUser, err := c.Repo.GetCurrentUserProfile(c.ExternalId)
	if err != nil {
		return err
	}

	return c.updateCache(updatedUser)
}

func (c *UpdateProfileImageCommand) updateCache(updatedUser *model.UserProfile) error {
	cacheKeyExternalId := fmt.Sprintf("%s:%s", userProfileCacheKey, updatedUser.ExternalId)
	cacheKeyUserName := fmt.Sprintf("%s:%s", userProfileCacheKey, updatedUser.UserName)

	jsonValue, err := json.Marshal(updatedUser)
	if err != nil {
		return fmt.Errorf("parsing error while updating user profile on cache")
	}

	err = c.Cache.Set(cacheKeyExternalId, jsonValue)
	if err != nil {
		return err
	}

	err = c.Cache.Set(cacheKeyUserName, jsonValue)
	if err != nil {
		return err
	}

	return nil

}
