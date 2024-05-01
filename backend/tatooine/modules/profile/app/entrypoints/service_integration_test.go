package entrypoints_test

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/go-redis/redis/v8"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/omerbeden/event-mate/backend/tatooine/modules/profile/app/adapters/postgresadapter"
	"github.com/omerbeden/event-mate/backend/tatooine/modules/profile/app/domain/model"
	"github.com/omerbeden/event-mate/backend/tatooine/modules/profile/app/entrypoints"
	"github.com/omerbeden/event-mate/backend/tatooine/pkg"
	"github.com/omerbeden/event-mate/backend/tatooine/pkg/cache"
	"github.com/omerbeden/event-mate/backend/tatooine/pkg/db/postgres"
	"github.com/stretchr/testify/assert"
	"go.uber.org/zap"
)

func TestCreateUserProfile(t *testing.T) {
	dbConfig := postgres.PostgresConfig{
		ConnectionString: "postgres://postgres:password@localhost:5432/test",
		Config:           pgxpool.Config{MinConns: 5, MaxConns: 10},
	}

	pool := postgres.NewConn(&dbConfig)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	logger := zap.NewNop().Sugar()
	ctx = context.WithValue(ctx, pkg.LoggerKey, logger)

	redis := cache.NewRedisClient(cache.RedisOption{
		Options: &redis.Options{
			Addr:     "Localhost:6379",
			Password: "",
			DB:       0,
		},
		ExpirationTime: 0,
	})

	pgxAdapter := postgres.NewPgxAdapter(pool)

	service := entrypoints.NewService(
		postgresadapter.NewUserProfileRepo(pgxAdapter),
		postgresadapter.NewUserProfileStatRepo(pgxAdapter),
		postgresadapter.NewUserProfileAddressRepo(pgxAdapter),
		postgresadapter.NewBadgeRepo(pgxAdapter),
		*redis,
		pgxAdapter,
	)
	user := &model.CreateUserProfileRequest{
		Name:            "ome1r",
		LastName:        "be1den",
		About:           "p1c",
		ExternalId:      "7",
		UserName:        "omrr7",
		Email:           "test7",
		Adress:          model.UserProfileAdress{City: "Cupertino"},
		ProfileImageUrl: "profileI1mage.png",
		IsVerified:      true,
	}

	err := service.CreateUser(ctx, user)

	assert.NoError(t, err)
}

func TestCreateUserProfileWithoutRedis(t *testing.T) {
	dbConfig := postgres.PostgresConfig{
		ConnectionString: "postgres://postgres:password@localhost:5432/test",
		Config:           pgxpool.Config{MinConns: 5, MaxConns: 10},
	}

	pool := postgres.NewConn(&dbConfig)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	logger := zap.NewNop().Sugar()
	ctx = context.WithValue(ctx, pkg.LoggerKey, logger)

	redis := cache.NewRedisClient(cache.RedisOption{
		Options: &redis.Options{
			Addr:     "unknown:6379",
			Password: "",
			DB:       0,
		},
		ExpirationTime: 0,
	})

	pgxAdapter := postgres.NewPgxAdapter(pool)
	service := entrypoints.NewService(
		postgresadapter.NewUserProfileRepo(pgxAdapter),
		postgresadapter.NewUserProfileStatRepo(pgxAdapter),
		postgresadapter.NewUserProfileAddressRepo(pgxAdapter),
		postgresadapter.NewBadgeRepo(pgxAdapter),
		*redis, pgxAdapter,
	)
	user := &model.CreateUserProfileRequest{
		Name:            "omer",
		LastName:        "beden",
		About:           "backend developer",
		ExternalId:      "redis14",
		UserName:        "omerbeden1",
		Email:           "test14@test1.com",
		Adress:          model.UserProfileAdress{City: "Sakarya"},
		ProfileImageUrl: "profileImage.png",
	}

	err := service.CreateUser(ctx, user)

	assert.NoError(t, err)
}

func TestUpdateUserProfileImage(t *testing.T) {
	dbConfig := postgres.PostgresConfig{
		ConnectionString: "postgres://postgres:password@localhost:5432/test",
		Config:           pgxpool.Config{MinConns: 5, MaxConns: 10},
	}

	pool := postgres.NewConn(&dbConfig)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()
	logger := zap.NewNop().Sugar()
	ctx = context.WithValue(ctx, pkg.LoggerKey, logger)

	redis := cache.NewRedisClient(cache.RedisOption{
		Options: &redis.Options{
			Addr:     "Localhost:6379",
			Password: "",
			DB:       0,
		},
		ExpirationTime: 0,
	})

	pgxAdapter := postgres.NewPgxAdapter(pool)
	service := entrypoints.NewService(
		postgresadapter.NewUserProfileRepo(pgxAdapter),
		postgresadapter.NewUserProfileStatRepo(pgxAdapter),
		postgresadapter.NewUserProfileAddressRepo(pgxAdapter),
		postgresadapter.NewBadgeRepo(pgxAdapter),
		*redis, pgxAdapter)

	err := service.UpdateProfileImage(ctx, "1a", "new profile image10.png")

	assert.NoError(t, err)
}

func TestUpdateUserProfileImageWithoutRedis(t *testing.T) {
	dbConfig := postgres.PostgresConfig{
		ConnectionString: "postgres://postgres:password@localhost:5432/test",
		Config:           pgxpool.Config{MinConns: 5, MaxConns: 10},
	}

	pool := postgres.NewConn(&dbConfig)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	logger := zap.NewNop().Sugar()
	ctx = context.WithValue(ctx, pkg.LoggerKey, logger)

	redis := cache.NewRedisClient(cache.RedisOption{
		Options: &redis.Options{
			Addr:     "unknown:6379",
			Password: "",
			DB:       0,
		},
		ExpirationTime: 0,
	})

	pgxAdapter := postgres.NewPgxAdapter(pool)

	service := entrypoints.NewService(
		postgresadapter.NewUserProfileRepo(pgxAdapter),
		postgresadapter.NewUserProfileStatRepo(pgxAdapter),
		postgresadapter.NewUserProfileAddressRepo(pgxAdapter),
		postgresadapter.NewBadgeRepo(pgxAdapter),
		*redis, pgxAdapter)
	err := service.UpdateProfileImage(ctx, "redis3", "new profile image10.png")

	assert.Error(t, err)
}

