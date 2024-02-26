package repo

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type activityRulesRepository struct {
	pool *pgxpool.Pool
}

func NewActivityRulesRepo(pool *pgxpool.Pool) *activityRulesRepository {
	return &activityRulesRepository{
		pool: pool,
	}
}
func (r *activityRulesRepository) Close() {
	r.pool.Close()
}

func (r *activityRulesRepository) CreateActivityRules(ctx context.Context, acitivtyId int64, rules []string) error {

	var ruleRows [][]interface{}
	for _, rule := range rules {
		ruleRows = append(ruleRows, []interface{}{acitivtyId, rule})
	}

	copyCount, err := r.pool.CopyFrom(ctx,
		pgx.Identifier{"activity_rules"},
		[]string{"activity_id", "description"},
		pgx.CopyFromRows(ruleRows))

	if err != nil {
		return err
	}

	if int(copyCount) != len(rules) {
		return err
	}
	return nil
}

func (r *activityRulesRepository) GetActivityRules(ctx context.Context, activityId int64) ([]string, error) {

	q := `SELECT description FROM activity_rules WHERE activity_id = $1`
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
