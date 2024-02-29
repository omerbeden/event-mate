package commands

import (
	"context"

	"github.com/omerbeden/event-mate/backend/tatooine/modules/profile/app/domain/model"
	"github.com/omerbeden/event-mate/backend/tatooine/modules/profile/app/domain/ports/repositories"
	"github.com/omerbeden/event-mate/backend/tatooine/pkg/cache"
)

type GetUserProfileStatsCommand struct {
	Repo   repositories.UserProfileRepository
	Cache  cache.Cache
	UserId int64
}

func (cmd *GetUserProfileStatsCommand) Handle(ctx context.Context) (*model.UserProfileStat, error) {
	return cmd.Repo.GetUserProfileStats(ctx, cmd.UserId)
}
