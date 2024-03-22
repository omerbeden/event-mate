package commands

import (
	"context"

	"github.com/omerbeden/event-mate/backend/tatooine/modules/profile/app/domain/model"
	"github.com/omerbeden/event-mate/backend/tatooine/modules/profile/app/domain/ports/repositories"
	"github.com/omerbeden/event-mate/backend/tatooine/pkg/db"
)

type CreateBadgeCommand struct {
	BadgeRepo repositories.ProfileBadgeRepository
	Badge     *model.ProfileBadge
	Tx        db.Tx
}

func (cmd *CreateBadgeCommand) Handle(ctx context.Context) error {
	return cmd.BadgeRepo.Insert(ctx, cmd.Tx, cmd.Badge)
}
