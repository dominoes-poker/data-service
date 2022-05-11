package gamerHandler

import (
	"data_service/database"
	"data_service/internal/models"

	"github.com/gofiber/fiber/v2"
)

func GetGames(c *fiber.Ctx) error {
	db := database.GetInstance()
	var games []models.Game

	// find all games in the database
	db.DB.Find(&games).Association("gamers")

	// Else return games
	return c.JSON(fiber.Map{"status": "success", "message": "Games Found", "data": games})
}

func CreateGame(c *fiber.Ctx) error {
	db := database.GetInstance()
	game := new(models.Game)

	// Store the body in the Game and return error if encountered
	if err := c.BodyParser(game); err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Review your input", "data": err})
	}

	// Create the Game and return error if encountered
	if err := db.Create(&game).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Could not create game", "data": err})
	}

	// Return the created note
	return c.JSON(fiber.Map{"status": "success", "message": "Created game", "data": game})
}
