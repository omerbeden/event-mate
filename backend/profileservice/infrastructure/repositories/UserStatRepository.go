package repositories

import (
	"github.com/omerbeden/event-mate/backend/profileservice/core"
	"github.com/omerbeden/event-mate/backend/profileservice/infrastructure/database"
	"gorm.io/gorm"
)

type UserStatRepo struct{}

func (r *UserStatRepo) GetUserStat(userId int) (core.UserProfileStat, error) {
	db := database.NewConnPG()
	var userStat core.UserProfileStat

	if err := db.First(&userStat, userId).Error; err != nil {
		return userStat, gorm.ErrRecordNotFound
	}

	return userStat, nil
}
