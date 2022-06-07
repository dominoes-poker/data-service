package router

import (
	"data_service/database"
	gameRouter "data_service/router/v1/game"
	gamerRoutes "data_service/router/v1/gamer"

	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(rootRouter fiber.Router, db *database.DataBase) {
	gameRootRouter := rootRouter.Group("/game")
	gameRouter.Setup(gameRootRouter, db)

	gamerRootRouter := rootRouter.Group("/gamer")
	gamerRoutes.Setup(gamerRootRouter, db)
}
