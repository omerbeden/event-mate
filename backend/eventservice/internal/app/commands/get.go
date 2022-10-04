package commands

import (
	"fmt"
	"strconv"

	"github.com/omerbeden/event-mate/backend/eventservice/internal/app/domain/model"
	"github.com/omerbeden/event-mate/backend/eventservice/internal/app/ports/repo"
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
