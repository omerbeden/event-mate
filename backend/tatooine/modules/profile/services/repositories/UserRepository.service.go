package repositories

import (
	"github.com/omerbeden/event-mate/backend/tatooine/modules/profile/core"
	"github.com/omerbeden/event-mate/backend/tatooine/modules/profile/infrastructure/repositories"
)

type UserRepoService struct {
	Repo repositories.UserRepositoryImpl
}

func (s *UserRepoService) GetUsers() ([]core.UserProfile, error) {
	return s.Repo.GetUsers()
}

func (s *UserRepoService) GetUserByID(ID uint) (core.UserProfile, error) {
	return s.Repo.GetUserById(ID)
}

func (s *UserRepoService) InsertUser(user core.UserProfile) error {
	return s.Repo.InsertUser(user)
}

func (s *UserRepoService) UpdateUser(userToUpdate *core.UserProfile) error {
	return s.Repo.UpdateUser(userToUpdate)
}

func (s *UserRepoService) DeleteUserById(ID uint) {
	s.Repo.DeleteUserById(ID)
}
