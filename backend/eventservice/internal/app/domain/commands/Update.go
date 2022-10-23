package commands

import (
	"strconv"

	"github.com/omerbeden/event-mate/backend/eventservice/internal/app/domain/model"
	"github.com/omerbeden/event-mate/backend/eventservice/internal/app/domain/ports/caching"
	"github.com/omerbeden/event-mate/backend/eventservice/internal/app/domain/ports/repo"
	"github.com/omerbeden/event-mate/backend/eventservice/internal/infra/grpc/pb"
)

type UpdateCommand struct {
	Event *pb.Event
	Repo  repo.Repository
}

func (uc *UpdateCommand) Handle() (bool, error) {
	model := &model.Event{
		Title:    uc.Event.Title,
		Category: uc.Event.Category,
	}

	intID, _ := strconv.Atoi(uc.Event.GetId())
	return uc.Repo.UpdateEventByID(int32(intID), *model)
}

type UpdateCacheCommand struct {
	Redis caching.Cache
	Key   string
	Posts []model.Event
}

func (uc *UpdateCacheCommand) Handle() (bool, error) {
	err := uc.Redis.UpdateCache(uc.Key, uc.Posts)
	if err != nil {
		return false, err
	}

	return true, nil
}
