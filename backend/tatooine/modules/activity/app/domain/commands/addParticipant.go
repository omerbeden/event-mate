package commands

import (
	"encoding/json"
	"fmt"

	"github.com/omerbeden/event-mate/backend/tatooine/modules/activity/app/domain/model"
	"github.com/omerbeden/event-mate/backend/tatooine/modules/activity/app/domain/ports/caching"
	"github.com/omerbeden/event-mate/backend/tatooine/modules/activity/app/domain/ports/repositories"
)

const PARTICIPANT_REDIS_KEY = "participant"
const ERR_PREFIX_ADD_PARTICIPANT = "commands:addParticipant"

type AddParticipantCommand struct {
	ActivityRepository repositories.ActivityRepository
	Redis              caching.Cache
	Participant        model.User
	ActivityId         int64
}

func (command *AddParticipantCommand) Handle() error {

	valueJSON, err := json.Marshal(&command.Participant)
	if err != nil {
		return fmt.Errorf("%s error while marshaling", ERR_PREFIX_ADD_PARTICIPANT)
	}
	err = command.addParticipantToRedis(command.ActivityId, valueJSON)
	if err != nil {
		return fmt.Errorf("%s could not add participant member to redis for activity id %d", ERR_PREFIX_ADD_PARTICIPANT, command.ActivityId)
	}

	return command.ActivityRepository.AddParticipant(command.ActivityId, command.Participant)
}

func (command *AddParticipantCommand) addParticipantToRedis(activityID int64, valueJSON []byte) error {
	redisKey := fmt.Sprintf("%s:%d", PARTICIPANT_REDIS_KEY, activityID)

	return command.Redis.AddMember(redisKey, valueJSON)

}
