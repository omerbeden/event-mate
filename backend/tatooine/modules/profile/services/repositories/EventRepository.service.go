package repositories

import (
	"github.com/omerbeden/event-mate/backend/tatooine/modules/profile/core"
	"github.com/omerbeden/event-mate/backend/tatooine/modules/profile/infrastructure/repositories"
)

type EventRepoService struct {
	Repo repositories.EventRepo
}

func (r *EventRepoService) GetUserEvent(userID int) ([]core.Event, error) {
	return r.Repo.GetUserEvent(userID)
}
