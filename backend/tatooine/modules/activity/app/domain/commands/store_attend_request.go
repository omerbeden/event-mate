package commands

import (
	"context"

	"github.com/omerbeden/event-mate/backend/tatooine/modules/activity/app/domain/model"
	"github.com/omerbeden/event-mate/backend/tatooine/modules/activity/app/domain/ports/repositories"
)

type StoreAttendRequest struct {
	ActivityRepository repositories.ActivityRepository
	Request            model.AttendRequest
}

func (cmd *StoreAttendRequest) Handle(ctx context.Context) error {
	return cmd.ActivityRepository.StoreAttendRequest(ctx, cmd.Request)
}
