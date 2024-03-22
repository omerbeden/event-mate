package commands

import (
	"context"

	"github.com/omerbeden/event-mate/backend/tatooine/modules/profile/app/domain/ports/repositories"
)

type UpdateVerificationCommand struct {
	Repo       repositories.UserProfileRepository
	IsVerified bool
	ExternalId string
}

func (cmd *UpdateVerificationCommand) Handle(ctx context.Context) error {
	return cmd.Repo.UpdateVerification(ctx, cmd.ExternalId, cmd.IsVerified)
}
