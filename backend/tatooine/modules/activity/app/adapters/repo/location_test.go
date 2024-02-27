package repo_test

import (
	"context"
	"testing"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/omerbeden/event-mate/backend/tatooine/modules/activity/app/adapters/repo"
	"github.com/omerbeden/event-mate/backend/tatooine/modules/activity/app/domain/model"
	postgres "github.com/omerbeden/event-mate/backend/tatooine/pkg/database"
	"github.com/stretchr/testify/assert"
)

func TestCreateLocation(t *testing.T) {
	dbConfig := postgres.PostgresConfig{
		ConnectionString: "postgres://postgres:password@localhost:5432/test",
		Config:           pgxpool.Config{MinConns: 5, MaxConns: 10},
	}
	pool := postgres.NewConn(&dbConfig)
	repository := repo.NewLocationRepo(pool)
	defer repository.Close()
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	loc := model.Location{
		ActivityId: 1,
		City:       "Sakarya",
	}

	res, err := repository.Create(ctx, &loc)
	assert.NotNil(t, res)
	assert.NoError(t, err)
}
func TestUpdateLocation(t *testing.T) {
	dbConfig := postgres.PostgresConfig{
		ConnectionString: "postgres://postgres:password@localhost:5432/test",
		Config:           pgxpool.Config{MinConns: 5, MaxConns: 10},
	}
	pool := postgres.NewConn(&dbConfig)

	repository := repo.NewLocationRepo(pool)
	defer repository.Close()
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	locationToBeUpdated := model.Location{
		ActivityId: 1,
		City:       "Istanbul",
	}
	res, err := repository.UpdateByID(ctx, locationToBeUpdated)
	assert.NotNil(t, res)
	assert.NoError(t, err)
}
