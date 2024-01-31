package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/omerbeden/event-mate/backend/tatooine/cmd/api/handlers"
	"github.com/omerbeden/event-mate/backend/tatooine/modules/profile/app/entrypoints"
)

func ProfileRouter(app fiber.Router, service entrypoints.UserService) {
	app.Post("/profiles", handlers.CreateUserProfile(service))
	app.Get("/profiles/currentUser/:externalId", handlers.GetCurrentUserProfile(service))
	app.Patch("/profiles/currentUser/:externalId/profile-image", handlers.UpdateProfileImageUrl(service))
}
