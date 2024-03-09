package commands

import (
	"github.com/omerbeden/event-mate/backend/tatooine/modules/validation/app/adapters/mernis"
)

type ValidateMernisCommand struct {
}

func NewValidateMernisCommand() *ValidateMernisCommand {
	return &ValidateMernisCommand{}
}

func (cmd *ValidateMernisCommand) ValidateMernis(nationalId, name, lastname string, birthYear int) (bool, error) {
	adapter := mernis.NewMernisAdapter(nationalId, name, lastname, birthYear)
	return adapter.Validate()
}
