package main

import (
	"data_service/database"
	"data_service/router"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func CreateApp() *fiber.App {
	var app *fiber.App = fiber.New()
	setupLoger(app)
	return app
}

func SetupRoutes(app *fiber.App, apiPreffix string, db *database.DataBase) fiber.Router {
	var api fiber.Router = app.Group(apiPreffix)
	router.SetupRoutes(api, db)
	return api
}

func createLogger() fiber.Handler {
	const loggerFormat = "[${ip}]:${port} ${status} - ${method} ${path}?${queryParams}\n"
	var loggerConfig logger.Config = logger.Config{
		Format: loggerFormat,
	}
	return logger.New(loggerConfig)
}

func setupLoger(app *fiber.App) {
	var logger fiber.Handler = createLogger()
	app.Use(logger)
}
