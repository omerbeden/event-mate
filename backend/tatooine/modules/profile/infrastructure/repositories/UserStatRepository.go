package repositories

import (
	"github.com/omerbeden/event-mate/backend/tatooine/modules/profile/application/interfaces"
	"github.com/omerbeden/event-mate/backend/tatooine/modules/profile/core"
)

type UserStatRepo struct{}

var _ interfaces.UserStatRepository = (*UserStatRepo)(nil)

func (r *UserStatRepo) GetUserStat(userId int) (core.UserProfileStat, error) {
	// db := database.NewConnPG()
	var userStat core.UserProfileStat

	// if err := db.First(&userStat, userId).Error; err != nil {
	// 	return userStat, gorm.ErrRecordNotFound
	// }

	return userStat, nil
}
