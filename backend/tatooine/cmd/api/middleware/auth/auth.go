package auth

import (
	"errors"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/omerbeden/event-mate/backend/tatooine/cmd/api/presenter"
	"github.com/omerbeden/event-mate/backend/tatooine/modules/validation/app/domain/derrors"
	"github.com/omerbeden/event-mate/backend/tatooine/modules/validation/app/entrypoints"
)

func New() fiber.Handler {

	return func(c *fiber.Ctx) error {
		service := entrypoints.NewValidationService(nil, nil)
		authHeader := c.Get("Authorization")
		idToken := strings.Split(authHeader, " ")[1]

		token, err := service.VerifyFirebaseToken(idToken)
		if errors.Is(err, derrors.ErrFirebaseAuth) {
			return c.Status(fiber.StatusUnauthorized).JSON(presenter.BaseResponse{
				APIVersion: presenter.APIVersion,
				Data:       nil,
				Error:      err.Error(),
			})
		}

		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(presenter.BaseResponse{
				APIVersion: presenter.APIVersion,
				Data:       nil,
				Error:      presenter.UNKNOW_ERR,
			})
		}
		c.Context().SetUserValue("userToken", token)
		c.Next()
		return nil
	}
}
