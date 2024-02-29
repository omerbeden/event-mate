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

type UpdateProfileImageCommand struct {
	Repo       repositories.UserProfileRepository
	Cache      cache.Cache
	ImageUrl   string
	ExternalId string
	Username   string
}

func (c *UpdateProfileImageCommand) Handle(ctx context.Context) error {
	err := c.Repo.UpdateProfileImage(ctx, c.ExternalId, c.ImageUrl)
	if err != nil {
		return err
	}

	updatedUser, err := c.Repo.GetCurrentUserProfile(ctx, c.ExternalId)
	if err != nil {
		return err
	}

	return c.updateCache(ctx, updatedUser)
}

func (c *UpdateProfileImageCommand) updateCache(ctx context.Context, updatedUser *model.UserProfile) error {
	cacheKeyExternalId := fmt.Sprintf("%s:%s", cachedapter.USER_PROFILE_CACHE_KEY, updatedUser.ExternalId)
	cacheKeyUserName := fmt.Sprintf("%s:%s", cachedapter.USER_PROFILE_CACHE_KEY, updatedUser.UserName)

	jsonValue, err := json.Marshal(updatedUser)
	if err != nil {
		return fmt.Errorf("parsing error while updating user profile on cache")
	}

	err = c.Cache.Set(ctx, cacheKeyExternalId, jsonValue)
	if err != nil {
		return err
	}

	err = c.Cache.Set(ctx, cacheKeyUserName, jsonValue)
	if err != nil {
		return err
	}

	return nil

}
