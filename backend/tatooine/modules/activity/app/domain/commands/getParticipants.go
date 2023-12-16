package commands

import (
	"encoding/json"
	"fmt"

	"github.com/omerbeden/event-mate/backend/tatooine/modules/activity/app/domain/model"
	"github.com/omerbeden/event-mate/backend/tatooine/modules/activity/app/domain/ports/caching"
	"github.com/omerbeden/event-mate/backend/tatooine/modules/activity/app/domain/ports/repositories"
)

const ERR_PREFIX_GET_PARTICIPANTS = "commands:getparticipants"

type GetParticipantsCommand struct {
	ActivityRepository repositories.ActivityRepository
	Redis              caching.Cache
	ActivityId         int64
}

func (command *GetParticipantsCommand) Handle() ([]model.User, error) {
	redisResult, err := command.getFromRedis()
	if err != nil {
		fmt.Printf("%s redis error returning from db", ERR_PREFIX_GET_PARTICIPANTS)
		return command.ActivityRepository.GetParticipants(command.ActivityId)
	}

	return redisResult, nil
}

func (command *GetParticipantsCommand) getFromRedis() ([]model.User, error) {
	redisKey := fmt.Sprintf("%s:%d", PARTICIPANT_REDIS_KEY, command.ActivityId)

	redisResult, err := command.Redis.GetMembers(redisKey)
	if err != nil {
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

	return participants, nil
}
