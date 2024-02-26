package commands

import (
	"context"

	"github.com/omerbeden/event-mate/backend/tatooine/modules/profile/app/adapters/cachedapter"
	"github.com/omerbeden/event-mate/backend/tatooine/modules/profile/app/domain/model"
	"github.com/omerbeden/event-mate/backend/tatooine/modules/profile/app/domain/ports/repositories"
)

type GetUserProfileStatsCommand struct {
	Repo   repositories.UserProfileRepository
	Cache  cachedapter.Cache
	UserId int64
}

func (cmd *GetUserProfileStatsCommand) Handle(ctx context.Context) (*model.UserProfileStat, error) {
	return cmd.Repo.GetUserProfileStats(ctx, cmd.UserId)
}
