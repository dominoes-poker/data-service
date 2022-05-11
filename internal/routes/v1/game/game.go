package gamerRoutes

import (
	gameHandler "data_service/internal/handlers/game"

	fiber "github.com/gofiber/fiber/v2"
)

func SetupGameRoutes(router fiber.Router) {
	gamer := router.Group("/game")

	// Create a Gamer
	gamer.Post("/", gameHandler.CreateGame)
	// Read all Gamers
	gamer.Get("/", gameHandler.GetGames)
}
