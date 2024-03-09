package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/omerbeden/event-mate/backend/tatooine/cmd/api/handlers"
	"github.com/omerbeden/event-mate/backend/tatooine/modules/validation/app/entrypoints"
)

func ValidationRouter(app fiber.Router, service *entrypoints.ValidationService) {
	app.Post("/validations/validateIdentity", handlers.ValidateIdentity(service))

}
