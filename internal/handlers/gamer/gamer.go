package gamerHandler

import (
	"data_service/database"
	"data_service/internal/models"
	"github.com/gofiber/fiber/v2"
)

func GetGamers(c *fiber.Ctx) error {
	db := database.GetInstance()
	var gamers []models.Gamer

	// find all gamers in the database
	db.DB.Find(&gamers)

	// Else return gamers
	return c.JSON(fiber.Map{"status": "success", "message": "Gamers Found", "data": gamers})
}

func CreateGamer(c *fiber.Ctx) error {
	db := database.GetInstance()
	gamer := new(models.Gamer)

	// Store the body in the note and return error if encountered
	err := c.BodyParser(gamer)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Review your input", "data": err})
	}

	// Create the Gamer and return error if encountered
	err = db.Create(&gamer).Error
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Could not create gamer", "data": err})
	}

	// Return the created note
	return c.JSON(fiber.Map{"status": "success", "message": "Created gamer", "data": gamer})
}
