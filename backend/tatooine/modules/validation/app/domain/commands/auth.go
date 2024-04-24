package commands

import (
	"context"
	"errors"
	"fmt"

	firebase "firebase.google.com/go"
	"github.com/omerbeden/event-mate/backend/tatooine/modules/validation/app/domain/derrors"
)

type authCommand struct {
	app     *firebase.App
	idToken string
}

func NewAuthCommand(app *firebase.App, idToken string) *authCommand {
	return &authCommand{
		idToken: idToken,
		app:     app,
	}
}

func (cmd authCommand) Handle(ctx context.Context) (string, error) {

	client, err := cmd.app.Auth(ctx)
	if err != nil {
		return "", fmt.Errorf("error getting Auth client: %w", err)
	}

	token, err := client.VerifyIDToken(ctx, cmd.idToken)
	if err != nil {
		return "", errors.Join(derrors.ErrFirebaseAuth, fmt.Errorf("error verifying ID token: %w", err))
	}

	return token.UID, nil
}
