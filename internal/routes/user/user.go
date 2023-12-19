package userRoutes

import (
	"github.com/gofiber/fiber/v2"
	userHandler "github.com/iqballm09/ezeats-backend-go/internal/handler/user"
)

func SetupUserRoutes(router fiber.Router) {
	user := router.Group("/user")
	// create a user
	user.Post("/", userHandler.CreateUser)
	// read all user
	user.Get("/", userHandler.GetUsers)
	// read user by id
	user.Get("/:userId", userHandler.GetUser)
	// update user
	user.Put("/:userId", userHandler.UpdateUser)
	// delete user
	user.Delete("/:userId", userHandler.DeleteUser)
}
