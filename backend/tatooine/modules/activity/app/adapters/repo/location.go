package repo

import (
	"context"
	"fmt"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/omerbeden/event-mate/backend/tatooine/modules/activity/app/domain/model"
)

const errLogPrefix = "repo:location"

type LocationRepo struct {
	pool *pgxpool.Pool
}

func NewLocationRepo(pool *pgxpool.Pool) *LocationRepo {
	return &LocationRepo{
		pool: pool,
	}
}

func (r *LocationRepo) Close() {
	r.pool.Close()
}

func (r *LocationRepo) Create(location *model.Location) (bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	ql := `INSERT INTO activity_locations(activity_id,city,district,description) Values($1,$2,$3,$4)`
	_, err := r.pool.Exec(ctx, ql, location.ActivityId, location.City, location.District, location.Description)
	if err != nil {
		return false, fmt.Errorf("%s could not create location for activity %d  %w", errLogPrefix, location.ActivityId, err)
	}

	return true, nil
}

func (r *LocationRepo) UpdateByID(loc model.Location) (bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	q := `UPDATE activity_locations
			SET city = $1
			Where activity_id = $2`
	_, err := r.pool.Exec(ctx, q, loc.City, loc.ActivityId)
	if err != nil {
		return false, fmt.Errorf("%s could not update activity id: %d %w", errLogPrefix, loc.ActivityId, err)
	}

	return true, nil
}
