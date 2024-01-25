package restoHandler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/iqballm09/ezeats-backend-go/database"
	"github.com/iqballm09/ezeats-backend-go/internal/models"
)

func CreateResto(c *fiber.Ctx) error {
	var db = database.DB
	resto := new(models.Resto)

	err := c.BodyParser(resto)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"status": "failed", "message": "Check your input data", "data": err})
	}

	// create resto
	err = db.Create(&resto).Error
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "failed", "message": "Failed to create resto", "data": err})
	}
	return c.Status(201).JSON(fiber.Map{"status": "success", "message": "Resto has been created!", "data": err})
}
