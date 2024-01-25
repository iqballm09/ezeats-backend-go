package restoHandler

import (
	"fmt"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/iqballm09/ezeats-backend-go/database"
	"github.com/iqballm09/ezeats-backend-go/internal/models"
)

type updateResto struct {
	Nama      string `gorm:"not null"`
	HargaMin  uint64 `gorm:"not null"`
	HargaMaks uint64 `gorm:"not null"`
	JamBuka   uint16 `gorm:"not null"`
	JamTutup  uint16 `gorm:"not null"`
}

func UpdateResto(c *fiber.Ctx) error {
	db := database.DB
	var resto models.Resto

	// get resto id
	id1, _ := strconv.Atoi(c.Params("restoId"))
	id := uint64(id1)
	db.Find(&resto, "id = ?", id)

	if resto.ID == 0 {
		return c.Status(404).JSON(fiber.Map{"status": "failed", "message": fmt.Sprintf("Resto by ID = %d not found!", id), "data": nil})
	}
	var updateRestoData updateResto
	err := c.BodyParser(&updateRestoData)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"status": "failed", "message": "Check your input!", "data": nil})
	}

	// update the data
	resto.Nama = updateRestoData.Nama
	resto.HargaMaks = updateRestoData.HargaMaks
	resto.HargaMin = updateRestoData.HargaMin
	resto.JamBuka = updateRestoData.JamBuka
	resto.JamTutup = updateRestoData.JamTutup

	db.Save(&resto)
	return c.Status(200).JSON(fiber.Map{"status": "success", "message": "Resto has been updated", "data": resto})
}
