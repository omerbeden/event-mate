package commands

import (
	"context"

	"github.com/omerbeden/event-mate/backend/tatooine/modules/profile/app/domain/model"
	"github.com/omerbeden/event-mate/backend/tatooine/modules/profile/app/domain/ports/repositories"
)

type GetEvaluationsCommand struct {
	Repo repositories.UserProfileStatRepository
}

func (cmd *GetEvaluationsCommand) Handle(ctx context.Context, userId int64) ([]model.GetUserEvaluations, error) {
	evaluations, err := cmd.Repo.GetEvaluations(ctx, userId)
	if err != nil {
		return nil, err
	}
	return evaluations, nil
}
