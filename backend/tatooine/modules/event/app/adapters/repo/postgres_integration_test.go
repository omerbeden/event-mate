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

	eventRepository := repo.NewEventRepo(pool)
	locationRepository := repo.NewLocationRepo(pool)

	defer eventRepository.Close()

	event := model.Event{
		Title:        "test title",
		Category:     "test category",
		CreatedBy:    model.User{ID: 1},
		Location:     model.Location{City: "Sakarya"},
		Participants: []model.User{{ID: 1}, {ID: 2}, {ID: 3}}}

	res, err := eventRepository.Create(event)
	assert.NoError(t, err)
	assert.NotNil(t, res)

	resLoc, err := locationRepository.Create(&res.Location)
	assert.NoError(t, err)
	assert.True(t, resLoc)

}

func TestAddParticipants(t *testing.T) {
	dbConfig := postgres.PostgresConfig{
		ConnectionString: "postgres://postgres:password@localhost:5432/test",
		Config:           pgxpool.Config{MinConns: 5, MaxConns: 10},
	}
	pool := postgres.NewConn(&dbConfig)
	eventRepository := repo.NewEventRepo(pool)
	event := model.Event{
		ID:           1,
		Title:        "test title",
		Category:     "test category",
		CreatedBy:    model.User{ID: 1},
		Location:     model.Location{City: "Sakarya"},
		Participants: []model.User{{ID: 1}, {ID: 2}, {ID: 3}}}

	err := eventRepository.AddParticipants(event)
	assert.NoError(t, err)

}

func TestAddParticipant(t *testing.T) {
	dbConfig := postgres.PostgresConfig{
		ConnectionString: "postgres://postgres:password@localhost:5432/test",
		Config:           pgxpool.Config{MinConns: 5, MaxConns: 10},
	}
	pool := postgres.NewConn(&dbConfig)
	eventRepository := repo.NewEventRepo(pool)
	event := model.Event{
		ID:           1,
		Title:        "test title",
		Category:     "test category",
		CreatedBy:    model.User{ID: 1},
		Location:     model.Location{City: "Sakarya"},
		Participants: []model.User{{ID: 1}, {ID: 2}, {ID: 3}}}

	user := model.User{ID: 4}

	err := eventRepository.AddParticipant(event.ID, user)
	assert.NoError(t, err)

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

func TestGetEventByLocation(t *testing.T) {
	dbConfig := postgres.PostgresConfig{
		ConnectionString: "postgres://postgres:password@localhost:5432/test",
		Config:           pgxpool.Config{MinConns: 5, MaxConns: 10},
	}
	pool := postgres.NewConn(&dbConfig)
	repository := repo.NewEventRepo(pool)
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
	repository := repo.NewEventRepo(pool)
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
	repository := repo.NewEventRepo(pool)
	defer repository.Close()

	res, err := repository.DeleteByID(1)
	assert.NotNil(t, res)
	assert.NoError(t, err)

}

func TestCreateLocation(t *testing.T) {
	dbConfig := postgres.PostgresConfig{
		ConnectionString: "postgres://postgres:password@localhost:5432/test",
		Config:           pgxpool.Config{MinConns: 5, MaxConns: 10},
	}
	pool := postgres.NewConn(&dbConfig)
	repository := repo.NewLocationRepo(pool)
	defer repository.Close()

	loc := model.Location{
		EventId: 1,
		City:    "Sakarya",
	}

	res, err := repository.Create(&loc)
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

	locationToBeUpdated := model.Location{
		EventId: 1,
		City:    "Istanbul",
	}
	res, err := repository.UpdateByID(locationToBeUpdated)
	assert.NotNil(t, res)
	assert.NoError(t, err)
}
