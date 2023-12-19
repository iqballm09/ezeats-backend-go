package userHandler

import (
	"fmt"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/iqballm09/ezeats-backend-go/database"
	"github.com/iqballm09/ezeats-backend-go/internal/models"
)

type updateUser struct {
	Email     string
	Name      string
	Deskripsi string
	Password  string
	NoTelp    string
	Alamat    string
	UrlGambar string
}

func UpdateUser(c *fiber.Ctx) error {
	db := database.DB
	var user models.User

	// get user id
	id1, _ := strconv.Atoi(c.Params("userId"))
	id := uint64(id1)
	db.Find(&user, "id = ?", id)

	if user.ID == 0 {
		return c.Status(404).JSON(fiber.Map{"status": "failed", "message": fmt.Sprintf("User by ID = %d not found!", id), "data": nil})
	}

	// store data of update user
	var updateUserData updateUser
	err := c.BodyParser(&updateUserData)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "failed", "message": "Check your input!", "data": err})
	}

	// update the data
	user.Name = updateUserData.Name
	user.Email = updateUserData.Email
	user.Deskripsi = updateUserData.Deskripsi
	user.Alamat = updateUserData.Alamat
	user.Password = updateUserData.Password
	user.NoTelp = updateUserData.NoTelp
	user.UrlGambar = updateUserData.UrlGambar

	db.Save(&user)
	return c.Status(200).JSON(fiber.Map{"status": "success", "message": "User has been updated!", "data": user})
}
