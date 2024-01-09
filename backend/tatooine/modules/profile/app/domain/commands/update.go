package commands

import (
	"encoding/json"
	"fmt"
	"strconv"

	"github.com/omerbeden/event-mate/backend/tatooine/modules/profile/app/adapters/cachedapter"
	"github.com/omerbeden/event-mate/backend/tatooine/modules/profile/app/domain/model"
	"github.com/omerbeden/event-mate/backend/tatooine/modules/profile/app/domain/ports/repositories"
)

type UpdateProfileImageCommand struct {
	Repo  repositories.UserProfileRepository
	Cache cachedapter.Cache

	ImageUrl string
	UserId   int64
}

func (c *UpdateProfileImageCommand) Handle() error {
	updatedUser, err := c.Repo.UpdateProfileImage(c.UserId, c.ImageUrl)
	if err != nil {
		return err
	}

	fmt.Printf(" updated : %+v\n", updatedUser)

	return c.updateCache(c.UserId, updatedUser)
}

func (c *UpdateProfileImageCommand) updateCache(userId int64, updatedUser *model.UserProfile) error {
	userIdStr := strconv.FormatInt(userId, 10)
	cacheKey := fmt.Sprintf("%s:%s", userProfileCacheKey, userIdStr)

	jsonValue, err := json.Marshal(updatedUser)
	if err != nil {
		return fmt.Errorf("parsing error while updating user profile on cache")
	}

	return c.Cache.Set(cacheKey, jsonValue)
}
