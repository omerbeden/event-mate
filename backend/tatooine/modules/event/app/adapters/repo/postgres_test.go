package repo_test

import (
	"testing"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/omerbeden/event-mate/backend/tatooine/modules/event/app/adapters/repo"
	"github.com/omerbeden/event-mate/backend/tatooine/modules/event/app/domain/model"
	postgres "github.com/omerbeden/event-mate/backend/tatooine/pkg/database"
	"github.com/stretchr/testify/assert"
)

func TestCreateEvent(t *testing.T) {
	dbConfig := postgres.PostgresConfig{
		ConnectionString: "postgres://postgres:password@localhost:5432/test",
		Config:           pgxpool.Config{MinConns: 5, MaxConns: 10},
	}
	pool := postgres.NewConn(&dbConfig)

	eventRepository := repo.NewEventRepo(pool, repo.NewLocationRepo(pool))

	defer eventRepository.Close()

	res, err := eventRepository.Create(
		model.Event{
			Title:        "test title",
			Category:     "test category",
			CreatedBy:    model.User{ID: 1},
			Location:     model.Location{City: "Sakarya"},
			Participants: []model.User{{ID: 1}, {ID: 2}, {ID: 3}},
		})

	assert.NoError(t, err)
	assert.True(t, res)

}

func TestGetEventByID(t *testing.T) {
	dbConfig := postgres.PostgresConfig{
		ConnectionString: "postgres://postgres:password@localhost:5432/test",
		Config:           pgxpool.Config{MinConns: 5, MaxConns: 10},
	}
	pool := postgres.NewConn(&dbConfig)
	repository := repo.NewEventRepo(pool, repo.NewLocationRepo(pool))
	defer repository.Close()

	res, err := repository.GetByID(1)

	assert.NotNil(t, res)
	assert.NoError(t, err)
}

func TestGetEventByLocation(t *testing.T) {
	dbConfig := postgres.PostgresConfig{
		ConnectionString: "postgres://postgres:password@localhost:5432/test",
		Config:           pgxpool.Config{MinConns: 5, MaxConns: 10},
	}
	pool := postgres.NewConn(&dbConfig)
	repository := repo.NewEventRepo(pool, repo.NewLocationRepo(pool))
	defer repository.Close()

	res, err := repository.GetByLocation(&model.Location{City: "Sakarya"})

	assert.NotNil(t, res)
	assert.NoError(t, err)
}

func TestUpdateEvent(t *testing.T) {

	dbConfig := postgres.PostgresConfig{
		ConnectionString: "postgres://postgres:password@localhost:5432/test",
		Config:           pgxpool.Config{MinConns: 5, MaxConns: 10},
	}
	pool := postgres.NewConn(&dbConfig)
	repository := repo.NewEventRepo(pool, repo.NewLocationRepo(pool))
	defer repository.Close()

	eventTobeUpdated := model.Event{
		Title:     "Updated title",
		Category:  "Updated Category",
		CreatedBy: model.User{ID: 2},
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
	repository := repo.NewEventRepo(pool, repo.NewLocationRepo(pool))
	defer repository.Close()

	res, err := repository.DeleteByID(1)
	assert.NotNil(t, res)
	assert.NoError(t, err)

}

//TODO: Migrate datebase and run the test from docker
