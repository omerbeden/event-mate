package commands

import (
	"github.com/omerbeden/event-mate/backend/tatooine/modules/profile/app/domain/model"
	"github.com/omerbeden/event-mate/backend/tatooine/modules/profile/app/domain/ports"
)

type DeleteProfileCommand struct {
	Profile model.UserProfile
	Repo    ports.UserProfileRepository
	userId  int64
}

func (c *DeleteProfileCommand) Handle() error {
	return c.Repo.DeleteUserById(c.userId)
}
