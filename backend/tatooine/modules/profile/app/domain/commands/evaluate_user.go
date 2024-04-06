package commands

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"

	"github.com/omerbeden/event-mate/backend/tatooine/modules/profile/app/adapters/cachedapter"
	"github.com/omerbeden/event-mate/backend/tatooine/modules/profile/app/domain/model"
	"github.com/omerbeden/event-mate/backend/tatooine/modules/profile/app/domain/ports/repositories"
	"github.com/omerbeden/event-mate/backend/tatooine/pkg"
	"github.com/omerbeden/event-mate/backend/tatooine/pkg/cache"
	customerrors "github.com/omerbeden/event-mate/backend/tatooine/pkg/customErrors"
	"go.uber.org/zap"
)

type EvaluateUserCommand struct {
	UserRepo   repositories.UserProfileRepository
	StatRepo   repositories.UserProfileStatRepository
	Cache      cache.Cache
	Evaluation model.UserEvaluation
}

func (cmd *EvaluateUserCommand) Handle(ctx context.Context) (*model.UserProfile, error) {
	logger, ok := ctx.Value(pkg.LoggerKey).(*zap.SugaredLogger)
	if !ok {
		return nil, customerrors.ErrGetLogger
	}

	err := cmd.StatRepo.EvaluateUser(ctx, cmd.Evaluation)
	if err != nil {
		if errors.Is(err, customerrors.ErrDublicateKey) {
			return nil, customerrors.ErrAlreadyEvaluated
		}
		return nil, err
	}

	user, err := cmd.UserRepo.GetCurrentUserProfile(ctx, cmd.Evaluation.ReceiverId)
	if err != nil {
		return nil, err
	}

	err = cmd.updateCache(ctx, user)
	if err != nil {
		logger.Infof("error whale updating cache %w", err)
	}

	return user, nil
}

func (cmd *EvaluateUserCommand) updateCache(ctx context.Context, user *model.UserProfile) error {
	jsonValue, err := json.Marshal(user)
	if err != nil {
		return fmt.Errorf("parsing json error %w", err)
	}

	cacheKeyExternalId := fmt.Sprintf("%s:%s", cachedapter.USER_PROFILE_CACHE_KEY, user.ExternalId)
	cacheKeyUserName := fmt.Sprintf("%s:%s", cachedapter.USER_PROFILE_CACHE_KEY, user.Header.UserName)

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
