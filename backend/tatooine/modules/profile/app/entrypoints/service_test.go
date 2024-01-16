package entrypoints_test

import (
	"testing"

	"github.com/go-redis/redis/v8"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/omerbeden/event-mate/backend/tatooine/modules/profile/app/adapters/repo"
	"github.com/omerbeden/event-mate/backend/tatooine/modules/profile/app/domain/model"
	"github.com/omerbeden/event-mate/backend/tatooine/modules/profile/app/entrypoints"
	"github.com/omerbeden/event-mate/backend/tatooine/pkg/cache"
	postgres "github.com/omerbeden/event-mate/backend/tatooine/pkg/database"
	"github.com/stretchr/testify/assert"
)

func TestCreateUserProfile(t *testing.T) {
	dbConfig := postgres.PostgresConfig{
		ConnectionString: "postgres://postgres:password@localhost:5432/test",
		Config:           pgxpool.Config{MinConns: 5, MaxConns: 10},
	}

	pool := postgres.NewConn(&dbConfig)

	redis := cache.NewRedisClient(cache.RedisOption{
		Options: &redis.Options{
			Addr:     "Localhost:6379",
			Password: "",
			DB:       0,
		},
		ExpirationTime: 0,
	})

	service := entrypoints.NewService(repo.NewUserProfileRepo(pool), *redis)
	user := &model.UserProfile{
		Name:               "omer",
		LastName:           "beden",
		About:              "backend developer",
		AttandedActivities: []model.Activity{},
		Adress:             model.UserProfileAdress{City: "Sakarya"},
		Stat: model.UserProfileStat{
			Followers:      1,
			Followings:     2,
			AttandedEvents: 3,
			Point:          5,
		},
		ProfileImageUrl: "profileImage.png",
	}

	err := service.CreateUser(user)

	assert.NoError(t, err)
}

func TestUpdateUserProfileImage(t *testing.T) {
	dbConfig := postgres.PostgresConfig{
		ConnectionString: "postgres://postgres:password@localhost:5432/test",
		Config:           pgxpool.Config{MinConns: 5, MaxConns: 10},
	}

	pool := postgres.NewConn(&dbConfig)

	redis := cache.NewRedisClient(cache.RedisOption{
		Options: &redis.Options{
			Addr:     "Localhost:6379",
			Password: "",
			DB:       0,
		},
		ExpirationTime: 0,
	})

	service := entrypoints.NewService(repo.NewUserProfileRepo(pool), *redis)

	err := service.UpdateProfileImage(4, "new profile image4.png")

	assert.NoError(t, err)
}
func TestDeleteUser(t *testing.T) {
	dbConfig := postgres.PostgresConfig{
		ConnectionString: "postgres://postgres:password@localhost:5432/test",
		Config:           pgxpool.Config{MinConns: 5, MaxConns: 10},
	}

	pool := postgres.NewConn(&dbConfig)

	redis := cache.NewRedisClient(cache.RedisOption{
		Options: &redis.Options{
			Addr:     "Localhost:6379",
			Password: "",
			DB:       0,
		},
		ExpirationTime: 0,
	})

	service := entrypoints.NewService(repo.NewUserProfileRepo(pool), *redis)
	err := service.DeleteUser(6)

	assert.NoError(t, err)

}
