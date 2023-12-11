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

	ql := `INSERT INTO event_locations(event_id,City) Values($1,$2)`
	_, err := r.pool.Exec(ctx, ql, location.EventId, location.City)
	if err != nil {
		return false, fmt.Errorf("%s could not create location for event %d  %w", errLogPrefix, location.EventId, err)
	}

	return true, nil
}

func (r *LocationRepo) UpdateByID(loc model.Location) (bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	q := `UPDATE event_locations
			SET city = $1
			Where event_id = $2`
	_, err := r.pool.Exec(ctx, q, loc.City, loc.EventId)
	if err != nil {
		return false, fmt.Errorf("%s could not update event id: %d %w", errLogPrefix, loc.EventId, err)
	}

	return true, nil
}
