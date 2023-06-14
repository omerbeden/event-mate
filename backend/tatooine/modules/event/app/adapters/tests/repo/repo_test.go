package repo_test

import (
	"testing"

	"github.com/omerbeden/event-mate/backend/eventservice/internal/app/adapters/database"
	adapters "github.com/omerbeden/event-mate/backend/eventservice/internal/app/adapters/repo"
	"github.com/omerbeden/event-mate/backend/eventservice/internal/app/domain/model"
	"github.com/stretchr/testify/suite"
)

type EventRepositoryTestSuite struct {
	suite.Suite
	repo *adapters.EventRepository
}

func (suite *EventRepositoryTestSuite) SetupTest() {
	suite.repo = &adapters.EventRepository{
		DB: database.InitPostgressConnection(),
	}

	suite.repo.DB.AutoMigrate(&model.Event{}, &model.Location{})
}

func (suite *EventRepositoryTestSuite) TearDownTest() {
	suite.repo = nil
}

func TestEventRepository(t *testing.T) {
	suite.Run(t, new(EventRepositoryTestSuite))
}

func (suite *EventRepositoryTestSuite) TestCreateEvent() {

	event := model.Event{
		Title:     "TEST",
		Category:  "TEST",
		CreatedBy: model.User{},
		Location:  model.Location{City: "Sakarya"},
	}

	res, err := suite.repo.CreateEvent(event)
	suite.True(res)
	suite.Nil(err)
	suite.NotNil(res)

}

func (suite *EventRepositoryTestSuite) TestGetEventByID() {

	var lastItem model.Event
	suite.repo.DB.Last(&lastItem)
	res, err := suite.repo.GetEventByID(int32(lastItem.ID))

	suite.Nil(err)
	suite.NotNil(res)
	suite.NotEmpty(res)
}

func (suite *EventRepositoryTestSuite) TestUpdateEvent() {

	eventTobeUpdated := model.Event{
		Title:    "Updated title",
		Category: "Updated Category",
	}

	var lastItem model.Event
	suite.repo.DB.Last(&lastItem)
	res, err := suite.repo.UpdateEventByID(int32(lastItem.ID), eventTobeUpdated)

	suite.True(res)
	suite.Nil(err)
	suite.NotNil(res)

}

func (suite *EventRepositoryTestSuite) TestDeleteEventByID() {

	var lastItem model.Event
	suite.repo.DB.Last(&lastItem)
	res, err := suite.repo.DeleteEventByID(int32(lastItem.ID))

	suite.True(res)
	suite.Nil(err)
	suite.NotNil(res)

}
