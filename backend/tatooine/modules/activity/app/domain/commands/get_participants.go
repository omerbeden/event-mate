package commands

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"

	"github.com/omerbeden/event-mate/backend/tatooine/modules/activity/app/adapters/cacheadapter"
	"github.com/omerbeden/event-mate/backend/tatooine/modules/activity/app/domain/model"
	"github.com/omerbeden/event-mate/backend/tatooine/modules/activity/app/domain/ports/repositories"
	"github.com/omerbeden/event-mate/backend/tatooine/pkg/cache"
	customerrors "github.com/omerbeden/event-mate/backend/tatooine/pkg/customErrors"
)

const ERR_PREFIX_GET_PARTICIPANTS = "commands:getparticipants"

type GetParticipantsCommand struct {
	ActivityRepository repositories.ActivityRepository
	Redis              cache.Cache
	ActivityId         int64
}

func (command *GetParticipantsCommand) Handle(ctx context.Context) ([]model.User, error) {
	redisResult, err := command.getFromRedis(ctx)
	if err != nil {
		fmt.Printf("%s redis error returning from db", ERR_PREFIX_GET_PARTICIPANTS)
		participants, err := command.ActivityRepository.GetParticipants(ctx, command.ActivityId)
		if err != nil {
			if errors.As(err, &customerrors.ErrActivityDoesNotHaveParticipants) {
				return nil, nil
			} else {
				return nil, err
			}
		}
		return participants, nil
	}

	return redisResult, nil
}

func (command *GetParticipantsCommand) getFromRedis(ctx context.Context) ([]model.User, error) {
	redisKey := fmt.Sprintf("%s:%d", cacheadapter.PARTICIPANT_CACHE_KEY, command.ActivityId)

	redisResult, err := command.Redis.GetMembers(ctx, redisKey)
	if err != nil || len(redisResult) < 1 {
		return nil, fmt.Errorf("%s could not get participants from redis , %d, %w", ERR_PREFIX_GET_PARTICIPANTS, command.ActivityId, err)
	}

	var participants []model.User

	for _, res := range redisResult {
		var participant model.User
		err := json.Unmarshal([]byte(res), &participant)
		if err != nil {
			return nil, fmt.Errorf("%s could not unmarshal participant from redis , %d , %w", ERR_PREFIX_GET_PARTICIPANTS, command.ActivityId, err)
		}
		participants = append(participants, participant)
	}

	fmt.Println(participants)
	return participants, nil
}
