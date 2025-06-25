package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/shubGupta10/shared-space-server/internals/handlers"
	"github.com/shubGupta10/shared-space-server/internals/middleware"
)

func AuthRoutes(app *fiber.App) {
	auth := app.Group("/auth")

	auth.Post("/register", handlers.Register)
	auth.Post("/login", handlers.Login)
	auth.Get("/get-profile", middleware.ProtectedRoute, handlers.GetProfile)
}
