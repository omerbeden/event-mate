package repo

import (
	"context"
	"fmt"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/omerbeden/event-mate/backend/tatooine/modules/event/app/domain/model"
)

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

func (r *LocationRepo) Create(location model.Location) (bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	ql := `INSERT INTO event_locations(event_id,City) Values($1,$2)`
	_, err := r.pool.Exec(ctx, ql, location.EventId, location.City)
	if err != nil {
		return false, fmt.Errorf("could not create location for event %d  %w", location.EventId, err)
	}

	return true, nil
}

func (r *LocationRepo) DeleteByID(id int32) (bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Millisecond*5)
	defer cancel()

	q := ``
	_, err := r.pool.Exec(ctx, q)
	if err != nil {
		return false, fmt.Errorf("could not delete %w", err)
	}
	return true, nil
}

func (r *LocationRepo) UpdateByID(id int32, loc model.Location) (bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	q := `UPDATE`
	_, err := r.pool.Exec(ctx, q)
	if err != nil {
		return false, fmt.Errorf("could not update event id: %d %w", id, err)
	}

	return true, nil
}
