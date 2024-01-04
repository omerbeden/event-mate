package commands

import (
	"strconv"

	"github.com/omerbeden/event-mate/backend/tatooine/modules/profile/app/adapters/cachedapter"
	"github.com/omerbeden/event-mate/backend/tatooine/modules/profile/app/domain/ports"
)

type UpdateProfileImageCommand struct {
	Repo  ports.UserProfileRepository
	Cache cachedapter.Cache

	ImageUrl string
	UserId   int64
}

func (c *UpdateProfileImageCommand) Handle() error {
	err := c.Repo.UpdateProfileImage(c.UserId, c.ImageUrl)
	if err != nil {
		return err
	}

	return c.updateCache(c.UserId, c.ImageUrl)
}

func (c *UpdateProfileImageCommand) updateCache(userId int64, imageURl string) error {
	userIdStr := strconv.FormatInt(userId, 10)
	return c.Cache.Set(userIdStr, []byte(imageURl))
}
