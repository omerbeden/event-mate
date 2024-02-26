package entrypoints_test

import (
	"context"
	"testing"
	"time"

	"github.com/go-redis/redis/v8"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/omerbeden/event-mate/backend/tatooine/modules/activity/app/adapters/repo"
	"github.com/omerbeden/event-mate/backend/tatooine/modules/activity/app/domain/model"
	"github.com/omerbeden/event-mate/backend/tatooine/modules/activity/app/entrypoints"
	postgres "github.com/omerbeden/event-mate/backend/tatooine/pkg/database"
	"github.com/stretchr/testify/assert"
)

func TestCreateActivity(t *testing.T) {
	dbConfig := postgres.PostgresConfig{
		ConnectionString: "postgres://postgres:password@localhost:5432/test",
		Config:           pgxpool.Config{MinConns: 5, MaxConns: 10},
	}
	pool := postgres.NewConn(&dbConfig)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()
	activityService := entrypoints.ActivityService{

		ActivityRepository: repo.NewActivityRepo(pool),
		LocationReposiroy:  repo.NewLocationRepo(pool),
		RedisClient: *redis.NewClient(&redis.Options{
			Addr:     "Localhost:6379",
			Password: "",
			DB:       0,
		}),
	}

	activity := model.Activity{
		Title:    "Test Title2",
		Category: "Test Category2",
		CreatedBy: model.User{
			ID: 1,
		},
		Location: model.Location{
			City: "Sakarya",
		},
	}

	res, err := activityService.CreateActivity(ctx, activity)
	assert.NoError(t, err)
	assert.True(t, res)
}

func TestGetActivitiesByLocation(t *testing.T) {
	dbConfig := postgres.PostgresConfig{
		ConnectionString: "postgres://postgres:password@localhost:5432/test",
		Config:           pgxpool.Config{MinConns: 5, MaxConns: 10},
	}
	pool := postgres.NewConn(&dbConfig)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()
	activityService := entrypoints.ActivityService{

		ActivityRepository: repo.NewActivityRepo(pool),
		LocationReposiroy:  repo.NewLocationRepo(pool),
		RedisClient: *redis.NewClient(&redis.Options{
			Addr:     "Localhost:6379",
			Password: "",
			DB:       0,
		}),
	}
	loc := model.Location{
		ActivityId: 2,
		City:       "Sakarya",
	}

	result, err := activityService.GetActivitiesByLocation(ctx, loc)

	assert.NoError(t, err)
	assert.NotEmpty(t, result)
	assert.NotEmpty(t, result[0].Participants)

}

func TestAddParticipant(t *testing.T) {
	dbConfig := postgres.PostgresConfig{
		ConnectionString: "postgres://postgres:password@localhost:5432/test",
		Config:           pgxpool.Config{MinConns: 5, MaxConns: 10},
	}
	pool := postgres.NewConn(&dbConfig)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()
	activityService := entrypoints.ActivityService{

		ActivityRepository: repo.NewActivityRepo(pool),
		LocationReposiroy:  repo.NewLocationRepo(pool),
		RedisClient: *redis.NewClient(&redis.Options{
			Addr:     "Localhost:6379",
			Password: "",
			DB:       0,
		}),
	}

	participant := model.User{
		ID: 2,
	}
	err := activityService.AddParticipant(ctx, participant, 2)
	assert.NoError(t, err)
}

func TestGetParticipants(t *testing.T) {
	dbConfig := postgres.PostgresConfig{
		ConnectionString: "postgres://postgres:password@localhost:5432/test",
		Config:           pgxpool.Config{MinConns: 5, MaxConns: 10},
	}
	pool := postgres.NewConn(&dbConfig)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()
	activityService := entrypoints.ActivityService{

		ActivityRepository: repo.NewActivityRepo(pool),
		LocationReposiroy:  repo.NewLocationRepo(pool),
		RedisClient: *redis.NewClient(&redis.Options{
			Addr:     "Localhost:6379",
			Password: "",
			DB:       0,
		}),
	}

	res, err := activityService.GetParticipants(ctx, 1)
	assert.NoError(t, err)
	assert.NotEmpty(t, res)

}
func TestGetActivityFromDBWhenRedisDown(t *testing.T) {
	dbConfig := postgres.PostgresConfig{
		ConnectionString: "postgres://postgres:password@localhost:5432/test",
		Config:           pgxpool.Config{MinConns: 5, MaxConns: 10},
	}
	pool := postgres.NewConn(&dbConfig)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()
	activityService := entrypoints.ActivityService{

		ActivityRepository: repo.NewActivityRepo(pool),
		LocationReposiroy:  repo.NewLocationRepo(pool),
		RedisClient: *redis.NewClient(&redis.Options{
			Addr:     "Localhost:6379",
			Password: "",
			DB:       0,
		}),
	}

	res, err := activityService.GetActivityById(ctx, 1)

	assert.NoError(t, err)
	assert.NotNil(t, res)
}

func TestGetActivityByIDReturnErrorWhenActivityIdNotFound(t *testing.T) {
	dbConfig := postgres.PostgresConfig{
		ConnectionString: "postgres://postgres:password@localhost:5432/test",
		Config:           pgxpool.Config{MinConns: 5, MaxConns: 10},
	}
	pool := postgres.NewConn(&dbConfig)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()
	activityService := entrypoints.ActivityService{

		ActivityRepository: repo.NewActivityRepo(pool),
		LocationReposiroy:  repo.NewLocationRepo(pool),
		RedisClient: *redis.NewClient(&redis.Options{
			Addr:     "Localhost:6379",
			Password: "",
			DB:       0,
		}),
	}

	res, err := activityService.GetActivityById(ctx, 3)

	assert.Error(t, err)
	assert.Nil(t, res)
}

func TestGetActivityByLocationFromDBWhenRedisDown(t *testing.T) {
	dbConfig := postgres.PostgresConfig{
		ConnectionString: "postgres://postgres:password@localhost:5432/test",
		Config:           pgxpool.Config{MinConns: 5, MaxConns: 10},
	}
	pool := postgres.NewConn(&dbConfig)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()
	activityService := entrypoints.ActivityService{

		ActivityRepository: repo.NewActivityRepo(pool),
		LocationReposiroy:  repo.NewLocationRepo(pool),
		RedisClient: *redis.NewClient(&redis.Options{
			Addr:     "Localhost:6379",
			Password: "",
			DB:       0,
		}),
	}

	loc := model.Location{
		ActivityId: 2,
		City:       "Sakarya",
	}

	res, err := activityService.GetActivitiesByLocation(ctx, loc)

	assert.NoError(t, err)
	assert.NotNil(t, res)
}

func TestGetActivityByLocationReturnErrorWhenCityNotFound(t *testing.T) {
	dbConfig := postgres.PostgresConfig{
		ConnectionString: "postgres://postgres:password@localhost:5432/test",
		Config:           pgxpool.Config{MinConns: 5, MaxConns: 10},
	}
	pool := postgres.NewConn(&dbConfig)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()
	activityService := entrypoints.ActivityService{

		ActivityRepository: repo.NewActivityRepo(pool),
		LocationReposiroy:  repo.NewLocationRepo(pool),
		RedisClient: *redis.NewClient(&redis.Options{
			Addr:     "Localhost:6379",
			Password: "",
			DB:       0,
		}),
	}

	loc := model.Location{
		ActivityId: 2,
		City:       "Istanbul",
	}

	res, _ := activityService.GetActivitiesByLocation(ctx, loc)

	assert.Nil(t, res)
}
