package repositories

import (
	"github.com/omerbeden/event-mate/backend/tatooine/modules/profile/core"
	"github.com/omerbeden/event-mate/backend/tatooine/modules/profile/infrastructure/repositories"
)

type UserStatRepoService struct {
	Repo repositories.UserStatRepo
}

func (s *UserStatRepoService) GetUserStat(userId int) (core.UserProfileStat, error) {
	return s.Repo.GetUserStat(userId)
}
