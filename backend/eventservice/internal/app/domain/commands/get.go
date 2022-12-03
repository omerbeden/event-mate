package commands

import (
	"fmt"
	"strconv"

	cacheadapter "github.com/omerbeden/event-mate/backend/eventservice/internal/app/adapters/cacheAdapter"
	"github.com/omerbeden/event-mate/backend/eventservice/internal/app/domain/model"
	"github.com/omerbeden/event-mate/backend/eventservice/internal/app/domain/ports/repo"
)

type GetCommand struct {
	EventID   string
	EventCity string
	Repo      repo.Repository
	Redis     *cacheadapter.RedisAdapter
}

func (gc *GetCommand) Handle() (model.Event, error) {

	intID, err := strconv.Atoi(gc.EventID)
	if err != nil {
		return model.Event{}, fmt.Errorf("get command: %w", err)
	}

	isCacheExist, err := cacheadapter.Exist(gc.EventCity, gc.Redis)
	if err != nil {
		return model.Event{}, err
	}
	if isCacheExist {
		return cacheadapter.GetEvent(intID, gc.EventCity, gc.Redis)

	} else {
		return gc.Repo.GetEventByID(int32(intID))
	}
}

type GetFeedCommand struct {
	Repo     repo.Repository
	Location *model.Location
	Redis    *cacheadapter.RedisAdapter
}

func (gf *GetFeedCommand) Handle() (*model.GetFeedCommandResult, error) {

	isCacheExist, err := cacheadapter.Exist(gf.Location.City, gf.Redis)
	if err != nil {
		return nil, err
	}

	if isCacheExist {
		cacheResult, cacheErr := cacheadapter.GetPosts(gf.Location.City, gf.Redis)
		if cacheErr != nil {
			return nil, cacheErr
		}
		return &model.GetFeedCommandResult{Events: &cacheResult, CacheHit: true}, nil
	} else {
		events, err := gf.Repo.GetEventByLocation(gf.Location)
		if err != nil {
			fmt.Println("Errror ocurred")
			return nil, err
		}
		return &model.GetFeedCommandResult{Events: &events, CacheHit: false}, nil
	}
}
