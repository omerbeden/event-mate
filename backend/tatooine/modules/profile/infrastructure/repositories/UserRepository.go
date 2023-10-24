package repositories

import (
	"github.com/omerbeden/event-mate/backend/tatooine/modules/profile/core"
)

type UserRepositoryImpl struct{}

// DeleteUserById implements interfaces.UserRepository
func (*UserRepositoryImpl) DeleteUserById(id uint) {
	panic("unimplemented")
}

func NewUserRepo() *UserRepositoryImpl {
	return &UserRepositoryImpl{}
}

func (repo *UserRepositoryImpl) GetUsers() ([]core.UserProfile, error) {
	//db := database.NewConnPG()
	var users []core.UserProfile
	// if err := db.Find(&users).Error; err != nil {
	// 	return users, err
	// }
	return users, nil
}

func (repo *UserRepositoryImpl) GetUserById(id uint) (core.UserProfile, error) {
	// db := database.NewConnPG()
	var user core.UserProfile
	// if err := db.Preload("Adress").First(&user, id).Error; err != nil {
	// 	return user, gorm.ErrRecordNotFound
	// }

	return user, nil
}

func (repo *UserRepositoryImpl) InsertUser(user core.UserProfile) error {
	// db := database.NewConnPG()
	// if err := db.Create(user).Error; err != nil {
	// 	db.Logger.Error(nil, "Error occurred while inserting User")
	// 	return err
	// }

	return nil
}

func (repo *UserRepositoryImpl) UpdateUser(usertoUpdate *core.UserProfile) error {
	// if usertoUpdate.ID == 0 {
	// 	return errors.New("ID has not been set")
	// }

	// if _, err := repo.GetUserById(usertoUpdate.ID); err != nil {
	// 	return err
	// }

	// db := database.NewConnPG()
	// db.Save(usertoUpdate)
	return nil
}
