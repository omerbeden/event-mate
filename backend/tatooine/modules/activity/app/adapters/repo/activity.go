package repo

import (
	"context"
	"fmt"
	"time"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/omerbeden/event-mate/backend/tatooine/modules/activity/app/domain/model"
)

const errlogprefix = "repo:activity"

type activityRepository struct {
	pool *pgxpool.Pool
}

func NewActivityRepo(pool *pgxpool.Pool) *activityRepository {
	return &activityRepository{
		pool: pool,
	}
}
func (r *activityRepository) Close() {
	r.pool.Close()
}

func (r *activityRepository) Create(activity model.Activity) (*model.Activity, error) {

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	var ID int64
	q := `INSERT INTO activities (title,category,created_by,background_image_url,start_at,content) 
	Values($1,$2,$3,$4,$5,$6) RETURNING ID`

	fmt.Printf("model: %+v \n", activity)
	err := r.pool.QueryRow(
		ctx,
		q,
		activity.Title,
		activity.Category,
		activity.CreatedBy.ID,
		activity.BackgroundImageUrl,
		activity.StartAt,
		activity.Content).Scan(&ID)

	if err != nil {

		return nil, fmt.Errorf("%s could not insert activity %w", errlogprefix, err)
	}

	activity.ID = ID
	activity.Location.ActivityId = ID

	return &activity, nil
}

func (r *activityRepository) AddParticipants(activity model.Activity) error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	var linkedParticipants [][]interface{}
	for _, parparticipants := range activity.Participants {
		linkedParticipants = append(linkedParticipants, []interface{}{activity.ID, parparticipants.ID})

	}

	copyCount, err := r.pool.CopyFrom(ctx,
		pgx.Identifier{"participants"},
		[]string{"activity_id", "user_id"},
		pgx.CopyFromRows(linkedParticipants),
	)

	if err != nil {
		return err
	}
	if int(copyCount) != len(activity.Participants) {
		return err
	}
	return nil
}

func (r *activityRepository) AddParticipant(activityId int64, user model.User) error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()
	q := `INSERT INTO participants(activity_id,user_id) VALUES($1,$2)`

	_, err := r.pool.Exec(ctx, q, activityId, user.ID)
	if err != nil {
		return fmt.Errorf("%s could not insert participant for activity %d , %w", errlogprefix, activityId, err)
	}

	return nil

}

func (r *activityRepository) GetParticipants(acitivityId int64) ([]model.User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()
	q := `SELECT u.id FROM user_profiles u
	RIGHT JOIN participants p ON p.user_id = u.id
	WHERE p.activity_id = $1
	`

	var participants []model.User
	rows, err := r.pool.Query(ctx, q, acitivityId)
	if err != nil {
		return nil, fmt.Errorf("%s could not get participants , acitivityId: %d  %w", errlogprefix, acitivityId, err)
	}
	for rows.Next() {
		var res model.User
		err := rows.Scan(&res.ID)
		if err != nil {
			return nil, fmt.Errorf("err getting rows %w ", err)
		}
		participants = append(participants, res)

	}

	return participants, nil

}

func (r *activityRepository) GetByID(id int64) (*model.Activity, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	q := `SELECT e.id, title, category, e.created_by, l.city
	FROM activities e
	LEFT JOIN user_profiles u ON e.created_by = u.id
	LEFT JOIN activity_locations l ON e.id = l.activity_id
	Where e.id = $1	
	`
	var activity model.Activity
	err := r.pool.QueryRow(ctx, q, id).Scan(&activity.ID, &activity.Title, &activity.Category, &activity.CreatedBy.ID, &activity.Location.City)
	if err != nil {
		return nil, fmt.Errorf("%s could not get activity by id: %d %w", errlogprefix, id, err)
	}

	return &activity, nil
}
func (r *activityRepository) GetByLocation(loc *model.Location) ([]model.Activity, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	q := `SELECT e.id, title, category, e.created_by, l.city
	FROM activities e
	LEFT JOIN user_profiles u ON e.created_by = u.id
	LEFT JOIN activity_locations l ON e.id = l.activity_id
	Where l.city= $1`

	var activity []model.Activity
	rows, err := r.pool.Query(ctx, q, loc.City)
	if err != nil {
		return nil, fmt.Errorf("%s could not get activity by loc: id: %s  %w", errlogprefix, loc.City, err)
	}

	for rows.Next() {
		var res model.Activity
		err := rows.Scan(&res.ID, &res.Title, &res.Category, &res.CreatedBy.ID, &res.Location.City)
		if err != nil {
			return nil, fmt.Errorf("err getting rows %w ", err)
		}
		activity = append(activity, res)

	}

	return activity, nil
}

func (r *activityRepository) UpdateByID(id int32, activity model.Activity) (bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	q := `UPDATE activities
	 SET title  = $1,
	  category = $2,
	  created_by = $3
	 WHERE id = $4
	 `
	_, err := r.pool.Exec(ctx, q, activity.Title, activity.Category, activity.CreatedBy.ID, id)
	if err != nil {
		return false, fmt.Errorf("%s could not update activity id: %d %w", errlogprefix, id, err)
	}

	return true, nil
}

func (r *activityRepository) DeleteByID(id int32) (bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	q := `DELETE FROM activities  WHERE id = $1`
	_, err := r.pool.Exec(ctx, q, id)
	if err != nil {
		return false, fmt.Errorf("%s could not delete activity id: %d %w", errlogprefix, id, err)
	}

	return true, nil
}
