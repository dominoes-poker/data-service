package main

import (
	"data_service/router"

	"github.com/gofiber/fiber/v2"
)

func CreateApp() *fiber.App {
	// Start a new fiber app
	app := fiber.New()
	router.SetupRoutes(app)
	return app
}
