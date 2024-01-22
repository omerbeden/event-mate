package commands

import (
	"encoding/json"
	"fmt"

	"github.com/omerbeden/event-mate/backend/tatooine/modules/profile/app/adapters/cachedapter"
	"github.com/omerbeden/event-mate/backend/tatooine/modules/profile/app/domain/model"
	"github.com/omerbeden/event-mate/backend/tatooine/modules/profile/app/domain/ports/repositories"
)

type GiveUserPointCommand struct {
	Repo             repositories.UserProfileRepository
	Cache            cachedapter.Cache
	ReceiverUserName string
	Point            float32
	ExternalId       string
}

func (cmd *GiveUserPointCommand) Handle() error {
	err := cmd.Repo.UpdateProfilePoints(cmd.ReceiverUserName, cmd.Point)
	if err != nil {
		return err
	}

	updatedUser, err := cmd.Repo.GetUserProfile(cmd.ReceiverUserName)
	if err != nil {
		return err
	}

	return cmd.updateCache(updatedUser)
}

func (cmd *GiveUserPointCommand) updateCache(updatedUser *model.UserProfile) error {
	jsonValue, err := json.Marshal(updatedUser)
	if err != nil {
		return fmt.Errorf("parsing json error %w", err)
	}

	cacheKeyExternalId := fmt.Sprintf("%s:%s", userProfileCacheKey, updatedUser.ExternalId)
	cacheKeyUserName := fmt.Sprintf("%s:%s", userProfileCacheKey, updatedUser.UserName)

	err = cmd.Cache.Set(cacheKeyExternalId, jsonValue)
	if err != nil {
		return err
	}

	err = cmd.Cache.Set(cacheKeyUserName, jsonValue)
	if err != nil {
		return err
	}

	return nil
}
