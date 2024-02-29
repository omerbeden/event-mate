package commands

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/omerbeden/event-mate/backend/tatooine/modules/profile/app/adapters/cachedapter"
	"github.com/omerbeden/event-mate/backend/tatooine/modules/profile/app/domain/model"
	"github.com/omerbeden/event-mate/backend/tatooine/modules/profile/app/domain/ports/repositories"
	"github.com/omerbeden/event-mate/backend/tatooine/pkg/cache"
)

type GiveUserPointCommand struct {
	Repo             repositories.UserProfileRepository
	Cache            cache.Cache
	ReceiverUserName string
	Point            float32
	ExternalId       string
}

func (cmd *GiveUserPointCommand) Handle(ctx context.Context) error {
	err := cmd.Repo.UpdateProfilePoints(ctx, cmd.ReceiverUserName, cmd.Point)
	if err != nil {
		return err
	}

	updatedUser, err := cmd.Repo.GetUserProfile(ctx, cmd.ReceiverUserName)
	if err != nil {
		return err
	}

	return cmd.updateCache(ctx, updatedUser)
}

func (cmd *GiveUserPointCommand) updateCache(ctx context.Context, updatedUser *model.UserProfile) error {
	jsonValue, err := json.Marshal(updatedUser)
	if err != nil {
		return fmt.Errorf("parsing json error %w", err)
	}

	cacheKeyExternalId := fmt.Sprintf("%s:%s", cachedapter.USER_PROFILE_CACHE_KEY, updatedUser.ExternalId)
	cacheKeyUserName := fmt.Sprintf("%s:%s", cachedapter.USER_PROFILE_CACHE_KEY, updatedUser.UserName)

	err = cmd.Cache.Set(ctx, cacheKeyExternalId, jsonValue)
	if err != nil {
		return err
	}

	err = cmd.Cache.Set(ctx, cacheKeyUserName, jsonValue)
	if err != nil {
		return err
	}

	return nil
}
