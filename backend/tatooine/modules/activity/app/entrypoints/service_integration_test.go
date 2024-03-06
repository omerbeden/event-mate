package entrypoints_test

import (
	"context"
	"testing"
	"time"

	"github.com/go-redis/redis/v8"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/omerbeden/event-mate/backend/tatooine/modules/activity/app/adapters/postgresadapter"
	"github.com/omerbeden/event-mate/backend/tatooine/modules/activity/app/domain/model"
	"github.com/omerbeden/event-mate/backend/tatooine/modules/activity/app/entrypoints"
	"github.com/omerbeden/event-mate/backend/tatooine/pkg/cache"
	"github.com/omerbeden/event-mate/backend/tatooine/pkg/db/postgres"
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

		ActivityRepository:      postgresadapter.NewActivityRepo(postgresadapter.NewPgxAdapter(pool)),
		LocationReposiroy:       postgresadapter.NewLocationRepo(postgresadapter.NewPgxAdapter(pool)),
		ActivityRulesRepository: postgresadapter.NewActivityRulesRepo(postgresadapter.NewPgxAdapter(pool)),
		ActivityFlowRepository:  postgresadapter.NewActivityFlowRepo(postgresadapter.NewPgxAdapter(pool)),
		RedisClient: *cache.NewRedisClient(cache.RedisOption{
			Options: &redis.Options{
				Addr:     "Localhost:6379",
				Password: "",
				DB:       0,
			},
			ExpirationTime: 0,
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

		ActivityRepository: postgresadapter.NewActivityRepo(postgresadapter.NewPgxAdapter(pool)),
		LocationReposiroy:  postgresadapter.NewLocationRepo(postgresadapter.NewPgxAdapter(pool)),
		RedisClient: *cache.NewRedisClient(cache.RedisOption{
			Options: &redis.Options{
				Addr:     "Localhost:6379",
				Password: "",
				DB:       0,
			},
			ExpirationTime: 0,
		}),
	}
	loc := model.Location{
		City: "Sakarya",
	}

	result, err := activityService.GetActivitiesByLocation(ctx, loc)

	assert.NoError(t, err)
	assert.NotEmpty(t, result)
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

		ActivityRepository: postgresadapter.NewActivityRepo(postgresadapter.NewPgxAdapter(pool)),
		LocationReposiroy:  postgresadapter.NewLocationRepo(postgresadapter.NewPgxAdapter(pool)),
		RedisClient: *cache.NewRedisClient(cache.RedisOption{
			Options: &redis.Options{
				Addr:     "Localhost:6379",
				Password: "",
				DB:       0,
			},
			ExpirationTime: 0,
		}),
	}

	participant := model.User{
		ID: 4,
	}
	activityId := int64(1)

	err := activityService.AddParticipant(ctx, participant, activityId)
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

		ActivityRepository: postgresadapter.NewActivityRepo(postgresadapter.NewPgxAdapter(pool)),
		LocationReposiroy:  postgresadapter.NewLocationRepo(postgresadapter.NewPgxAdapter(pool)),
		RedisClient: *cache.NewRedisClient(cache.RedisOption{
			Options: &redis.Options{
				Addr:     "Localhost:6379",
				Password: "",
				DB:       0,
			},
			ExpirationTime: 0,
		}),
	}

	activityID := int64(1)

	res, err := activityService.GetParticipants(ctx, activityID)

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

		ActivityRepository:      postgresadapter.NewActivityRepo(postgresadapter.NewPgxAdapter(pool)),
		LocationReposiroy:       postgresadapter.NewLocationRepo(postgresadapter.NewPgxAdapter(pool)),
		ActivityRulesRepository: postgresadapter.NewActivityRulesRepo(postgresadapter.NewPgxAdapter(pool)),
		ActivityFlowRepository:  postgresadapter.NewActivityFlowRepo(postgresadapter.NewPgxAdapter(pool)),
		RedisClient: *cache.NewRedisClient(cache.RedisOption{
			Options:        &redis.Options{},
			ExpirationTime: 0,
		}),
	}

	activityId := 2

	res, err := activityService.GetActivityById(ctx, int64(activityId))

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

		ActivityRepository: postgresadapter.NewActivityRepo(postgresadapter.NewPgxAdapter(pool)),
		LocationReposiroy:  postgresadapter.NewLocationRepo(postgresadapter.NewPgxAdapter(pool)),
		RedisClient: *cache.NewRedisClient(cache.RedisOption{
			Options: &redis.Options{
				Addr:     "Localhost:6379",
				Password: "",
				DB:       0,
			},
			ExpirationTime: 0,
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

		ActivityRepository:      postgresadapter.NewActivityRepo(postgresadapter.NewPgxAdapter(pool)),
		LocationReposiroy:       postgresadapter.NewLocationRepo(postgresadapter.NewPgxAdapter(pool)),
		ActivityRulesRepository: postgresadapter.NewActivityRulesRepo(postgresadapter.NewPgxAdapter(pool)),
		ActivityFlowRepository:  postgresadapter.NewActivityFlowRepo(postgresadapter.NewPgxAdapter(pool)),
		RedisClient: *cache.NewRedisClient(cache.RedisOption{
			Options:        &redis.Options{},
			ExpirationTime: 0,
		}),
	}

	loc := model.Location{
		City: "Sakarya",
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

		ActivityRepository:      postgresadapter.NewActivityRepo(postgresadapter.NewPgxAdapter(pool)),
		LocationReposiroy:       postgresadapter.NewLocationRepo(postgresadapter.NewPgxAdapter(pool)),
		ActivityRulesRepository: postgresadapter.NewActivityRulesRepo(postgresadapter.NewPgxAdapter(pool)),
		ActivityFlowRepository:  postgresadapter.NewActivityFlowRepo(postgresadapter.NewPgxAdapter(pool)),
		RedisClient: *cache.NewRedisClient(cache.RedisOption{
			Options: &redis.Options{
				Addr:     "Localhost:6379",
				Password: "",
				DB:       0,
			},
			ExpirationTime: 0,
		}),
	}

	loc := model.Location{
		City: "Istanbul",
	}

	res, _ := activityService.GetActivitiesByLocation(ctx, loc)

	assert.Nil(t, res)
}
