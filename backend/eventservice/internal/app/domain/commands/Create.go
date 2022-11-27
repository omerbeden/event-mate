package commands

import (
	"github.com/omerbeden/event-mate/backend/eventservice/internal/app/domain/model"
	"github.com/omerbeden/event-mate/backend/eventservice/internal/app/domain/ports/caching"
	"github.com/omerbeden/event-mate/backend/eventservice/internal/app/domain/ports/repo"
)

type CreateCommand struct {
	Event model.Event
	Repo  repo.Repository
	Redis caching.Cache
}

func (ccmd *CreateCommand) Handle() (bool, error) {
	err := ccmd.Redis.Push(ccmd.Event.Location.City, ccmd.Event)
	if err != nil {
		return false, err
	}

	return ccmd.Repo.CreateEvent(ccmd.Event)

}

type CreateCacheCommand struct {
	Redis caching.Cache
	Key   string
	Posts []model.Event
}

func (uc *CreateCacheCommand) Handle() (bool, error) {
	err := uc.Redis.Push(uc.Key, uc.Posts)
	if err != nil {
		return false, err
	}

	return true, nil
}
