package commands

import "github.com/omerbeden/event-mate/backend/tatooine/modules/profile/app/domain/ports"

type UpdateProfileImageCommand struct {
	Repo     ports.UserProfileRepository
	ImageUrl string
	UserId   int64
}

func (c *UpdateProfileImageCommand) Handle() error {
	return c.Repo.UpdateProfileImage(c.UserId, c.ImageUrl)
}
