package repo_test

import (
	"testing"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/omerbeden/event-mate/backend/tatooine/modules/profile/app/adapters/repo"
	"github.com/omerbeden/event-mate/backend/tatooine/modules/profile/app/domain/model"
	postgres "github.com/omerbeden/event-mate/backend/tatooine/pkg/database"
	"github.com/stretchr/testify/assert"
)

func TestInsertUser(t *testing.T) {
	dbConfig := postgres.PostgresConfig{
		ConnectionString: "postgres://postgres:password@localhost:5432/test",
		Config:           pgxpool.Config{MinConns: 5, MaxConns: 10},
	}
	pool := postgres.NewConn(&dbConfig)

	repository := repo.NewUserProfileRepo(pool)

	user := model.UserProfile{
		Name:               "oner1",
		LastName:           "beden",
		About:              "about",
		AttandedActivities: []model.Activity{},
		Adress: model.UserProfileAdress{
			City: "Sakarya",
		},
		Stat: model.UserProfileStat{
			Point: 3.5,
		},
		ProfileImageUrl: "image url",
	}

	result, err := repository.InsertUser(&user)

	assert.NoError(t, err)
	assert.NotNil(t, result)

}

func TestGetUsersByAddress(t *testing.T) {
	dbConfig := postgres.PostgresConfig{
		ConnectionString: "postgres://postgres:password@localhost:5432/test",
		Config:           pgxpool.Config{MinConns: 5, MaxConns: 10},
	}
	pool := postgres.NewConn(&dbConfig)

	repository := repo.NewUserProfileRepo(pool)

	address := model.UserProfileAdress{
		City: "Sakarya",
	}

	result, err := repository.GetUsersByAddress(address)

	assert.NoError(t, err)
	assert.NotNil(t, result)
	assert.NotEmpty(t, result)
	assert.Equal(t, result[0].Adress.City, address.City)
}

func TestUpdateProfileImage(t *testing.T) {
	dbConfig := postgres.PostgresConfig{
		ConnectionString: "postgres://postgres:password@localhost:5432/test",
		Config:           pgxpool.Config{MinConns: 5, MaxConns: 10},
	}
	pool := postgres.NewConn(&dbConfig)

	repository := repo.NewUserProfileRepo(pool)

	updatedUSer, err := repository.UpdateProfileImage("", "new image.png")
	assert.NoError(t, err)
	assert.NotNil(t, updatedUSer)
}

func TestDeleteUserById(t *testing.T) {
	dbConfig := postgres.PostgresConfig{
		ConnectionString: "postgres://postgres:password@localhost:5432/test",
		Config:           pgxpool.Config{MinConns: 5, MaxConns: 10},
	}
	pool := postgres.NewConn(&dbConfig)

	repository := repo.NewUserProfileRepo(pool)

	err := repository.DeleteUserById(1)
	assert.NoError(t, err)
}
