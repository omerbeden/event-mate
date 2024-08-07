package postgresadapter

import (
	"context"
	"fmt"

	"github.com/omerbeden/event-mate/backend/tatooine/modules/activity/app/domain/model"
	"github.com/omerbeden/event-mate/backend/tatooine/pkg/db"
)

const errLogPrefix = "repo:location"

type LocationRepo struct {
	pool db.Executor
}

func NewLocationRepo(pool db.Executor) *LocationRepo {
	return &LocationRepo{
		pool: pool,
	}
}

func (r *LocationRepo) Create(ctx context.Context, tx db.Tx, location *model.Location) (bool, error) {

	ql := `INSERT INTO activity_locations(activity_id,city,district,description,latitude,longitude) Values($1,$2,$3,$4,$5,$6)`
	_, err := tx.Exec(ctx, ql, location.ActivityId, location.City, location.District, location.Description, location.Latitude, location.Longitude)
	if err != nil {
		return false, fmt.Errorf("%s could not create location for activity %d  %w", errLogPrefix, location.ActivityId, err)
	}

	return true, nil
}

func (r *LocationRepo) UpdateByID(ctx context.Context, loc model.Location) (bool, error) {

	q := `UPDATE activity_locations
			SET city = $1
			Where activity_id = $2`
	_, err := r.pool.Exec(ctx, q, loc.City, loc.ActivityId)
	if err != nil {
		return false, fmt.Errorf("%s could not update activity id: %d %w", errLogPrefix, loc.ActivityId, err)
	}

	return true, nil
}
