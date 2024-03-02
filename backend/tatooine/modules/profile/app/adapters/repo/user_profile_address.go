package repo

import (
	"context"
	"fmt"

	"github.com/omerbeden/event-mate/backend/tatooine/modules/profile/app/domain/model"
	"github.com/omerbeden/event-mate/backend/tatooine/pkg/db"
)

type userProfileAddressRepo struct {
	pool db.DBExecutor
}

func NewUserProfileAddressRepo(pool db.DBExecutor) *userProfileRepo {
	return &userProfileRepo{
		pool: pool,
	}
}

func (r *userProfileAddressRepo) Insert(ctx context.Context, address model.UserProfileAdress) error {
	q := `INSERT INTO user_profile_addresses (profile_id,city) Values($1,$2)`
	_, err := r.pool.Exec(ctx, q, address.ProfileId, address.City)
	if err != nil {
		return fmt.Errorf("%s could not insert profile adress %w", errlogprefix, err)
	}

	return nil
}
