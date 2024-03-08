package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/omerbeden/event-mate/backend/tatooine/modules/validation/app/entrypoints"
)

func ValidateIdentity(service entrypoints.ValidationService) fiber.Handler {
	return func(c *fiber.Ctx) error {

		service.ValidateIdentity()
		return nil
	}
}
