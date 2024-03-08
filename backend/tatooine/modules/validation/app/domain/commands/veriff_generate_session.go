package commands

import (
	"github.com/omerbeden/event-mate/backend/tatooine/modules/validation/app/adapters/veriff"
)

type VeriffGenerateSessionCommand struct{}

func (cmd *VeriffGenerateSessionCommand) Handle() (string, error) {

	verifAdapter := veriff.NewVeriffAdapter()
	return verifAdapter.GenerateSession()
}
