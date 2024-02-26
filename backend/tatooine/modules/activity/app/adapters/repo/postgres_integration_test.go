package repo_test

import (
	"context"
	"testing"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/omerbeden/event-mate/backend/tatooine/modules/activity/app/adapters/repo"
	"github.com/omerbeden/event-mate/backend/tatooine/modules/activity/app/domain/model"
	postgres "github.com/omerbeden/event-mate/backend/tatooine/pkg/database"
	"github.com/stretchr/testify/assert"
)

func TestCreateActivity(t *testing.T) {
	dbConfig := postgres.PostgresConfig{
		ConnectionString: "postgres://postgres:password@localhost:5432/test",
		Config:           pgxpool.Config{MinConns: 5, MaxConns: 10},
	}
	pool := postgres.NewConn(&dbConfig)

	activityRepository := repo.NewActivityRepo(pool)
	locationRepository := repo.NewLocationRepo(pool)

	defer activityRepository.Close()
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	activity := model.Activity{
		Title:              "test title",
		Category:           "test category",
		CreatedBy:          model.User{ID: 2},
		Location:           model.Location{City: "Sakarya"},
		BackgroundImageUrl: "image url",
		StartAt:            time.Now(),
		Content:            "test activity content",
	}

	res, err := activityRepository.Create(ctx, activity)
	assert.NoError(t, err)
	assert.NotNil(t, res)

	resLoc, err := locationRepository.Create(ctx, &res.Location)
	assert.NoError(t, err)
	assert.True(t, resLoc)

}

func TestAddParticipants(t *testing.T) {
	dbConfig := postgres.PostgresConfig{
		ConnectionString: "postgres://postgres:password@localhost:5432/test",
		Config:           pgxpool.Config{MinConns: 5, MaxConns: 10},
	}
	pool := postgres.NewConn(&dbConfig)
	activityRepository := repo.NewActivityRepo(pool)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	activity := model.Activity{
		ID:           1,
		Title:        "test title",
		Category:     "test category",
		CreatedBy:    model.User{ID: 1},
		Location:     model.Location{City: "Sakarya"},
		Participants: []model.User{{ID: 1}, {ID: 2}, {ID: 3}}}

	err := activityRepository.AddParticipants(ctx, activity)
	assert.NoError(t, err)

}

func TestAddParticipant(t *testing.T) {
	dbConfig := postgres.PostgresConfig{
		ConnectionString: "postgres://postgres:password@localhost:5432/test",
		Config:           pgxpool.Config{MinConns: 5, MaxConns: 10},
	}
	pool := postgres.NewConn(&dbConfig)
	activityRepository := repo.NewActivityRepo(pool)
	activity := model.Activity{
		ID:        1,
		Title:     "test title",
		Category:  "test category",
		CreatedBy: model.User{ID: 1},
		Location:  model.Location{City: "Sakarya"},
	}
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	user := model.User{ID: 1}

	err := activityRepository.AddParticipant(ctx, activity.ID, user)
	assert.NoError(t, err)

}

func TestGetActivityByID(t *testing.T) {
	dbConfig := postgres.PostgresConfig{
		ConnectionString: "postgres://postgres:password@localhost:5432/test",
		Config:           pgxpool.Config{MinConns: 5, MaxConns: 10},
	}
	pool := postgres.NewConn(&dbConfig)
	repository := repo.NewActivityRepo(pool)
	defer repository.Close()
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	res, err := repository.GetByID(ctx, 1)

	assert.NotNil(t, res)
	assert.NoError(t, err)
}

func TestGetActivitiesByLocation(t *testing.T) {
	dbConfig := postgres.PostgresConfig{
		ConnectionString: "postgres://postgres:password@localhost:5432/test",
		Config:           pgxpool.Config{MinConns: 5, MaxConns: 10},
	}
	pool := postgres.NewConn(&dbConfig)
	repository := repo.NewActivityRepo(pool)
	defer repository.Close()
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	res, err := repository.GetByLocation(ctx, &model.Location{City: "Sakarya"})

	assert.NotNil(t, res)
	assert.NoError(t, err)
}

func TestUpdateActivity(t *testing.T) {

	dbConfig := postgres.PostgresConfig{
		ConnectionString: "postgres://postgres:password@localhost:5432/test",
		Config:           pgxpool.Config{MinConns: 5, MaxConns: 10},
	}
	pool := postgres.NewConn(&dbConfig)
	repository := repo.NewActivityRepo(pool)
	defer repository.Close()
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	activityTobeUpdated := model.Activity{
		Title:     "Updated title",
		Category:  "Updated Category",
		CreatedBy: model.User{ID: 2},
	}

	res, err := repository.UpdateByID(ctx, 1, activityTobeUpdated)
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
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	loc := model.Location{
		ActivityId: 1,
		City:       "Sakarya",
	}

	res, err := repository.Create(ctx, &loc)
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
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	locationToBeUpdated := model.Location{
		ActivityId: 1,
		City:       "Istanbul",
	}
	res, err := repository.UpdateByID(ctx, locationToBeUpdated)
	assert.NotNil(t, res)
	assert.NoError(t, err)
}

func TestDeleteActivityByID(t *testing.T) {

	dbConfig := postgres.PostgresConfig{
		ConnectionString: "postgres://postgres:password@localhost:5432/test",
		Config:           pgxpool.Config{MinConns: 5, MaxConns: 10},
	}
	pool := postgres.NewConn(&dbConfig)
	repository := repo.NewActivityRepo(pool)
	defer repository.Close()
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	res, err := repository.DeleteByID(ctx, 1)
	assert.NotNil(t, res)
	assert.NoError(t, err)

}
