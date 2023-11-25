package repo

import (
	"context"
	"fmt"
	"time"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/omerbeden/event-mate/backend/tatooine/modules/event/app/domain/model"
	"github.com/omerbeden/event-mate/backend/tatooine/modules/event/app/domain/ports/repositories"
)

type EventRepository struct {
	pool    *pgxpool.Pool
	locRepo repositories.LocationRepository
}

func NewEventRepo(pool *pgxpool.Pool) *EventRepository {
	return &EventRepository{
		pool: pool,
	}
}
func (r *EventRepository) Close() {
	r.pool.Close()
}

func (r *EventRepository) Create(event model.Event) (bool, error) {

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	tx, err := r.pool.Begin(ctx)
	if err != nil {
		return false, err
	}
	defer func() {
		if err != nil {
			tx.Rollback(ctx)
		} else {
			tx.Commit(ctx)
		}
	}()

	q := `INSERT INTO events (title,category,createdbyuserid,locationid) Values($1,$2,$3,$4) `
	_, err1 := tx.Exec(ctx, q, event.Title, event.Category, event.CreatedBy.ID, event.Location.EventId)

	if err1 != nil {
		return false, fmt.Errorf("could not create event %w", err1)
	}
	err2 := r.AddParticipants(event)
	if err2 != nil {
		tx.Rollback(ctx)
		return false, err2
	}

	_, err3 := r.locRepo.Create(event.Location)
	if err3 != nil {
		tx.Rollback(ctx)
		return false, err3
	}

	return true, nil
}

func (r *EventRepository) AddParticipants(event model.Event) error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	copyCount, err := r.pool.CopyFrom(ctx,
		pgx.Identifier{"participants"},
		[]string{"event_id", "user_id"},
		pgx.CopyFromSlice(len(event.Participants), func(i int) ([]any, error) {
			return []any{event.ID, event.Participants[i]}, nil
		}),
	)
	if err != nil {
		return err
	}
	if int(copyCount) != len(event.Participants) {
		return err
	}
	return nil
}

func (r *EventRepository) GetByID(id int32) (*model.Event, error) {
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
func (r *EventRepository) GetByLocation(loc *model.Location) ([]model.Event, error) {
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

func (r *EventRepository) UpdateByID(id int32, event model.Event) (bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	q := `UPDATE`
	_, err := r.pool.Exec(ctx, q)
	if err != nil {
		return false, fmt.Errorf("could not update event id: %d %w", id, err)
	}

	return true, nil
}

func (r *EventRepository) DeleteByID(id int32) (bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	q := `Delete`
	_, err := r.pool.Exec(ctx, q)
	if err != nil {
		return false, fmt.Errorf("could not delete event id: %d %w", id, err)
	}

	return true, nil
}
