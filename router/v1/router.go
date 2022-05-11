package router

import (
	gameRoutes "data_service/internal/routes/v1/game"
	gamerRoutes "data_service/internal/routes/v1/gamer"

	fiber "github.com/gofiber/fiber/v2"
	logger "github.com/gofiber/fiber/v2/middleware/logger"
)

func SetupRoutes(app *fiber.App) {
	api := app.Group("/api", logger.New())
	gamerRoutes.SetupGamerRoutes(api)
	gameRoutes.SetupGameRoutes(api)
}
