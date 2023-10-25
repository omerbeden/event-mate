package repo

import (
	"context"
	"fmt"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/omerbeden/event-mate/backend/tatooine/modules/event/app/domain/model"
)

type LocationRepository struct {
	pool *pgxpool.Pool
}

func NewLocationRepo(pool *pgxpool.Pool) *LocationRepository {
	return &LocationRepository{
		pool: pool,
	}
}

func (r *LocationRepository) Close() {
	r.pool.Close()
}

func (r *LocationRepository) Create(location model.Location) (bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	q := `INSERT INTO locations (city) Values($1) `
	_, err := r.pool.Exec(ctx, q, location.City)
	if err != nil {
		return false, fmt.Errorf("could not create %w", err)
	}

	return true, nil
}

func (r *LocationRepository) DeleteByID(id int32) (bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Millisecond*5)
	defer cancel()

	q := ``
	_, err := r.pool.Exec(ctx, q)
	if err != nil {
		return false, fmt.Errorf("could not delete %w", err)
	}
	return true, nil
}

func (r *LocationRepository) UpdateByID(id int32, loc model.Location) (bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	q := `UPDATE`
	_, err := r.pool.Exec(ctx, q)
	if err != nil {
		return false, fmt.Errorf("could not update event id: %d %w", id, err)
	}

	return true, nil
}
