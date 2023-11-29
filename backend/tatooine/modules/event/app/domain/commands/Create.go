package commands

import (
	"strconv"

	"github.com/omerbeden/event-mate/backend/tatooine/modules/event/app/domain/model"
	"github.com/omerbeden/event-mate/backend/tatooine/modules/event/app/domain/ports/caching"
	repo "github.com/omerbeden/event-mate/backend/tatooine/modules/event/app/domain/ports/repositories"
)

type CreateCommand struct {
	Event     model.Event
	EventRepo repo.EventRepository
	LocRepo   repo.LocationRepository
	Redis     caching.Cache
}

func (ccmd *CreateCommand) Handle() (bool, error) {

	event, errCreate := ccmd.EventRepo.Create(ccmd.Event)
	if errCreate != nil {
		return false, errCreate
	}

	_, errLoc := ccmd.LocRepo.Create(&ccmd.Event.Location)
	if errLoc != nil {
		return false, errLoc
	}

	eventId := strconv.FormatInt(event.ID, 10)
	err := ccmd.Redis.Set(eventId, ccmd.Event)
	if err != nil {
		return false, err
	}

	return true, nil

}
