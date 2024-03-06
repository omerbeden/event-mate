package postgresadapter

import (
	"context"
	"errors"
	"fmt"

	"github.com/jackc/pgx/v5/pgconn"
	"github.com/omerbeden/event-mate/backend/tatooine/modules/profile/app/domain/model"
	customerrors "github.com/omerbeden/event-mate/backend/tatooine/pkg/customErrors"
	"github.com/omerbeden/event-mate/backend/tatooine/pkg/db"
)

type userProfileStatRepo struct {
	pool db.Executor
}

func NewUserProfileStatRepo(pool db.Executor) *userProfileStatRepo {
	return &userProfileStatRepo{
		pool: pool,
	}
}

func (r *userProfileStatRepo) Insert(ctx context.Context, tx db.Tx, stat model.UserProfileStat) error {
	q := fmt.Sprintf(
		`INSERT INTO user_profile_stats
		 (profile_id, point, attanded_activities)
		 Values(%d,%f,%d)`, stat.ProfileId, stat.Point, stat.AttandedActivities)
	_, err := tx.Exec(ctx, q)
	if err != nil {
		return fmt.Errorf("could not insert profile stats %w", err)
	}

	return nil
}

func (r *userProfileStatRepo) EvaluateUser(ctx context.Context, eval model.UserEvaluation) error {

	q := `INSERT INTO user_points(giver_id,receiver_id,points,comment) 
		VALUES ($1,$2,$3,$4);`

	_, err := r.pool.Exec(ctx, q, eval.GiverId, eval.ReceiverId, eval.Points, eval.Comment)

	if err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) {
			if pgErr.Code == "23505" {
				return customerrors.ErrDublicateKey
			}
		}
		return fmt.Errorf("could not evaluate user %s , %w", eval.ReceiverId, err)
	}

	return nil
}
