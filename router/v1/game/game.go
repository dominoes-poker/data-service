package playerRoute

import (
	"data_service/database"
	gameHandler "data_service/handlers/game"
	"data_service/handlers/results"
	routerUtils "data_service/router/common"

	"github.com/gofiber/fiber/v2"
)

func Setup(rootRouter fiber.Router, db *database.DataBase) {
	handler := gameHandler.New(db)

	// Get all Games
	rootRouter.Get("/", func(ctx *fiber.Ctx) error {
		return handler.GetAll(ctx)
	})

	// Read the Game with specified id
	rootRouter.Get("/:gameId", func(ctx *fiber.Ctx) error {
		gameId, err := routerUtils.GetUintParam(ctx, "gameId")
		if err != nil {
			return results.BadRequestResult(ctx, err)
		}
		return handler.GetOne(gameId, ctx)
	})

	// Create a Game
	rootRouter.Post("/", func(ctx *fiber.Ctx) error {
		return handler.Create(ctx)
	})

	// Start a new round
	rootRouter.Post("/:gameId/new-round", func(context *fiber.Ctx) error {
		gameId, err := routerUtils.GetUintParam(context, "gameId")
		if err != nil {
			return results.BadRequestResult(context, err)
		}
		return handler.StartRound(gameId, context)
	})

	// User make a bet
	rootRouter.Post("/:gameId/round/:roundNumber/bet", func(context *fiber.Ctx) error {

		gameId, err := routerUtils.GetUintParam(context, "gameId")
		if err != nil {
			return results.BadRequestResult(context, err)
		}

		roundNumber, err := routerUtils.GetUintParam(context, "roundNumber")
		if err != nil {
			return results.BadRequestResult(context, err)
		}

		return handler.MakeBet(gameId, roundNumber, context)
	})

	// Set a bribe
	rootRouter.Post("/:gameId/round/:roundNumber/bribe", func(context *fiber.Ctx) error {

		gameId, err := routerUtils.GetUintParam(context, "gameId")
		if err != nil {
			return results.BadRequestResult(context, err)
		}

		roundNumber, err := routerUtils.GetUintParam(context, "roundNumber")
		if err != nil {
			return results.BadRequestResult(context, err)
		}

		return handler.SetBribe(gameId, roundNumber, context)
	})

}
