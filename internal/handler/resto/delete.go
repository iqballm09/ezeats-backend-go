package restoHandler

import (
	"fmt"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/iqballm09/ezeats-backend-go/database"
	"github.com/iqballm09/ezeats-backend-go/internal/models"
)

func DeleteResto(c *fiber.Ctx) error {
	db := database.DB
	var resto models.Resto

	// get resto id
	id1, _ := strconv.Atoi(c.Params("restoId"))
	id := uint64(id1)
	db.Find(&resto, "id = ?", id)

	// check if resto exists
	if resto.ID == 0 {
		return c.Status(404).JSON(fiber.Map{"status": "failed", "message": fmt.Sprintf("Resto by ID = %d not found!", id), "data": nil})
	}
	result := db.Delete(&resto, "id = ?", id)

	if result.Error != nil {
		return c.Status(500).JSON(fiber.Map{"status": "failed", "message": fmt.Sprintf("Failed to delete resto with ID = %d has been deleted", id), "data": result.Error})
	}
	return c.Status(200).JSON(fiber.Map{"status": "success", "message": fmt.Sprintf("Resto with ID = %d has been deleted", id)})
}
