package repo

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/omerbeden/event-mate/backend/tatooine/modules/event/app/domain/model"
)

type Repository struct {
	pool *pgxpool.Pool
}

func New(cnnStr string) *Repository {
	//dbUrl := os.Getenv("Db_Conn_Str")
	config, err := pgxpool.ParseConfig(cnnStr)
	if err != nil {
		fmt.Fprintf(os.Stderr, "unable to parse config: %v\n", err)
		os.Exit(1)
	}
	config.MinConns = 5
	config.MaxConns = 10
	//later import db tracer

	pool, err := pgxpool.NewWithConfig(context.Background(), config)
	if err != nil {
		fmt.Fprintf(os.Stderr, "unable to parse config: %v\n", err)
		os.Exit(1)
	}
	return &Repository{
		pool: pool,
	}
}
func (r *Repository) Close() {
	r.pool.Close()
}

func (r *Repository) Create(event model.Event) (bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	q := `INSERT INTO events (title,category,createdbyuserid,locationid) Values($1,$2,$3,$4) `
	_, err := r.pool.Exec(ctx, q, event.Title, event.Category, event.CreatedBy.ID, event.Location.ID)
	if err != nil {
		return false, fmt.Errorf("could not create %w", err)
	}

	return true, nil
}
func (r *Repository) GetByID(id int32) (*model.Event, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	q := `SELECT e.id, e.title, e.category, e.created_by, e.location_city
	FROM events e
	LEFT JOIN users u ON e.created_by = u.user_id
	LEFT JOIN locations l ON e.locationId = locations.id
	Where u.user_id = $1	
`
	var event model.Event
	err := r.pool.QueryRow(ctx, q, id).Scan(&event.Category, &event.CreatedBy, &event.Location, &event.Title)
	if err != nil {
		return nil, fmt.Errorf("could not get event by id: %d %w", id, err)
	}

	return &event, nil
}
func (r *Repository) GetByLocation(loc *model.Location) ([]model.Event, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	q := `SELECT e.id, e.title, e.category, e.created_by, e.location_city
	FROM events e
	LEFT JOIN users u ON e.created_by = u.user_id
	LEFT JOIN locations l ON e.locationId = locations.id
	Where l.city= $1`
	var events []model.Event
	rows, err := r.pool.Query(ctx, q, loc.City)
	if err != nil {
		return nil, fmt.Errorf("could not get event by loc: id: %s  %w", loc.City, err)
	}

	for rows.Next() {
		var res model.Event
		err := rows.Scan(&res.Category, &res.CreatedBy, &res.Location, &res.Title, &res)
		if err != nil {
			return nil, fmt.Errorf("err getting rows %w ", err)
		}
		events = append(events, res)

	}

	return events, nil
}

func (r *Repository) UpdateByID(id int32, event model.Event) (bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	q := `UPDATE`
	_, err := r.pool.Exec(ctx, q)
	if err != nil {
		return false, fmt.Errorf("could not update event id: %d %w", id, err)
	}

	return true, nil
}

func (r *Repository) DeleteByID(id int32) (bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	q := `Delete`
	_, err := r.pool.Exec(ctx, q)
	if err != nil {
		return false, fmt.Errorf("could not delete event id: %d %w", id, err)
	}

	return true, nil
}
