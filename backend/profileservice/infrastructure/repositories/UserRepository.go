package repositories

import (
	"github.com/omerbeden/event-mate/backend/profileservice/core"
	"github.com/omerbeden/event-mate/backend/profileservice/infrastructure/database"
)

//TODO: do implement
type UserRepositoryImpl struct{}

func NewUserRepo() *UserRepositoryImpl {
	return &UserRepositoryImpl{}
}

func (repo *UserRepositoryImpl) GetUsers() []core.User {
	dbConn := database.NewConnPG()
	var users []core.User
	dbConn.Find(&users)
	return users
}

func (repo *UserRepositoryImpl) InsertUser(user *core.User) {
	dbConn := database.NewConnPG()

	dbConn.Create(user)

}
