package restoHandler

import (
	"fmt"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/iqballm09/ezeats-backend-go/database"
	"github.com/iqballm09/ezeats-backend-go/internal/models"
)

func GetRestos(c *fiber.Ctx) error {
	db := database.DB
	var restos []models.Resto

	// find all restos
	db.Find(&restos)

	if len(restos) == 0 {
		return c.Status(404).JSON(fiber.Map{"status": "failed", "message": "Restos not found!", "data": nil})
	}
	return c.Status(200).JSON(fiber.Map{"status": "success", "message": "Restos found", "data": restos})
}

func GetResto(c *fiber.Ctx) error {
	db := database.DB
	var resto models.Resto

	// get resto by id
	id1, _ := strconv.Atoi(c.Params("restoId"))
	id := uint64(id1)
	db.Find(&resto, "id = ?", id)

	if resto.ID == 0 {
		return c.Status(404).JSON(fiber.Map{"status": "failed", "message": fmt.Sprintf("Resto by ID = %d not found!", id), "data": nil})
	}
	return c.Status(200).JSON(fiber.Map{"status": "success", "message": fmt.Sprintf("Resto by ID = %d found", id), "data": resto})
}
