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

func (r *userProfileStatRepo) UpdateProfilePoints(ctx context.Context, receiverUserName string, point float32) error {

	q := `UPDATE user_profile_stats
		SET point = point + $1,
		point_giving_count = point_giving_count + 1
		FROM user_profiles
		WHERE user_profile_stas.profile_id = user_profiles.id
		AND user_profiles.user_name = $2`

	_, err := r.pool.Exec(ctx, q, point, receiverUserName)
	if err != nil {
		return fmt.Errorf("%s could not update user %s , %w", errlogprefix, receiverUserName, err)
	}

	return nil
}
