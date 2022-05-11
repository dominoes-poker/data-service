package main

import (
	"data_service/database"
	"data_service/internal/models"
	"data_service/router/v1"

	"github.com/gofiber/fiber/v2"
)

func setup_database(db_url string) *database.DataBase {
	db := database.GetInstance()
	db.Connect(db_url)
	db.InitializeTables(models.Gamer{}, models.Game{})
	return db
}

func create_app() *fiber.App {
	// Start a new fiber app
	app := fiber.New()
	router.SetupRoutes(app)
	return app
}
