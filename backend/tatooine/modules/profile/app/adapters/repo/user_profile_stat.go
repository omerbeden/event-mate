package repo

import (
	"context"
	"fmt"

	"github.com/omerbeden/event-mate/backend/tatooine/modules/profile/app/domain/model"
	"github.com/omerbeden/event-mate/backend/tatooine/pkg/db"
)

type userProfileStatRepo struct {
	pool db.DBExecutor
}

func NewUserProfileStatRepo(pool db.DBExecutor) *userProfileRepo {
	return &userProfileRepo{
		pool: pool,
	}
}

func (r *userProfileStatRepo) Insert(ctx context.Context, user model.UserProfile) error {
	q := fmt.Sprintf(
		`INSERT INTO user_profile_stats
		 (profile_id, point, attanded_activities)
		 Values(%d,%f,%d)`, user.Id, user.Stat.Point, user.Stat.AttandedActivities)
	_, err := r.pool.Exec(ctx, q)
	if err != nil {
		return fmt.Errorf("%s could not insert profile stats %w", errlogprefix, err)
	}

	return nil
}
