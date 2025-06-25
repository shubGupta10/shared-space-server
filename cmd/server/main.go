package main

import (
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/shubGupta10/shared-space-server/internals/config"
)

func main() {
	app := fiber.New()

	//connect to the database
	config.ConnectToDatabase()

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
