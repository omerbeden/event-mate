package commands

import (
	"github.com/omerbeden/event-mate/backend/tatooine/modules/profile/domain/model"
	"github.com/omerbeden/event-mate/backend/tatooine/modules/profile/domain/ports"
)

type CreateProfileCommand struct {
	Profile model.UserProfile
	repo    ports.UserProfileRepository
}

func (ccmd *CreateProfileCommand) Handle() (bool, error) {
	return ccmd.repo.InsertUser(&ccmd.Profile)
}
