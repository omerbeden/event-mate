package commands

import (
	"context"

	"github.com/omerbeden/event-mate/backend/tatooine/modules/profile/app/domain/model"
	"github.com/omerbeden/event-mate/backend/tatooine/modules/profile/app/domain/ports/repositories"
)

type CreateBadgeCommand struct {
	BadgeRepo repositories.ProfileBadgeRepository
	Badge     *model.ProfileBadge
}

func (cmd *CreateBadgeCommand) Handle(ctx context.Context) error {
	return cmd.BadgeRepo.Insert(ctx, cmd.Badge)
}