func TestGetAttandedActivities(t *testing.T) {
	dbConfig := postgres.PostgresConfig{
		ConnectionString: "postgres://postgres:password@localhost:5432/test",
		Config:           pgxpool.Config{MinConns: 5, MaxConns: 10},
	}

	pool := postgres.NewConn(&dbConfig)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()
	logger := zap.NewNop().Sugar()
	ctx = context.WithValue(ctx, pkg.LoggerKey, logger)

	redis := cache.NewRedisClient(cache.RedisOption{
		Options: &redis.Options{
			Addr:     "Localhost:6379",
			Password: "",
			DB:       0,
		},
		ExpirationTime: 0,
	})
	pgxAdapter := postgres.NewPgxAdapter(pool)
	service := entrypoints.NewService(
		postgresadapter.NewUserProfileRepo(pgxAdapter),
		postgresadapter.NewUserProfileStatRepo(pgxAdapter),
		postgresadapter.NewUserProfileAddressRepo(pgxAdapter),
		postgresadapter.NewBadgeRepo(pgxAdapter),
		*redis, pgxAdapter)

	userId := int64(2)
	attandedActivities, err := service.GetAttandedActivities(ctx, userId)

	assert.NoError(t, err)
	assert.NotNil(t, attandedActivities)
	assert.NotEmpty(t, attandedActivities)
}

func TestGetUserProfile(t *testing.T) {
	dbConfig := postgres.PostgresConfig{
		ConnectionString: "postgres://postgres:password@localhost:5432/test",
		Config:           pgxpool.Config{MinConns: 5, MaxConns: 10},
	}

	pool := postgres.NewConn(&dbConfig)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()
	logger := zap.NewNop().Sugar()
	ctx = context.WithValue(ctx, pkg.LoggerKey, logger)

	redis := cache.NewRedisClient(cache.RedisOption{
		Options: &redis.Options{
			Addr:     "Localhost:6379",
			Password: "",
			DB:       0,
		},
		ExpirationTime: 0,
	})
	pgxAdapter := postgres.NewPgxAdapter(pool)

	service := entrypoints.NewService(
		postgresadapter.NewUserProfileRepo(pgxAdapter),
		postgresadapter.NewUserProfileStatRepo(pgxAdapter),
		postgresadapter.NewUserProfileAddressRepo(pgxAdapter),
		postgresadapter.NewBadgeRepo(pgxAdapter),
		*redis, pgxAdapter)

	userName := "onerbed"
	user, err := service.GetUserProfile(ctx, userName)
	fmt.Printf("user: %+v", user)
	assert.NoError(t, err)
	assert.NotNil(t, user)
}

func TestEvaluateUser(t *testing.T) {
	dbConfig := postgres.PostgresConfig{
		ConnectionString: "postgres://postgres:password@localhost:5432/test",
		Config:           pgxpool.Config{MinConns: 5, MaxConns: 10},
	}

	pool := postgres.NewConn(&dbConfig)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	logger := zap.NewNop().Sugar()
	ctx = context.WithValue(ctx, pkg.LoggerKey, logger)

	redis := cache.NewRedisClient(cache.RedisOption{
		Options: &redis.Options{
			Addr:     "Localhost:6379",
			Password: "",
			DB:       0,
		},
		ExpirationTime: 0,
	})

	evaluation := model.UserEvaluation{
		ReceiverId:        "e2",
		GiverId:           "16",
		Points:            8.7,
		Comment:           "test comment",
		RelatedActivityId: 2,
	}
	pgxAdapter := postgres.NewPgxAdapter(pool)

	service := entrypoints.NewService(
		postgresadapter.NewUserProfileRepo(pgxAdapter),
		postgresadapter.NewUserProfileStatRepo(pgxAdapter),
		postgresadapter.NewUserProfileAddressRepo(pgxAdapter),
		postgresadapter.NewBadgeRepo(pgxAdapter),
		*redis, pgxAdapter)

	err := service.EvaluateUser(ctx, evaluation)

	assert.NoError(t, err)

}

func TestDeleteUser(t *testing.T) {
	dbConfig := postgres.PostgresConfig{
		ConnectionString: "postgres://postgres:password@localhost:5432/test",
		Config:           pgxpool.Config{MinConns: 5, MaxConns: 10},
	}

	pool := postgres.NewConn(&dbConfig)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	logger := zap.NewNop().Sugar()
	ctx = context.WithValue(ctx, pkg.LoggerKey, logger)

	redis := cache.NewRedisClient(cache.RedisOption{
		Options: &redis.Options{
			Addr:     "Localhost:6379",
			Password: "",
			DB:       0,
		},
		ExpirationTime: 0,
	})
	pgxAdapter := postgres.NewPgxAdapter(pool)

	service := entrypoints.NewService(
		postgresadapter.NewUserProfileRepo(pgxAdapter),
		postgresadapter.NewUserProfileStatRepo(pgxAdapter),
		postgresadapter.NewUserProfileAddressRepo(pgxAdapter),
		postgresadapter.NewBadgeRepo(pgxAdapter),
		*redis, pgxAdapter)
	err := service.DeleteUser(ctx, "1a", "omr")

	assert.NoError(t, err)

}
