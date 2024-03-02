//go:build integration
// +build integration

package integration_test

import (
	"context"
	"os"
	"testing"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/omerbeden/event-mate/backend/tatooine/modules/activity/app/adapters/repo"
	"github.com/omerbeden/event-mate/backend/tatooine/modules/activity/app/domain/model"
	"github.com/omerbeden/event-mate/backend/tatooine/pkg/db/postgres"
	"github.com/stretchr/testify/assert"
)

var pool *pgxpool.Pool

func TestMain(m *testing.M) {
	dbConfig := postgres.PostgresConfig{
		ConnectionString: "postgres://postgres:password@localhost:5432/test",
		Config:           pgxpool.Config{MinConns: 5, MaxConns: 10},
	}

	pool = postgres.NewConn(&dbConfig)

	code := m.Run()

	os.Exit(code)
}

func TestActivityRepository_Create(t *testing.T) {
	activityRepo := repo.NewActivityRepo(pool)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	activity := model.Activity{
		ID:        int64(1),
		Title:     "test",
		Category:  "test",
		CreatedBy: model.User{ID: 1},
		StartAt:   time.Now(),
		EndAt:     time.Now(),
		Content:   "test",
		Quota:     1,
		Location:  model.Location{ActivityId: int64(1), City: "London"},
	}

	result, err := activityRepo.Create(ctx, activity)

	assert.NotNil(t, result)
	assert.NoError(t, err)
	assert.Equal(t, &activity, result)
}

func TestActivityRepository_AddParticipants(t *testing.T) {
	activityRepo := repo.NewActivityRepo(pool)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	activityId := int64(1)
	participants := []model.User{
		{ID: 6},
		{ID: 7},
	}
	err := activityRepo.AddParticipants(ctx, activityId, participants)

	assert.NoError(t, err)
}
func TestActivityRepository_AddParticipant(t *testing.T) {
	activityRepo := repo.NewActivityRepo(pool)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	activityId := int64(1)
	participant := model.User{ID: 8}

	err := activityRepo.AddParticipant(ctx, activityId, participant)

	assert.NoError(t, err)

}
func TestActivityRepository_GetParticipants(t *testing.T) {
	activityRepo := repo.NewActivityRepo(pool)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	activityId := int64(1)
	participants, err := activityRepo.GetParticipants(ctx, activityId)

	assert.NotNil(t, participants)
	assert.NotEmpty(t, participants)
	assert.NoError(t, err)

}
func TestActivityRepository_GetByID(t *testing.T) {
	activityRepo := repo.NewActivityRepo(pool)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()
	activityId := int64(1)

	activity, err := activityRepo.GetByID(ctx, activityId)

	assert.NotNil(t, activity)
	assert.NoError(t, err)

}
func TestActivityRepository_GetByLocation(t *testing.T) {
	activityRepo := repo.NewActivityRepo(pool)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	location := model.Location{City: "London"}

	activities, err := activityRepo.GetByLocation(ctx, &location)

	assert.NotNil(t, activities)
	assert.NotEmpty(t, activities)
	assert.NoError(t, err)

}
func TestActivityRepository_UpdateByID(t *testing.T) {
	activityRepo := repo.NewActivityRepo(pool)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	activityId := int64(1)
	activity := model.Activity{Category: "sport", Title: "changed title", CreatedBy: model.User{ID: 8}, Quota: 3}

	res, err := activityRepo.UpdateByID(ctx, activityId, activity)

	assert.NoError(t, err)
	assert.True(t, res)
}
func TestActivityRepository_DeleteByID(t *testing.T) {
	activityRepo := repo.NewActivityRepo(pool)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	activityId := int64(1)
	res, err := activityRepo.DeleteByID(ctx, activityId)

	assert.NoError(t, err)
	assert.True(t, res)
}
