package gamerRoutes

import (
	"data_service/database"
	"data_service/handlers/common"
	gameHandler "data_service/handlers/game"
	"data_service/handlers/results"

	fiber "github.com/gofiber/fiber/v2"
)

func Setup(rootRouter fiber.Router, db *database.DataBase) {
	handler := gameHandler.New(db)

	// Get all Games
	rootRouter.Get("/", func(ctx *fiber.Ctx) error {
		return handler.GetAll(ctx)
	})

	// Read the Game with specified id
	rootRouter.Get("/:gameId", func(ctx *fiber.Ctx) error {
		gameId, err := common.GetUintParam(ctx, "gameId")
		if err != nil {
			return results.BadRequestResult(ctx, err)
		}
		return handler.GetOne(gameId, ctx)
	})

	// Create a Game
	rootRouter.Post("/", func(ctx *fiber.Ctx) error {
		return handler.Create(ctx)
	})

	// Add gamers to the game
	rootRouter.Post("/:gameId/add-gamers", func(ctx *fiber.Ctx) error {
		gameId, err := common.GetUintParam(ctx, "gameId")
		if err != nil {
			return results.BadRequestResult(ctx, err)
		}
		return handler.AddGamersToGame(gameId, ctx)
	})

	// Start a new round
	rootRouter.Post("/:gameId/new-round", func(context *fiber.Ctx) error {
		gameId, err := common.GetUintParam(context, "gameId")
		if err != nil {
			return results.BadRequestResult(context, err)
		}
		return handler.StartRound(gameId, context)
	})

	// User make a bet
	rootRouter.Post("/:gameId/round/:roundNumber/make-bet", func(context *fiber.Ctx) error {

		gameId, err := common.GetUintParam(context, "gameId")
		if err != nil {
			return results.BadRequestResult(context, err)
		}

		roundNumber, err := common.GetUintParam(context, "roundNumber")
		if err != nil {
			return results.BadRequestResult(context, err)
		}

		return handler.MakeBet(gameId, roundNumber, context)
	})

}
