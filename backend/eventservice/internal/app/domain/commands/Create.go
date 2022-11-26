package commands

import (
	cacheadapter "github.com/omerbeden/event-mate/backend/eventservice/internal/app/adapters/cacheAdapter"
	"github.com/omerbeden/event-mate/backend/eventservice/internal/app/domain/model"
	"github.com/omerbeden/event-mate/backend/eventservice/internal/app/domain/ports/caching"
	"github.com/omerbeden/event-mate/backend/eventservice/internal/app/domain/ports/repo"
)

type CreateCommand struct {
	Event model.Event
	Repo  repo.Repository
	Redis cacheadapter.RedisAdapter
}

func (ccmd *CreateCommand) Handle() (bool, error) {

	//todo: refactor : get yapılırken key -> array pairi şeklinde çekiyoruz , ama add yapılırken bu nasıl olacak araştır
	// cache deki arraye ekleme yapılıyor mu ?
	// yapılmıyorsa her defasında
	cacheadapter.Set(ccmd.Event.Location.City, ccmd.Event, &ccmd.Redis)
	return ccmd.Repo.CreateEvent(ccmd.Event)

}

type CreateCacheCommand struct {
	Redis caching.Cache
	Key   string
	Posts []model.Event
}

func (uc *CreateCacheCommand) Handle() (bool, error) {
	err := uc.Redis.AddToCache(uc.Key, uc.Posts)
	if err != nil {
		return false, err
	}

	return true, nil
}
