package main

import (
	"data_service/config"
	"data_service/database"
	"data_service/internal/models"
	"data_service/router"

	"github.com/gofiber/fiber/v2"
)

func main() {
	// Start a new fiber app
	app := fiber.New()

	db_url := config.Config("DB_URL")
	db := database.GetInstance()
	db.Connect(db_url)
	db.InitializeTables(models.Gamer{})

	router.SetupRoutes(app)

	// Listen on PORT 300

	app.Listen(":3000")
}
