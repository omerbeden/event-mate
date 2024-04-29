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
	"github.com/omerbeden/event-mate/backend/tatooine/pkg"
	"github.com/omerbeden/event-mate/backend/tatooine/pkg/cache"
	"github.com/omerbeden/event-mate/backend/tatooine/pkg/db/postgres"
	"github.com/stretchr/testify/assert"
	"go.uber.org/zap"
)

func TestCreateActivity(t *testing.T) {
	dbConfig := postgres.PostgresConfig{
		ConnectionString: "postgres://postgres:password@localhost:5432/test",
		Config:           pgxpool.Config{MinConns: 5, MaxConns: 10},
	}
	pool := postgres.NewConn(&dbConfig)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	logger := zap.NewNop().Sugar()
	ctx = context.WithValue(ctx, pkg.LoggerKey, logger)

	pgxAdapter := postgres.NewPgxAdapter(pool)

	activityService := entrypoints.NewService(
		postgresadapter.NewActivityRepo(pgxAdapter),
		postgresadapter.NewActivityRulesRepo(pgxAdapter),
		postgresadapter.NewActivityFlowRepo(pgxAdapter),
		postgresadapter.NewLocationRepo(pgxAdapter),
		*cache.NewRedisClient(cache.RedisOption{
			Options: &redis.Options{
				Addr:     "Localhost:6379",
				Password: "",
				DB:       0,
			},
			ExpirationTime: 0,
		}),
		pgxAdapter,
	)

	activity := model.Activity{
		Title:    "running",
		Category: "sport",
		StartAt:  time.Now(),
		EndAt:    time.Now().Add(time.Hour * 2),
		Content:  "Test Content2",
		Quota:    3,
		Flow:     []string{"meet", "ice breaker", "running"},
		Rules:    []string{"no dogs", "no cigarete", "no talking while running"},
		CreatedBy: model.User{
			ID: 2,
		},
		Location: model.Location{
			City:        "Sakarya",
			District:    "Hendek",
			Description: "in front of the gas station",
			Latitude:    41.000000,
			Longitude:   28.000000,
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

	logger := zap.NewNop().Sugar()
	ctx = context.WithValue(ctx, pkg.LoggerKey, logger)

	pgxAdapter := postgres.NewPgxAdapter(pool)

	activityService := entrypoints.NewService(
		postgresadapter.NewActivityRepo(pgxAdapter),
		postgresadapter.NewActivityRulesRepo(pgxAdapter),
		postgresadapter.NewActivityFlowRepo(pgxAdapter),
		postgresadapter.NewLocationRepo(pgxAdapter),
		*cache.NewRedisClient(cache.RedisOption{
			Options: &redis.Options{
				Addr:     "Localhost:6379",
				Password: "",
				DB:       0,
			},
			ExpirationTime: 0,
		}),
		pgxAdapter,
	)
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

	logger := zap.NewNop().Sugar()
	ctx = context.WithValue(ctx, pkg.LoggerKey, logger)

	pgxAdapter := postgres.NewPgxAdapter(pool)

	activityService := entrypoints.NewService(
		postgresadapter.NewActivityRepo(pgxAdapter),
		postgresadapter.NewActivityRulesRepo(pgxAdapter),
		postgresadapter.NewActivityFlowRepo(pgxAdapter),
		postgresadapter.NewLocationRepo(pgxAdapter),
		*cache.NewRedisClient(cache.RedisOption{
			Options: &redis.Options{
				Addr:     "Localhost:6379",
				Password: "",
				DB:       0,
			},
			ExpirationTime: 0,
		}),
		pgxAdapter,
	)

	participant := model.User{
		ID:              2,
		Name:            "ome1r",
		LastName:        "be1den",
		ProfileImageUrl: "https://lh3.googleusercontent.com/a/ACg8ocL3WjhfOizacrfFw12c3m2oBr708t9vWJZffSoySetuXGQw41Ej=s100",
		ProfilePoint:    0,
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

	logger := zap.NewNop().Sugar()
	ctx = context.WithValue(ctx, pkg.LoggerKey, logger)

	pgxAdapter := postgres.NewPgxAdapter(pool)

	activityService := entrypoints.NewService(
		postgresadapter.NewActivityRepo(pgxAdapter),
		postgresadapter.NewActivityRulesRepo(pgxAdapter),
		postgresadapter.NewActivityFlowRepo(pgxAdapter),
		postgresadapter.NewLocationRepo(pgxAdapter),
		*cache.NewRedisClient(cache.RedisOption{
			Options: &redis.Options{
				Addr:     "Localhost:6379",
				Password: "",
				DB:       0,
			},
			ExpirationTime: 0,
		}),
		pgxAdapter,
	)

	activityID := int64(2)

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

	logger := zap.NewNop().Sugar()
	ctx = context.WithValue(ctx, pkg.LoggerKey, logger)

	pgxAdapter := postgres.NewPgxAdapter(pool)

	activityService := entrypoints.NewService(
		postgresadapter.NewActivityRepo(pgxAdapter),
		postgresadapter.NewActivityRulesRepo(pgxAdapter),
		postgresadapter.NewActivityFlowRepo(pgxAdapter),
		postgresadapter.NewLocationRepo(pgxAdapter),
		*cache.NewRedisClient(cache.RedisOption{
			Options: &redis.Options{
				Addr:     "Localhost:6379",
				Password: "",
				DB:       0,
			},
			ExpirationTime: 0,
		}),
		pgxAdapter,
	)

	activityId := 2

	res, err := activityService.GetActivityDetail(ctx, int64(activityId))

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

	logger := zap.NewNop().Sugar()
	ctx = context.WithValue(ctx, pkg.LoggerKey, logger)

	pgxAdapter := postgres.NewPgxAdapter(pool)
	activityService := entrypoints.NewService(
		postgresadapter.NewActivityRepo(pgxAdapter),
		postgresadapter.NewActivityRulesRepo(pgxAdapter),
		postgresadapter.NewActivityFlowRepo(pgxAdapter),
		postgresadapter.NewLocationRepo(pgxAdapter),
		*cache.NewRedisClient(cache.RedisOption{
			Options: &redis.Options{
				Addr:     "Localhost:6379",
				Password: "",
				DB:       0,
			},
			ExpirationTime: 0,
		}),
		pgxAdapter,
	)
	res, err := activityService.GetActivityDetail(ctx, 3)

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

	logger := zap.NewNop().Sugar()
	ctx = context.WithValue(ctx, pkg.LoggerKey, logger)

	pgxAdapter := postgres.NewPgxAdapter(pool)

	activityService := entrypoints.NewService(
		postgresadapter.NewActivityRepo(pgxAdapter),
		postgresadapter.NewActivityRulesRepo(pgxAdapter),
		postgresadapter.NewActivityFlowRepo(pgxAdapter),
		postgresadapter.NewLocationRepo(pgxAdapter),
		*cache.NewRedisClient(cache.RedisOption{
			Options: &redis.Options{
				Addr:     "Localhost:6379",
				Password: "",
				DB:       0,
			},
			ExpirationTime: 0,
		}),
		pgxAdapter,
	)

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

	logger := zap.NewNop().Sugar()
	ctx = context.WithValue(ctx, pkg.LoggerKey, logger)

	pgxAdapter := postgres.NewPgxAdapter(pool)

	activityService := entrypoints.NewService(
		postgresadapter.NewActivityRepo(pgxAdapter),
		postgresadapter.NewActivityRulesRepo(pgxAdapter),
		postgresadapter.NewActivityFlowRepo(pgxAdapter),
		postgresadapter.NewLocationRepo(pgxAdapter),
		*cache.NewRedisClient(cache.RedisOption{
			Options: &redis.Options{
				Addr:     "Localhost:6379",
				Password: "",
				DB:       0,
			},
			ExpirationTime: 0,
		}),
		pgxAdapter,
	)

	loc := model.Location{
		City: "Istanbul",
	}

	res, _ := activityService.GetActivitiesByLocation(ctx, loc)

	assert.Nil(t, res)
}
