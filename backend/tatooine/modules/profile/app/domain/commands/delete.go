package commands

import (
	"context"
	"fmt"

	"github.com/omerbeden/event-mate/backend/tatooine/modules/profile/app/domain/ports/repositories"
	"github.com/omerbeden/event-mate/backend/tatooine/pkg/cache"
)

type DeleteProfileCommand struct {
	Repo       repositories.UserProfileRepository
	Cache      cache.Cache
	ExternalId string
	UserName   string
}

func (c *DeleteProfileCommand) Handle(ctx context.Context) error {
	err := c.Repo.DeleteUser(ctx, c.ExternalId)
	if err != nil {
		return err
	}

	return c.deleteFromCache(ctx)
}

func (c *DeleteProfileCommand) deleteFromCache(ctx context.Context) error {
	cacheKeyExternalId := fmt.Sprintf("%s:%s", cache.USER_PROFILE_CACHE_KEY, c.ExternalId)
	cacheKeyUserName := fmt.Sprintf("%s:%s", cache.USER_PROFILE_CACHE_KEY, c.UserName)

	err := c.Cache.Delete(ctx, cacheKeyExternalId)
	if err != nil {
		return err
	}

	err = c.Cache.Delete(ctx, cacheKeyUserName)
	if err != nil {
		return err
	}

	return nil
}
