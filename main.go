package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/iqballm09/ezeats-backend-go/database"
)

func main() {
	// define fiber app
	app := fiber.New()

	// establish database connection
	database.ConnectDB()

	// define root route
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	app.Listen(":3000")
}
