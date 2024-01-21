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
}

func (c *UpdateProfileImageCommand) Handle() error {
	updatedUser, err := c.Repo.UpdateProfileImage(c.ExternalId, c.ImageUrl)
	if err != nil {
		return err
	}

	fmt.Printf(" updated : %+v\n", updatedUser)

	return c.updateCache(c.ExternalId, updatedUser)
}

func (c *UpdateProfileImageCommand) updateCache(externalId string, updatedUser *model.UserProfile) error {
	cacheKey := fmt.Sprintf("%s:%s", userProfileCacheKey, externalId)

	jsonValue, err := json.Marshal(updatedUser)
	if err != nil {
		return fmt.Errorf("parsing error while updating user profile on cache")
	}

	return c.Cache.Set(cacheKey, jsonValue)
}
