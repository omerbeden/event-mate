package commands

import (
	"encoding/json"
	"fmt"

	"github.com/omerbeden/event-mate/backend/tatooine/modules/profile/app/adapters/cachedapter"
	"github.com/omerbeden/event-mate/backend/tatooine/modules/profile/app/domain/ports/repositories"
)

type GiveUserPointCommand struct {
	Repo       repositories.UserProfileRepository
	Cache      cachedapter.Cache
	ReceiverId int64
	Point      float32
}

func (cmd *GiveUserPointCommand) Handle() error {
	err := cmd.Repo.UpdateProfilePoints(cmd.ReceiverId, cmd.Point)
	if err != nil {
		return err
	}

	updatedUser, err := cmd.Repo.GetUserProfile(cmd.ReceiverId)
	if err != nil {
		return err
	}

	jsonValue, err := json.Marshal(updatedUser)
	if err != nil {
		return fmt.Errorf("parsing json error %w", err)
	}

	return cmd.updateCache(cmd.ReceiverId, jsonValue)
}

func (cmd *GiveUserPointCommand) updateCache(userId int64, value []byte) error {
	cacheKey := fmt.Sprintf("%s:%d", userProfileCacheKey, userId)
	return cmd.Cache.Set(cacheKey, value)
}
