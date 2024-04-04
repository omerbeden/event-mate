package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/omerbeden/event-mate/backend/tatooine/cmd/api/handlers"
	"github.com/omerbeden/event-mate/backend/tatooine/modules/profile/app/entrypoints"
)

func ProfileRouter(app fiber.Router, service entrypoints.UserService) {
	app.Post("/profiles", handlers.CreateUserProfile(service))
	app.Get("/profiles/:id", handlers.GetUserProfile(service))
	app.Get("/profiles/currentUser/:externalId", handlers.GetCurrentUserProfile(service))
	app.Get("/profiles/currentUser/badges/:externalId", handlers.GetProfileBadges(service))
	app.Patch("/profiles/currentUser/:externalId/profile-image", handlers.UpdateProfileImageUrl(service))
	app.Patch("/profiles/currentUser/:externalId/verification", handlers.UpdateProfileVerification(service))
}
