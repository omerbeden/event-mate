package repo_test

import (
	"testing"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/omerbeden/event-mate/backend/tatooine/modules/event/app/adapters/repo"
	"github.com/omerbeden/event-mate/backend/tatooine/modules/event/app/domain/model"
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
	res, err := repository.Create(
		model.Location{City: "Istanbul"})

	assert.NoError(t, err)
	assert.True(t, res)

}

func TestCreateEvent(t *testing.T) {

	dbConfig := postgres.PostgresConfig{
		ConnectionString: "postgres://postgres:password@localhost:5432/test",
		Config:           pgxpool.Config{MinConns: 5, MaxConns: 10},
	}
	pool := postgres.NewConn(&dbConfig)
	repository := repo.NewEventRepo(pool)
	defer repository.Close()
	res, err := repository.Create(
		model.Event{ID: 1,
			Title:     "test title",
			Category:  "test category",
			CreatedBy: model.User{ID: 1},
			Location:  model.Location{City: "Sakarya"}})

	assert.NoError(t, err)
	assert.True(t, res)

}

func TestGetEventByID(t *testing.T) {
	dbConfig := postgres.PostgresConfig{
		ConnectionString: "postgres://postgres:password@localhost:5432/test",
		Config:           pgxpool.Config{MinConns: 5, MaxConns: 10},
	}
	pool := postgres.NewConn(&dbConfig)
	repository := repo.NewEventRepo(pool)
	defer repository.Close()

	res, err := repository.GetByID(1)

	assert.NotNil(t, res)
	assert.NoError(t, err)
}

func TestUpdateEvent(t *testing.T) {

	dbConfig := postgres.PostgresConfig{
		ConnectionString: "postgres://postgres:password@localhost:5432/test",
		Config:           pgxpool.Config{MinConns: 5, MaxConns: 10},
	}
	pool := postgres.NewConn(&dbConfig)
	repository := repo.NewEventRepo(pool)
	defer repository.Close()

	eventTobeUpdated := model.Event{
		Title:    "Updated title",
		Category: "Updated Category",
	}

	res, err := repository.UpdateByID(1, eventTobeUpdated)
	assert.NotNil(t, res)
	assert.NoError(t, err)

}

func TestDeleteEventByID(t *testing.T) {

	dbConfig := postgres.PostgresConfig{
		ConnectionString: "postgres://postgres:password@localhost:5432/test",
		Config:           pgxpool.Config{MinConns: 5, MaxConns: 10},
	}
	pool := postgres.NewConn(&dbConfig)
	repository := repo.NewEventRepo(pool)
	defer repository.Close()

	res, err := repository.DeleteByID(1)
	assert.NotNil(t, res)
	assert.NoError(t, err)

}

//TODO: Migrate datebase and run the test from docker
