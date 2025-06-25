package main

import (
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/shubGupta10/shared-space-server/internals/config"
	"github.com/shubGupta10/shared-space-server/internals/routes"
)

func main() {
	app := fiber.New()

	//validate frontend url
	frontendURL := os.Getenv("FRONTEND_URL")
	if frontendURL == "" {
		log.Fatal("FRONTEND_URL not set in env")
	}

	//we can do this AllowOrigins: "*" if not working for native apps

	//cors configuration
	app.Use(cors.New(cors.Config{
		AllowOrigins: frontendURL,
		AllowMethods: "GET,POST,HEAD,PUT,DELETE,PATCH",
		AllowHeaders: "Origin,Content-Type,Accept, Authorization",
	}))

	//connect to the database
	config.ConnectToDatabase()

	//connect to redis
	config.ConnectToRedis()

	// routes
	routes.AuthRoutes(app)
	routes.SpaceRoutes(app)
	routes.NoteRoutes(app)

	//basic health route
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}
	// Start the server on the specified port
	if err := app.Listen(":" + port); err != nil {
		log.Fatal(err)
	}
}
