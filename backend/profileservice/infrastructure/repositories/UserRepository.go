package repositories

import (
	"errors"
	"fmt"

	"github.com/omerbeden/event-mate/backend/profileservice/core"
	"github.com/omerbeden/event-mate/backend/profileservice/infrastructure/database"
	"gorm.io/gorm"
)

//TODO: do implement
type UserRepositoryImpl struct{}

func NewUserRepo() *UserRepositoryImpl {
	return &UserRepositoryImpl{}
}

func (repo *UserRepositoryImpl) GetUsers() *[]core.User {
	dbConn := database.NewConnPG()
	var users []core.User
	dbConn.Find(&users)
	return &users
}

func (repo *UserRepositoryImpl) GetUserById(id uint) (core.User, error) {
	db := database.NewConnPG()
	var user core.User
	if result := db.Preload("Adress").First(&user, id); result.Error != nil {
		fmt.Printf("can not get user by id %d", id)
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		}
		return user, gorm.ErrRecordNotFound
	}

	return user, nil
}

func (repo *UserRepositoryImpl) InsertUser(user *core.User) {
	dbConn := database.NewConnPG()
	dbConn.Create(user)
}

func (repo *UserRepositoryImpl) UpdateUser(usertoUpdate *core.User) error {
	if usertoUpdate.ID == 0 {
		return errors.New("ID has not been set")
	}
	_, err := repo.GetUserById(usertoUpdate.ID)
	if err != nil {
		return err
	}

	db := database.NewConnPG()
	db.Save(usertoUpdate)
	return nil
}
