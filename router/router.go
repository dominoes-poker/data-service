package router

import (
	"data_service/database"
	gameRouter "data_service/router/v1/game"
	playerRoutes "data_service/router/v1/player"

	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(rootRouter fiber.Router, db *database.DataBase) {
	gameRootRouter := rootRouter.Group("/game")
	gameRouter.Setup(gameRootRouter, db)

	playerRootRouter := rootRouter.Group("/player")
	playerRoutes.Setup(playerRootRouter, db)
}
