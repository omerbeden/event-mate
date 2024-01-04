package commands

import (
	"fmt"
	"strconv"

	"github.com/omerbeden/event-mate/backend/tatooine/modules/profile/app/adapters/cachedapter"
	"github.com/omerbeden/event-mate/backend/tatooine/modules/profile/app/domain/model"
	"github.com/omerbeden/event-mate/backend/tatooine/modules/profile/app/domain/ports"
)

var errLogPrefixDeleteCommand = "profile:createCommand"

type DeleteProfileCommand struct {
	Profile model.UserProfile
	Repo    ports.UserProfileRepository
	Cache   cachedapter.Cache
	userId  int64
}

func (c *DeleteProfileCommand) Handle() error {
	err := c.Repo.DeleteUserById(c.userId)
	userId := strconv.FormatInt(c.userId, 10)
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
	return c.Cache.Delete(key)
}
