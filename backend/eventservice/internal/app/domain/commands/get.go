package commands

import (
	"fmt"
	"strconv"

	cacheadapter "github.com/omerbeden/event-mate/backend/eventservice/internal/app/adapters/cacheAdapter"
	"github.com/omerbeden/event-mate/backend/eventservice/internal/app/domain/model"
	"github.com/omerbeden/event-mate/backend/eventservice/internal/app/domain/ports/repo"
)

type GetCommand struct {
	EventID string
	Repo    repo.Repository
}

func (gc *GetCommand) Handle() (model.Event, error) {

	intID, err := strconv.Atoi(gc.EventID)
	if err != nil {
		fmt.Printf("Err: TODO")
	}
	result, err := gc.Repo.GetEventByID(int32(intID))
	if err != nil {
		fmt.Printf("Err: TODO")
	}

	return result, nil
}

type GetFeedCommand struct {
	Repo     repo.Repository
	Location *model.Location
	Redis    *cacheadapter.RedisAdapter
}

func (gf *GetFeedCommand) Handle() (*model.GetFeedCommandResult, error) {

	isExist := cacheadapter.Exist(gf.Location.City, gf.Redis)

	if isExist {
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
		return &model.GetFeedCommandResult{Events: &events, CacheHit: true}, nil
	}
}
