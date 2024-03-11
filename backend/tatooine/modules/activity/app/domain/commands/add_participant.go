package commands

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/omerbeden/event-mate/backend/tatooine/modules/activity/app/adapters/cacheadapter"
	"github.com/omerbeden/event-mate/backend/tatooine/modules/activity/app/domain/model"
	"github.com/omerbeden/event-mate/backend/tatooine/modules/activity/app/domain/ports/repositories"
	"github.com/omerbeden/event-mate/backend/tatooine/pkg"
	"github.com/omerbeden/event-mate/backend/tatooine/pkg/cache"
	"go.uber.org/zap"
)

const ERR_PREFIX_ADD_PARTICIPANT = "commands:addParticipant"

type AddParticipantCommand struct {
	ActivityRepository repositories.ActivityRepository
	Redis              cache.Cache
	Participant        model.User
	ActivityId         int64
}

func (command *AddParticipantCommand) Handle(ctx context.Context) error {

	logger, ok := ctx.Value(pkg.LoggerKey).(*zap.SugaredLogger)
	if !ok {
		return fmt.Errorf("failed to get logger for AddParticipantCommand")
	}

	valueJSON, err := json.Marshal(&command.Participant)
	if err != nil {
		return fmt.Errorf("%s error while marshaling", ERR_PREFIX_ADD_PARTICIPANT)
	}
	err = command.addParticipantToRedis(ctx, command.ActivityId, valueJSON)
	if err != nil {
		logger.Infof("%s could not add participant member to redis for activity id %d, adding to db", ERR_PREFIX_ADD_PARTICIPANT, command.ActivityId)
		return command.ActivityRepository.AddParticipant(ctx, command.ActivityId, command.Participant)
	}

	err = command.addAttandedActivitiesForUser(ctx, command.Participant.ID, valueJSON)
	if err != nil {
		fmt.Printf("%s could not add atttanded activities to redis ", ERR_PREFIX_ADD_PARTICIPANT)
	}

	return command.ActivityRepository.AddParticipant(ctx, command.ActivityId, command.Participant)
}

func (command *AddParticipantCommand) addParticipantToRedis(ctx context.Context, activityID int64, valueJSON []byte) error {
	redisKey := fmt.Sprintf("%s:%d", cacheadapter.PARTICIPANT_CACHE_KEY, activityID)

	return command.Redis.AddMember(ctx, redisKey, valueJSON)

}

func (command *AddParticipantCommand) addAttandedActivitiesForUser(ctx context.Context, userId int64, valueJSON []byte) error {
	redisKey := fmt.Sprintf("%s:%d", cache.ATTANDED_ACTIVITIES_CACHE_KEY, userId)
	return command.Redis.AddMember(ctx, redisKey, valueJSON)
}
