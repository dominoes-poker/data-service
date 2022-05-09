package gamerRoutes

import (
	gamerHandler "data_service/internal/handlers/gamer"
	"github.com/gofiber/fiber/v2"
)

func SetupGamerRoutes(router fiber.Router) {
	note := router.Group("/gamer")
	// Create a Gamer
	note.Post("/", gamerHandler.CreateGamer)
	// Read all Gamers
	note.Get("/", gamerHandler.GetGamers)
	// Read one Gamer
	// note.Get("/:gamerId", func(c *fiber.Ctx) error {})
	// Update one Gamer
	// note.Put("/:gamerId", func(c *fiber.Ctx) error {})
	// Delete one Gamer
	// note.Delete("/:gamerId", func(c *fiber.Ctx) error {})
}

func SetupGameRoutes(router fiber.Router) {
	note := router.Group("/game")
	// Create a Gamer
	note.Post("/", gamerHandler.CreateGamer)
	// Read all Gamers
	note.Get("/", gamerHandler.GetGamers)
	// Read one Gamer
	// note.Get("/:gamerId", func(c *fiber.Ctx) error {})
	// Update one Gamer
	// note.Put("/:gamerId", func(c *fiber.Ctx) error {})
	// Delete one Gamer
	// note.Delete("/:gamerId", func(c *fiber.Ctx) error {})
}
