package main

import (
	"data_service/database"
	"data_service/router"

	"github.com/gofiber/fiber/v2"
)

func CreateApp(db *database.DataBase) *fiber.App {
	// Start a new fiber app
	app := fiber.New()
	api := app.Group("/api")
	router.SetupRoutes(api, db)
	return app
}
