package entrypoints

import (
	"github.com/omerbeden/event-mate/backend/tatooine/modules/validation/app/domain/commands"
	"go.uber.org/zap"
)

type ValidationService struct {
	Logger *zap.SugaredLogger
}

func NewValidationService(logger *zap.SugaredLogger) *ValidationService {
	return &ValidationService{
		Logger: logger,
	}
}

func (service *ValidationService) ValidateMernis(nationalId, name, lastname string, birthYear int) (bool, error) {
	cmd := commands.NewValidateMernisCommand()

	return cmd.ValidateMernis(nationalId, name, lastname, birthYear)

}
