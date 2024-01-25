package userHandler

import (
	"fmt"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/iqballm09/ezeats-backend-go/database"
	"github.com/iqballm09/ezeats-backend-go/internal/models"
)

func DeleteUser(c *fiber.Ctx) error {
	db := database.DB
	var user models.User

	// get user id
	id1, _ := strconv.Atoi(c.Params("userId"))
	id := uint64(id1)
	db.Find(&user, "id = ?", id)

	// check if user exists
	if user.ID == 0 {
		return c.Status(404).JSON(fiber.Map{"status": "failed", "message": fmt.Sprintf("User by ID = %d not found!", id), "data": nil})
	}

	result := db.Delete(&user, "id = ?", id)
	if result.Error != nil {
		return c.Status(400).JSON(fiber.Map{"status": "failed", "message": fmt.Sprintf("Failed to delete user by ID = %d!", id), "data": result.Error})
	}
	return c.Status(200).JSON(fiber.Map{"status": "success", "message": fmt.Sprintf("User with ID = %d has been deleted", id)})
}
