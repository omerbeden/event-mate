package entrypoints

import (
	"github.com/omerbeden/event-mate/backend/tatooine/modules/validation/app/domain/commands"
)

type ValidationService struct{}

func (service *ValidationService) ValidateMernis(nationalId, name, lastname string, birthYear int) (bool, error) {
	return commands.ValidateMernis(nationalId, name, lastname, birthYear)

}
