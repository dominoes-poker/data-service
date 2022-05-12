package gamerHandler

import (
	"data_service/database"
	"data_service/models"

	"github.com/gofiber/fiber/v2"
)

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

func GetAllGames(ctx *fiber.Ctx) error {
	db := database.GetInstance()
	var games []models.Game

	// find all games in the database
	result := db.DB.Preload("Gamers").Find(&games)
	if result.Error != nil {
		return ctx.Status(500).JSON(fiber.Map{"status": "error", "message": "Cannot make a select opration", "data": result.Error})
	}
	// Else return games
	return ctx.JSON(fiber.Map{"status": "success", "message": "Games Found", "data": games})
}

func GetGame(ctx *fiber.Ctx) error {
	db := database.GetInstance()
	var game models.Game

	gameId := ctx.Params("gameId")

	result := db.DB.First(&game, gameId)
	if result.Error != nil {
		return ctx.Status(500).JSON(fiber.Map{"status": "error", "message": "Cannot make a select opration", "data": result.Error})
	}

	// Else return games
	return ctx.JSON(fiber.Map{"status": "success", "message": "Game Found", "data": game})
}

func AddGamersToGame(ctx *fiber.Ctx) error {
	payload := struct {
		GamerIds []int `json:"gamerIds"`
	}{}

	if err := ctx.BodyParser(&payload); err != nil {
		return ctx.Status(500).JSON(fiber.Map{"status": "error", "message": "Cannot parse the body", "data": err})
	}

	var game models.Game
	gamers := make([]models.Gamer, len(payload.GamerIds))
	gameId := ctx.Params("gameId")
	db := database.GetInstance()

	if result := db.DB.First(&game, gameId); result.Error != nil {
		return ctx.Status(500).JSON(fiber.Map{"status": "error", "message": "Cannot make a select opration", "data": result.Error})
	}

	for index, gamerId := range payload.GamerIds {
		gamers[index].ID = uint(gamerId)
	}

	assosiation := db.DB.Model(&game).Association("Gamers")
	assosiation.Append(gamers)

	return ctx.JSON(fiber.Map{"status": "success", "message": "Created game"})
}
