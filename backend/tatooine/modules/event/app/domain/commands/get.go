package commands

import (
	"fmt"
	"strconv"

	"github.com/omerbeden/event-mate/backend/tatooine/modules/event/app/adapters/redisadapter"
	"github.com/omerbeden/event-mate/backend/tatooine/modules/event/app/domain/model"
	repo "github.com/omerbeden/event-mate/backend/tatooine/modules/event/app/domain/ports/repositories"
)

type GetCommand struct {
	EventID   string
	EventCity string
	Repo      repo.EventRepository
	Redis     *redisadapter.RedisAdapter
}

func (gc *GetCommand) Handle() (*model.Event, error) {

	intID, err := strconv.Atoi(gc.EventID)
	if err != nil {
		return nil, fmt.Errorf("get command: %w", err)
	}

	result, redisErr := gc.Redis.Get(gc.EventID)
	if redisErr != nil {
		return nil, redisErr
	}

	if result != nil {
		return result.(*model.Event), err
	}

	return gc.Repo.GetByID(int64(intID))

}
