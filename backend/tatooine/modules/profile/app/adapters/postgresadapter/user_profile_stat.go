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

func (r *userProfileStatRepo) EvaluateUser(ctx context.Context, eval model.UserEvaluation) error {

	tx, err := r.pool.Begin(ctx)

	defer func() {
		if err != nil {
			tx.Rollback(ctx)
		}
	}()

	if err != nil {
		return fmt.Errorf("failed to begin transaction: %w", err)
	}
	iq := `INSERT INTO user_points(giver_id,receiver_id,points,comment) 
		VALUES ($1,$2,$3,$4);`

	_, err = tx.Exec(ctx, iq, eval.GiverId, eval.ReceiverId, eval.Points, eval.Comment)

	if err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) {
			if pgErr.Code == "23505" {
				return customerrors.ErrDublicateKey
			}
		}
		return fmt.Errorf("could not evaluate user %s , %w", eval.ReceiverId, err)
	}

	uq := `WITH average_points AS (
				SELECT
					up.id AS profile_id,
					ROUND(AVG(points),1) AS mean_point
				FROM
					user_points upoints
				JOIN
					user_profiles up ON up.external_id = upoints.receiver_id
				GROUP BY
					up.id
			)
			UPDATE
				user_profile_stats ups
			SET
			average_point = ap.mean_point
			FROM
				average_points ap
			WHERE
				ups.profile_id = ap.profile_id;`

	_, err = tx.Exec(ctx, uq)
	if err != nil {
		return fmt.Errorf("failed to calculate average point")
	}

	err = tx.Commit(ctx)
	if err != nil {
		return fmt.Errorf("failed to commit transaction: %w", err)
	}

	return nil
}
