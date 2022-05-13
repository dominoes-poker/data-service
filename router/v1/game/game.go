package gamerRoutes

import (
	"data_service/database"
	gameHandler "data_service/handlers/game"

	fiber "github.com/gofiber/fiber/v2"
)

func Setup(rootRouter fiber.Router, db *database.DataBase) {
	handler := gameHandler.New(db)

	// Read all Games
	rootRouter.Get("/", func(ctx *fiber.Ctx) error {
		return handler.GetAll(ctx)
	})

	// Read the Game with specified id
	rootRouter.Get("/:id", func(ctx *fiber.Ctx) error {
		return handler.GetOne(ctx)
	})

	// Create a Game
	rootRouter.Post("/", func(ctx *fiber.Ctx) error {
		return handler.Create(ctx)
	})

	// Add gamers to the game
	rootRouter.Post("/:id/add-gamers", func(ctx *fiber.Ctx) error {
		return handler.AddGamersToGame(ctx)
	})
}
