package commands

import (
	"fmt"

	"github.com/omerbeden/event-mate/backend/tatooine/modules/profile/app/adapters/cachedapter"
	"github.com/omerbeden/event-mate/backend/tatooine/modules/profile/app/domain/ports/repositories"
)

type DeleteProfileCommand struct {
	Repo       repositories.UserProfileRepository
	Cache      cachedapter.Cache
	ExternalId string
	UserName   string
}

func (c *DeleteProfileCommand) Handle() error {
	err := c.Repo.DeleteUser(c.ExternalId)
	if err != nil {
		return err
	}

	return c.deleteFromCache()
}

func (c *DeleteProfileCommand) deleteFromCache() error {
	cacheKeyExternalId := fmt.Sprintf("%s:%s", userProfileCacheKey, c.ExternalId)
	cacheKeyUserName := fmt.Sprintf("%s:%s", userProfileCacheKey, c.UserName)

	err := c.Cache.Delete(cacheKeyExternalId)
	if err != nil {
		return err
	}

	err = c.Cache.Delete(cacheKeyUserName)
	if err != nil {
		return err
	}

	return nil
}
