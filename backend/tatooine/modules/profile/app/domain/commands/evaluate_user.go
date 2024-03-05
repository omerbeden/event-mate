package commands

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"

	"github.com/omerbeden/event-mate/backend/tatooine/modules/profile/app/adapters/cachedapter"
	"github.com/omerbeden/event-mate/backend/tatooine/modules/profile/app/domain/model"
	"github.com/omerbeden/event-mate/backend/tatooine/modules/profile/app/domain/ports/repositories"
	"github.com/omerbeden/event-mate/backend/tatooine/pkg/cache"
	customerrors "github.com/omerbeden/event-mate/backend/tatooine/pkg/customErrors"
)

type EvaluateUserCommand struct {
	UserRepo   repositories.UserProfileRepository
	StatRepo   repositories.UserProfileStatRepository
	Cache      cache.Cache
	Evaluation model.UserEvaluation
}

func (cmd *EvaluateUserCommand) Handle(ctx context.Context) error {
	err := cmd.StatRepo.EvaluateUser(ctx, cmd.Evaluation)
	if err != nil {
		if errors.Is(err, customerrors.ErrDublicateKey) {
			return customerrors.ErrAlreadyEvaluated
		}
		return err
	}

	updatedUser, err := cmd.UserRepo.GetCurrentUserProfile(ctx, cmd.Evaluation.ReceiverId)
	if err != nil {
		return err
	}

	return cmd.updateCache(ctx, updatedUser)
}

func (cmd *EvaluateUserCommand) updateCache(ctx context.Context, updatedUser *model.UserProfile) error {
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
