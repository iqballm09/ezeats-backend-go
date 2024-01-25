package restoRoutes

import (
	"github.com/gofiber/fiber/v2"
	restoHandler "github.com/iqballm09/ezeats-backend-go/internal/handler/resto"
)

func SetupRestoRoutes(router fiber.Router) {
	resto := router.Group("/resto")
	// create resto
	resto.Post("/", restoHandler.CreateResto)
	// read all resto
	resto.Get("/", restoHandler.GetRestos)
	// read resto by id
	resto.Get("/:restoId", restoHandler.GetResto)
	// update resto
	resto.Put("/:restoId", restoHandler.UpdateResto)
	// delete resto
	resto.Delete("/:restoId", restoHandler.DeleteResto)
}
