package commands

import (
	"github.com/omerbeden/event-mate/backend/tatooine/modules/profile/app/domain/model"
	"github.com/omerbeden/event-mate/backend/tatooine/modules/profile/app/domain/ports"
)

type CreateProfileCommand struct {
	Profile model.UserProfile
	Repo    ports.UserProfileRepository
}

func (ccmd *CreateProfileCommand) Handle() (bool, error) {
	return ccmd.Repo.InsertUser(&ccmd.Profile)
}
