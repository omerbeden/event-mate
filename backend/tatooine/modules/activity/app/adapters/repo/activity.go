package repo

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/omerbeden/event-mate/backend/tatooine/modules/activity/app/domain/model"
)

const errlogprefix = "repo:activity"

type activityRepository struct {
	pool DBExecutor
}

type DBExecutor interface {
	Begin(ctx context.Context) (pgx.Tx, error)
	BeginTx(ctx context.Context, txOptions pgx.TxOptions) (pgx.Tx, error)
	Close()
	Config() *pgxpool.Config
	CopyFrom(ctx context.Context, tableName pgx.Identifier, columnNames []string, rowSrc pgx.CopyFromSource) (int64, error)
	Exec(ctx context.Context, sql string, arguments ...any) (pgconn.CommandTag, error)
	Ping(ctx context.Context) error
	Query(ctx context.Context, sql string, args ...any) (pgx.Rows, error)
	QueryRow(ctx context.Context, sql string, args ...any) pgx.Row
}

var _ DBExecutor = (*pgxpool.Pool)(nil)

func NewActivityRepo(pool DBExecutor) *activityRepository {
	return &activityRepository{
		pool: pool,
	}
}

func (r *activityRepository) Close() {
	r.pool.Close()
}

func (r *activityRepository) Create(ctx context.Context, activity model.Activity) (*model.Activity, error) {

	var ID int64
	q := `INSERT INTO activities (title,category,created_by,start_at,end_at,content,quota) 
	Values($1,$2,$3,$4,$5,$6,$7) RETURNING id`

	err := r.pool.QueryRow(
		ctx,
		q,
		activity.Title,
		activity.Category,
		activity.CreatedBy.ID,
		activity.StartAt,
		activity.EndAt,
		activity.Content,
		activity.Quota).Scan(&ID)
	if err != nil {
		return nil, fmt.Errorf("%s could not insert activity %w", errlogprefix, err)
	}

	activity.ID = ID
	activity.Location.ActivityId = ID

	return &activity, nil
}

func (r *activityRepository) AddParticipants(ctx context.Context, activityId int64, participants []model.User) error {

	var linkedParticipants [][]interface{}
	for _, parparticipants := range participants {
		linkedParticipants = append(linkedParticipants, []interface{}{activityId, parparticipants.ID})

	}

	copyCount, err := r.pool.CopyFrom(ctx,
		pgx.Identifier{"participants"},
		[]string{"activity_id", "user_id"},
		pgx.CopyFromRows(linkedParticipants),
	)

	if err != nil {
		return err
	}
	if int(copyCount) != len(participants) {
		return err
	}
	return nil
}

func (r *activityRepository) AddParticipant(ctx context.Context, activityId int64, user model.User) error {

	q := `INSERT INTO participants(activity_id,user_id) VALUES($1,$2)`

	_, err := r.pool.Exec(ctx, q, activityId, user.ID)
	if err != nil {
		return fmt.Errorf("%s could not insert participant for activity %d , %w", errlogprefix, activityId, err)
	}

	return nil

}

func (r *activityRepository) GetParticipants(ctx context.Context, acitivityId int64) ([]model.User, error) {

	q := `SELECT u.id, u.name, u.last_name, 
	COALESCE(stats.point, 0.0) AS point
	FROM user_profiles u
	RIGHT JOIN participants p ON p.user_id = u.id
	LEFT JOIN user_profile_stats stats ON stats.profile_id = u.id
	WHERE p.activity_id = $1
	`

	var participants []model.User
	rows, err := r.pool.Query(ctx, q, acitivityId)
	if err != nil {
		return nil, fmt.Errorf("%s could not get participants , acitivityId: %d  %w", errlogprefix, acitivityId, err)
	}
	for rows.Next() {
		var res model.User
		err := rows.Scan(&res.ID, &res.Name, &res.LastName, &res.ProfilePoint)
		if err != nil {
			return nil, fmt.Errorf("err getting rows %w ", err)
		}
		participants = append(participants, res)

	}

	return participants, nil

}

func (r *activityRepository) GetByID(ctx context.Context, id int64) (*model.Activity, error) {

	q := `SELECT a.id, a.title, a.category, a.created_by,a.start_at,a.end_at,a.content,a.quota,
	 l.city
	FROM activities a
	LEFT JOIN user_profiles u ON a.created_by = u.id
	LEFT JOIN activity_locations l ON a.id = l.activity_id
	Where a.id = $1	
	`
	var activity model.Activity
	err := r.pool.QueryRow(ctx, q, id).Scan(&activity.ID, &activity.Title, &activity.Category, &activity.CreatedBy.ID,
		&activity.StartAt, &activity.EndAt, &activity.Content, &activity.Quota, &activity.Location.City)
	if err != nil {
		return nil, fmt.Errorf("%s could not get activity by id: %d %w", errlogprefix, id, err)
	}

	return &activity, nil
}
func (r *activityRepository) GetByLocation(ctx context.Context, loc *model.Location) ([]model.Activity, error) {

	q := `SELECT a.id, a.title, a.category,a.start_at, a.end_at,a.content,a.quota,
	u.id, u.name, u.last_name, u.profile_image_url, 
	COALESCE(stats.point, 0.0)  as point
	, l.city
	FROM activities a
	LEFT JOIN user_profiles u ON a.created_by = u.id
	LEFT JOIN user_profile_stats stats ON stats.profile_id = u.id
	LEFT JOIN activity_locations l ON a.id = l.activity_id
	Where l.city= $1`

	var activities []model.Activity
	rows, err := r.pool.Query(ctx, q, loc.City)
	if err != nil {
		return nil, fmt.Errorf("%s could not get activity by loc: id: %s  %w", errlogprefix, loc.City, err)
	}

	for rows.Next() {
		var res model.Activity
		err := rows.Scan(&res.ID, &res.Title, &res.Category, &res.StartAt, &res.EndAt, &res.Content, &res.Quota,
			&res.CreatedBy.ID, &res.CreatedBy.Name, &res.CreatedBy.LastName, &res.CreatedBy.ProfileImageUrl, &res.CreatedBy.ProfilePoint,
			&res.Location.City)
		if err != nil {
			return nil, fmt.Errorf("err getting rows %w ", err)
		}
		activities = append(activities, res)
	}

	return activities, nil
}

func (r *activityRepository) UpdateByID(ctx context.Context, id int64, activity model.Activity) (bool, error) {

	q := `UPDATE activities
	 SET title  = $1,
	  category = $2,
	  created_by = $3,
	  quota = $4
	 WHERE id = $5
	 `
	_, err := r.pool.Exec(ctx, q, activity.Title, activity.Category, activity.CreatedBy.ID, activity.Quota, id)
	if err != nil {
		return false, fmt.Errorf("%s could not update activity id: %d %w", errlogprefix, id, err)
	}

	return true, nil
}

func (r *activityRepository) DeleteByID(ctx context.Context, id int64) (bool, error) {

	q := `DELETE FROM activities  WHERE id = $1`
	_, err := r.pool.Exec(ctx, q, id)
	if err != nil {
		return false, fmt.Errorf("%s could not delete activity id: %d %w", errlogprefix, id, err)
	}

	return true, nil
}
