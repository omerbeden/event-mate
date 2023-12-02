package entrypoints_test

import (
	"context"
	"testing"

	"github.com/go-redis/redis/v8"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/omerbeden/event-mate/backend/tatooine/modules/event/app/adapters/repo"
	"github.com/omerbeden/event-mate/backend/tatooine/modules/event/app/domain/model"
	"github.com/omerbeden/event-mate/backend/tatooine/modules/event/app/entrypoints"
	postgres "github.com/omerbeden/event-mate/backend/tatooine/pkg/database"
	"github.com/stretchr/testify/assert"
)

func TestGetEventFromDBWhenRedisDown(t *testing.T) {
	dbConfig := postgres.PostgresConfig{
		ConnectionString: "postgres://postgres:password@localhost:5432/test",
		Config:           pgxpool.Config{MinConns: 5, MaxConns: 10},
	}
	pool := postgres.NewConn(&dbConfig)
	eventService := entrypoints.EventService{

		EventRepository:   repo.NewEventRepo(pool),
		LocationReposiroy: repo.NewLocationRepo(pool),
		RedisClient: *redis.NewClient(&redis.Options{
			Addr:     "Localhost:6379",
			Password: "",
			DB:       0,
		}),
	}

	res, err := eventService.GetEventById(context.Background(), 2)

	assert.NoError(t, err)
	assert.NotNil(t, res)
}

func TestGetEventByIDReturnErrorWhenEventIdNotFound(t *testing.T) {
	dbConfig := postgres.PostgresConfig{
		ConnectionString: "postgres://postgres:password@localhost:5432/test",
		Config:           pgxpool.Config{MinConns: 5, MaxConns: 10},
	}
	pool := postgres.NewConn(&dbConfig)
	eventService := entrypoints.EventService{

		EventRepository:   repo.NewEventRepo(pool),
		LocationReposiroy: repo.NewLocationRepo(pool),
		RedisClient: *redis.NewClient(&redis.Options{
			Addr:     "Localhost:6379",
			Password: "",
			DB:       0,
		}),
	}

	res, err := eventService.GetEventById(context.Background(), 3)

	assert.Error(t, err)
	assert.Nil(t, res)
}

func TestGetEventByLocationFromDBWhenRedisDown(t *testing.T) {
	dbConfig := postgres.PostgresConfig{
		ConnectionString: "postgres://postgres:password@localhost:5432/test",
		Config:           pgxpool.Config{MinConns: 5, MaxConns: 10},
	}
	pool := postgres.NewConn(&dbConfig)
	eventService := entrypoints.EventService{

		EventRepository:   repo.NewEventRepo(pool),
		LocationReposiroy: repo.NewLocationRepo(pool),
		RedisClient: *redis.NewClient(&redis.Options{
			Addr:     "Localhost:6379",
			Password: "",
			DB:       0,
		}),
	}

	loc := model.Location{
		EventId: 2,
		City:    "Sakarya",
	}
	res, err := eventService.GetEventsByLocation(context.Background(), loc)

	assert.NoError(t, err)
	assert.NotNil(t, res)
}

func TestGetEventByLocationReturnErrorWhenEventIdNotFound(t *testing.T) {
	dbConfig := postgres.PostgresConfig{
		ConnectionString: "postgres://postgres:password@localhost:5432/test",
		Config:           pgxpool.Config{MinConns: 5, MaxConns: 10},
	}
	pool := postgres.NewConn(&dbConfig)
	eventService := entrypoints.EventService{

		EventRepository:   repo.NewEventRepo(pool),
		LocationReposiroy: repo.NewLocationRepo(pool),
		RedisClient: *redis.NewClient(&redis.Options{
			Addr:     "Localhost:6379",
			Password: "",
			DB:       0,
		}),
	}

	loc := model.Location{
		EventId: 2,
		City:    "Istanbul",
	}

	res, _ := eventService.GetEventsByLocation(context.Background(), loc)

	assert.Nil(t, res)
}
