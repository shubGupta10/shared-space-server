package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/shubGupta10/shared-space-server/internals/handlers"
	"github.com/shubGupta10/shared-space-server/internals/middleware"
)

func SpaceRoutes(app *fiber.App) {
	space := app.Group("/space", middleware.ProtectedRoute)

	space.Post("/create", handlers.CreateSpace)
	space.Get("/fetch-spaces", handlers.FetchSpace)
	space.Delete("/delete", handlers.DeleteSpace)
}
