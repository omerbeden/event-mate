package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/omerbeden/event-mate/backend/tatooine/cmd/api/handlers"
	"github.com/omerbeden/event-mate/backend/tatooine/modules/activity/app/entrypoints"
)

func ActivityRouter(app fiber.Router, service entrypoints.ActivityService) {
	app.Post("/activities", handlers.CreateActivity(service))
	app.Post("/activities/:activityId/participants", handlers.AddParticipant(service))
	app.Get("/activities/:activityId/participants", handlers.GetParticipants(service))
	app.Get("/activities/:city", handlers.GetActivitiesByLocation(service))
}
