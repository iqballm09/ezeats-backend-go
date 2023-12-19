package userHandler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/iqballm09/ezeats-backend-go/database"
	"github.com/iqballm09/ezeats-backend-go/internal/models"
)

func CreateUser(c *fiber.Ctx) error {
	var db = database.DB
	user := new(models.User)

	err := c.BodyParser(user)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "failed", "message": "Check your input", "data": err})
	}

	// create user
	err = db.Create(&user).Error
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "failed", "message": "Failed to create user", "data": err})
	}

	// return create user
	return c.Status(201).JSON(fiber.Map{"status": "success", "message": "User has been created!", "data": user})
}
