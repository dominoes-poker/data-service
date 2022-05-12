package gamerRoutes

import (
	gameHandler "data_service/internal/handlers/game"

	fiber "github.com/gofiber/fiber/v2"
)

func SetupGameRoutes(router fiber.Router) {
	gamer := router.Group("/game")

	// Create a Game
	gamer.Post("/", gameHandler.CreateGame)
	// Read all Games
	gamer.Get("/", gameHandler.GetAllGames)
	// Read the Game with specified id
	gamer.Get("/:id", gameHandler.GetGame)
	// Add gamers to the game
	gamer.Post("/:id/add-gamers", gameHandler.AddGamersToGame)

}
