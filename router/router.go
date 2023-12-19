package router

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	userRoutes "github.com/iqballm09/ezeats-backend-go/internal/routes/user"
)

func SetupRouter(app *fiber.App) {
	api := app.Group("/api", logger.New())

	// setup routes
	userRoutes.SetupUserRoutes(api)
}
