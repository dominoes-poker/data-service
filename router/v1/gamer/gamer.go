package gamerRoutes

import (
	gamerHandler "data_service/handlers/gamer"

	fiber "github.com/gofiber/fiber/v2"
)

func SetupGamerRoutes(router fiber.Router) {
	gamer := router.Group("/gamer")

	// Create a Gamer
	gamer.Post("/", gamerHandler.CreateGamer)
	// Read all Gamers
	gamer.Get("/", gamerHandler.GetGamers)
}
