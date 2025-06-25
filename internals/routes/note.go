package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/shubGupta10/shared-space-server/internals/handlers"
	"github.com/shubGupta10/shared-space-server/internals/middleware"
)

func NoteRoutes(app *fiber.App) {
	note := app.Group("/note", middleware.ProtectedRoute)

	note.Post("/create-note", handlers.CreateNote)
	note.Delete("/delete-note", handlers.DeleteNote)
}
