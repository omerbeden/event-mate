package auth

import (
	"errors"
	"fmt"
	"strings"

	firebase "firebase.google.com/go"
	"github.com/gofiber/fiber/v2"
	"github.com/omerbeden/event-mate/backend/tatooine/cmd/api/presenter"
	"github.com/omerbeden/event-mate/backend/tatooine/modules/validation/app/domain/derrors"
	"github.com/omerbeden/event-mate/backend/tatooine/modules/validation/app/entrypoints"
)

func New(firebaseApp *firebase.App) fiber.Handler {
	return func(c *fiber.Ctx) error {
		service := entrypoints.NewValidationService(nil, firebaseApp)
		authHeader := c.Get("Authorization")
		fmt.Printf("authHeader: %v\n", authHeader)
		if authHeader != "" {
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
			return c.Next()
		}

		return c.Status(fiber.StatusUnauthorized).JSON(presenter.BaseResponse{
			APIVersion: presenter.APIVersion,
			Data:       nil,
			Error:      "Authorization header not found",
		})

	}
}
