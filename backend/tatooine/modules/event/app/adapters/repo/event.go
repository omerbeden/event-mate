package repo

import (
	"context"
	"fmt"
	"time"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/omerbeden/event-mate/backend/tatooine/modules/event/app/domain/model"
)

const errlogprefix = "repo:event"

type eventRepository struct {
	pool *pgxpool.Pool
}

func NewEventRepo(pool *pgxpool.Pool) *eventRepository {
	return &eventRepository{
		pool: pool,
	}
}
func (r *eventRepository) Close() {
	r.pool.Close()
}

func (r *eventRepository) Create(event model.Event) (*model.Event, error) {

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	var ID int64
	q := `INSERT INTO events (title,category,created_user_id) Values($1,$2,$3) RETURNING ID`
	errInsertingEvent := r.pool.QueryRow(ctx, q, event.Title, event.Category, event.CreatedBy.ID).Scan(&ID)

	if errInsertingEvent != nil {

		return nil, fmt.Errorf("%s could not insert event %w", errlogprefix, errInsertingEvent)
	}

	event.ID = ID
	event.Location.EventId = ID

	return &event, nil
}

func (r *eventRepository) AddParticipants(event model.Event) error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	var linkedParticipants [][]interface{}
	for _, parparticipants := range event.Participants {
		linkedParticipants = append(linkedParticipants, []interface{}{event.ID, parparticipants.ID})

	}

	copyCount, err := r.pool.CopyFrom(ctx,
		pgx.Identifier{"participants"},
		[]string{"event_id", "user_id"},
		pgx.CopyFromRows(linkedParticipants),
	)

	if err != nil {
		return err
	}
	if int(copyCount) != len(event.Participants) {
		return err
	}
	return nil
}

func (r *eventRepository) AddParticipant(eventId int64, user model.User) error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()
	q := `INSERT INTO participants(event_id,user_id) VALUES($1,$2)`

	_, err := r.pool.Exec(ctx, q, eventId, user.ID)
	if err != nil {
		return fmt.Errorf("%s could not insert participant for event %d , %w", errlogprefix, eventId, err)
	}

	return nil

}

func (r *eventRepository) GetByID(id int32) (*model.Event, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	q := `SELECT e.id, title, category, e.created_user_id, l.city
	FROM events e
	LEFT JOIN user_profiles u ON e.created_user_id = u.id
	LEFT JOIN event_locations l ON e.id = l.event_id
	Where e.id = $1	
	`
	var event model.Event
	err := r.pool.QueryRow(ctx, q, id).Scan(&event.ID, &event.Title, &event.Category, &event.CreatedBy.ID, &event.Location.City)
	if err != nil {
		return nil, fmt.Errorf("%s could not get event by id: %d %w", errlogprefix, id, err)
	}

	return &event, nil
}
func (r *eventRepository) GetByLocation(loc *model.Location) ([]model.Event, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	q := `SELECT e.id, title, category, e.created_user_id, l.city
	FROM events e
	LEFT JOIN user_profiles u ON e.created_user_id = u.id
	LEFT JOIN event_locations l ON e.id = l.event_id
	Where l.city= $1`

	var events []model.Event
	rows, err := r.pool.Query(ctx, q, loc.City)
	if err != nil {
		return nil, fmt.Errorf("%s could not get event by loc: id: %s  %w", errlogprefix, loc.City, err)
	}

	for rows.Next() {
		var res model.Event
		err := rows.Scan(&res.ID, &res.Title, &res.Category, &res.CreatedBy.ID, &res.Location.City)
		if err != nil {
			return nil, fmt.Errorf("err getting rows %w ", err)
		}
		events = append(events, res)

	}

	return events, nil
}

func (r *eventRepository) UpdateByID(id int32, event model.Event) (bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	q := `UPDATE events
	 SET title  = $1,
	  category = $2,
	  created_user_id = $3
	 WHERE id = $4
	 `
	_, err := r.pool.Exec(ctx, q, event.Title, event.Category, event.CreatedBy.ID, id)
	if err != nil {
		return false, fmt.Errorf("%s could not update event id: %d %w", errlogprefix, id, err)
	}

	return true, nil
}

func (r *eventRepository) DeleteByID(id int32) (bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	q := `DELETE FROM events  WHERE id = $1`
	_, err := r.pool.Exec(ctx, q, id)
	if err != nil {
		return false, fmt.Errorf("%s could not delete event id: %d %w", errlogprefix, id, err)
	}

	return true, nil
}
