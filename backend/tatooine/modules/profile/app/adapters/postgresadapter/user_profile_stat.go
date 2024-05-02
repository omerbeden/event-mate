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
	iq := `INSERT INTO user_points(giver_id,receiver_id,points,comment,related_activity_id) 
		VALUES ($1,$2,$3,$4,$5);`

	_, err = tx.Exec(ctx, iq, eval.GiverId, eval.ReceiverId, eval.Points, eval.Comment, eval.RelatedActivityId)
	fmt.Println(err)
	if err != nil {
		fmt.Print("error ocurred ")
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) {
			if pgErr.Code == "23505" {
				return customerrors.ErrDublicateKey
			}
		}
		return fmt.Errorf("could not evaluate user %d , %w", eval.ReceiverId, err)
	}

	uq := `WITH average_points AS (
				SELECT
					up.id AS profile_id,
					ROUND(AVG(points),1) AS mean_point
				FROM
					user_points upoints
				JOIN
					user_profiles up ON up.id = upoints.receiver_id
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
		return fmt.Errorf("failed to calculate average point %w", err)
	}

	err = tx.Commit(ctx)
	if err != nil {
		return fmt.Errorf("failed to commit transaction: %w", err)
	}

	return nil
}

func (r *userProfileStatRepo) GetEvaluations(ctx context.Context, userId int64) ([]model.GetUserEvaluations, error) {
	q := `	SELECT
		giver.user_name,
		giver.profile_image_url,
		pts.given_at,
		pts.comment,
		pts.points
	FROM user_profiles p
	JOIN user_points pts ON p.id = pts.receiver_id
	JOIN user_profiles giver ON  giver.id = pts.giver_id
	WHERE p.id = $1;`

	rows, err := r.pool.Query(ctx, q, userId)
	if err != nil {
		return nil, fmt.Errorf("could not get evaluations for user: %d %w", userId, err)
	}

	var evaluations []model.GetUserEvaluations
	for rows.Next() {
		var eval model.GetUserEvaluations
		err := rows.Scan(&eval.GiverUserName, &eval.GiverProfileImageUrl, &eval.GivenAt, &eval.Comment, &eval.GivenPoint)
		if err != nil {
			return nil, fmt.Errorf("error getting rows  evaluation for user: %w", err)
		}
		evaluations = append(evaluations, eval)
	}

	return evaluations, nil
}
