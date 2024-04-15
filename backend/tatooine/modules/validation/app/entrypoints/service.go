package entrypoints

import (
	"context"

	firebase "firebase.google.com/go"
	"github.com/omerbeden/event-mate/backend/tatooine/modules/validation/app/domain/commands"
	"go.uber.org/zap"
)

type ValidationService struct {
	Logger      *zap.SugaredLogger
	firebaseApp *firebase.App
}

func NewValidationService(logger *zap.SugaredLogger, app *firebase.App) *ValidationService {
	return &ValidationService{
		Logger:      logger,
		firebaseApp: app,
	}
}

func (service *ValidationService) ValidateMernis(nationalId, name, lastname string, birthYear int) (bool, error) {
	cmd := commands.NewValidateMernisCommand()

	return cmd.ValidateMernis(nationalId, name, lastname, birthYear)

}

func (service *ValidationService) VerifyFirebaseToken(token string) (string, error) {
	cmd := commands.NewAuthCommand(service.firebaseApp, token)
	return cmd.Handle(context.Background())
}
