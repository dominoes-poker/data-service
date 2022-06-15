package playerrRoutes

import (
	"data_service/database"
	playerHandler "data_service/handlers/player"
	"data_service/handlers/results"
	routerUtils "data_service/router/common"

	"github.com/gofiber/fiber/v2"
)

func Setup(rootRouter fiber.Router, db *database.DataBase) {
	handler := playerHandler.New(db)

	// Read all Players
	rootRouter.Get("/", func(ctx *fiber.Ctx) error {
		username := ctx.Query("username")
		identificator := ctx.Query("identificator")

		return handler.GetAll(ctx, username, identificator)
	})

	// Read the Player with specified id
	rootRouter.Get("/:playerId", func(ctx *fiber.Ctx) error {
		playerId, err := routerUtils.GetUintParam(ctx, "playerId")
		if err != nil {
			return results.BadRequestResult(ctx, err)
		}
		return handler.GetOne(&playerId, ctx)
	})

	// Create a Player
	rootRouter.Post("/", func(ctx *fiber.Ctx) error {
		return handler.Create(ctx)
	})
}
