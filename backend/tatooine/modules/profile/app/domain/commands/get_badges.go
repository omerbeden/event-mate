package commands

import (
	"context"

	"github.com/omerbeden/event-mate/backend/tatooine/modules/profile/app/domain/model"
	"github.com/omerbeden/event-mate/backend/tatooine/modules/profile/app/domain/ports/repositories"
)

type GetProfileBadgesCommand struct {
	BadgeRepo   repositories.ProfileBadgeRepository
	ProfileRepo repositories.UserProfileRepository
	ExternalId  string
}

func (cmd *GetProfileBadgesCommand) Handle(ctx context.Context) (map[int64]*model.ProfileBadge, error) {
	id, err := cmd.ProfileRepo.GetId(ctx, cmd.ExternalId)
	if err != nil {
		return nil, err
	}

	return cmd.BadgeRepo.GetBadges(ctx, id)
}
