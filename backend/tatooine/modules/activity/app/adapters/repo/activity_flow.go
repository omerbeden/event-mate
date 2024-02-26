package repo

import (
	"context"
	"fmt"
	"time"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type activityFLowRepository struct {
	pool *pgxpool.Pool
}

func NewActivityFlowRepo(pool *pgxpool.Pool) *activityFLowRepository {
	return &activityFLowRepository{
		pool: pool,
	}
}
func (r *activityFLowRepository) Close() {
	r.pool.Close()
}

func (r *activityFLowRepository) CreateActivityFlow(acitivtyId int64, flows []string) error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	var flowRows [][]interface{}
	for _, flow := range flows {
		flowRows = append(flowRows, []interface{}{acitivtyId, flow})
	}

	copyCount, err := r.pool.CopyFrom(ctx,
		pgx.Identifier{"activity_flows"},
		[]string{"activity_id", "description"},
		pgx.CopyFromRows(flowRows))

	if err != nil {
		return err
	}

	if int(copyCount) != len(flows) {
		return err
	}
	return nil
}

func (r *activityFLowRepository) GetActivityFlow(activityId int64) ([]string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	q := `SELECT description FROM activity_flows WHERE activity_id = $1`
	rows, err := r.pool.Query(ctx, q, activityId)
	if err != nil {
		return nil, fmt.Errorf("could not get rules for activity: %d %w", activityId, err)
	}

	var activityRules []string
	for rows.Next() {
		var rule string
		err := rows.Scan(&rule)
		if err != nil {
			return nil, fmt.Errorf("err getting rows %w ", err)
		}
		activityRules = append(activityRules, rule)
	}

	return activityRules, nil

}
