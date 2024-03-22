package postgresadapter

import (
	"context"
	"fmt"

	"github.com/omerbeden/event-mate/backend/tatooine/modules/profile/app/domain/model"
	"github.com/omerbeden/event-mate/backend/tatooine/pkg/db"
)

type BadgeRepo struct {
	pool db.Executor
}

func NewBadgeRepo(pool db.Executor) *BadgeRepo {
	return &BadgeRepo{
		pool: pool,
	}
}

func (r *BadgeRepo) Insert(ctx context.Context, tx db.Tx, badge *model.ProfileBadge) error {

	q := `INSERT INTO profile_badges(profile_id,badge_id,image_url,text) VALUES($1,$2,$3,$4)`

	if tx != nil {
		_, err := tx.Exec(ctx, q, badge.ProfileId, badge.BadgeId, badge.ImageUrl, badge.Text)

		if err != nil {
			return fmt.Errorf("could not insert badge %w", err)
		}
	} else {
		_, err := r.pool.Exec(ctx, q, badge.ProfileId, badge.BadgeId, badge.ImageUrl, badge.Text)

		if err != nil {
			return fmt.Errorf("could not insert badge %w", err)
		}
	}

	return nil
}

func (r *BadgeRepo) GetBadges(ctx context.Context, profileId int64) (map[int64]*model.ProfileBadge, error) {

	q := `SELECT badge_id,profile_id,image_url,text from profile_badges 
	WHERE profile_id = $1`

	rows, err := r.pool.Query(ctx, q, profileId)
	if err != nil {
		return nil, fmt.Errorf("could not get badges %w", err)
	}

	badges := make(map[int64]*model.ProfileBadge)
	for rows.Next() {
		var badge model.ProfileBadge
		err := rows.Scan(&badge.BadgeId, &badge.ProfileId, &badge.ImageUrl, &badge.Text)
		if err != nil {
			return nil, fmt.Errorf("could not get rows of  badges %w", err)
		}
		badges[badge.BadgeId] = &badge
	}

	return badges, nil
}
