package commands

import (
	"context"

	"github.com/omerbeden/event-mate/backend/tatooine/modules/profile/app/domain/ports/repositories"
)

type UpdateProfileImageCommand struct {
	Repo       repositories.UserProfileRepository
	ImageUrl   string
	ExternalId string
}

func (c *UpdateProfileImageCommand) Handle(ctx context.Context) error {
	return c.Repo.UpdateProfileImage(ctx, c.ExternalId, c.ImageUrl)

}
