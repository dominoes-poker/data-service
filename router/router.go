package router

import (
	gamerRoutes "data_service/internal/routes/v1"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func SetupRoutes(app *fiber.App) {
	api := app.Group("/api", logger.New())

	// Setup the Node Routes
	gamerRoutes.SetupGamerRoutes(api)
}
