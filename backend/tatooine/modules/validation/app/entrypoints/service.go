package entrypoints

import (
	"github.com/omerbeden/event-mate/backend/tatooine/modules/validation/app/domain/commands"
)

type ValidationService struct{}

func (service *ValidationService) ValidateIdentity() (string, error) {
	cmd := commands.VeriffGenerateSessionCommand{}

	return cmd.Handle()
}
