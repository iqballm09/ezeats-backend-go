package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/iqballm09/ezeats-backend-go/database"
	"github.com/iqballm09/ezeats-backend-go/router"
)

func main() {
	// define fiber app
	app := fiber.New()

	// establish database connection
	database.ConnectDB()

	// add routes to app
	router.SetupRouter(app)

	app.Listen(":3003")
}
