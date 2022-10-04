package repo_test

import (
	"testing"

	"github.com/omerbeden/event-mate/backend/eventservice/internal/app/adapters/database"
	adapters "github.com/omerbeden/event-mate/backend/eventservice/internal/app/adapters/repo"
	"github.com/omerbeden/event-mate/backend/eventservice/internal/app/domain/model"
	"github.com/stretchr/testify/assert"
)

func TestCreateEvent(t *testing.T) {

	model := model.Event{
		Title:     "TEST",
		Category:  "TEST",
		CreatedBy: model.User{},
	}

	repo := adapters.EventRepository{
		DB: database.InitPostgressConnection(),
	}

	res, err := repo.CreateEvent(model)

	assert.True(t, res)
	assert.Nil(t, err)
	assert.NotNil(t, res)

}

func TestGetEventByID(t *testing.T) {
	repo := adapters.EventRepository{
		DB: database.InitPostgressConnection(),
	}

	res, err := repo.GetEventByID(1)

	assert.Nil(t, err)
	assert.NotNil(t, res)
	assert.NotEmpty(t, res)
}

func TestUpdateEvent(t *testing.T) {
	repo := adapters.EventRepository{
		DB: database.InitPostgressConnection(),
	}

	eventTobeUpdated := model.Event{
		Title:    "Updated title",
		Category: "Updated Category",
	}
	res, err := repo.UpdateEvent(eventTobeUpdated)

	assert.True(t, res)
	assert.Nil(t, err)
	assert.NotNil(t, res)

}

func TestDeleteEventByID(t *testing.T) {
	repo := adapters.EventRepository{
		DB: database.InitPostgressConnection(),
	}

	res, err := repo.DeleteEventByID(1)

	assert.True(t, res)
	assert.Nil(t, err)
	assert.NotNil(t, res)

}
