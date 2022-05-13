package gamerRoutes

import (
	"data_service/database"
	gamerHandler "data_service/handlers/gamer"

	fiber "github.com/gofiber/fiber/v2"
)

func Setup(rootRouter fiber.Router, db *database.DataBase) {
	handler := gamerHandler.New(db)

	// Read all Gamers
	rootRouter.Get("/", func(ctx *fiber.Ctx) error {
		return handler.GetAll(ctx)
	})

	// Read the Gamer with specified id
	rootRouter.Get("/:id", func(ctx *fiber.Ctx) error {
		return handler.GetOne(ctx)
	})

	// Create a Gamer
	rootRouter.Post("/", func(ctx *fiber.Ctx) error {
		return handler.Create(ctx)
	})
}
