package userHandler

import (
	"fmt"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/iqballm09/ezeats-backend-go/database"
	"github.com/iqballm09/ezeats-backend-go/internal/models"
)

func GetUsers(c *fiber.Ctx) error {
	db := database.DB
	var users []models.User

	// find all users
	db.Find(&users)

	return c.Status(200).JSON(fiber.Map{"status": "success", "data": users})
}

func GetUser(c *fiber.Ctx) error {
	db := database.DB
	var user models.User
	// get user id
	id1, _ := strconv.Atoi(c.Params("userId"))
	id := uint64(id1)
	db.Find(&user, "id = ?", id)

	if user.ID == 0 {
		return c.Status(404).JSON(fiber.Map{"status": "failed", "message": fmt.Sprintf("User by ID = %d not found!", id), "data": nil})
	}
	return c.Status(200).JSON(fiber.Map{"status": "success", "data": user})
}
