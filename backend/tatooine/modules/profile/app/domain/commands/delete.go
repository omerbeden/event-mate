package commands

import (
	"fmt"
	"strconv"

	"github.com/omerbeden/event-mate/backend/tatooine/modules/profile/app/adapters/cachedapter"
	"github.com/omerbeden/event-mate/backend/tatooine/modules/profile/app/domain/ports/repositories"
)

var errLogPrefixDeleteCommand = "profile:createCommand"

type DeleteProfileCommand struct {
	Repo   repositories.UserProfileRepository
	Cache  cachedapter.Cache
	UserId int64
}

func (c *DeleteProfileCommand) Handle() error {
	err := c.Repo.DeleteUserById(c.UserId)
	userId := strconv.FormatInt(c.UserId, 10)
	if err != nil {
		return err
	}

	err = c.deleteFromCache(userId)
	if err != nil {
		return fmt.Errorf("%s error while deleting user from cache,id:  %s", errLogPrefixDeleteCommand, userId)
	}

	return nil
}

func (c *DeleteProfileCommand) deleteFromCache(key string) error {
	cacheKey := fmt.Sprintf("%s:%s", userProfileCacheKey, key)
	return c.Cache.Delete(cacheKey)
}
