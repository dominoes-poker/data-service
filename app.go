package main

import (
	"data_service/database"
	"data_service/router"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func CreateApp(db *database.DataBase) *fiber.App {
	// Start a new fiber app
	app := fiber.New()
	app.Use(logger.New(logger.Config{
		Format: "[${ip}]:${port} ${status} - ${method} ${path}?${queryParams}\n",
	}))
	api := app.Group("/api")
	router.SetupRoutes(api, db)
	return app
}
