package gamerHandler

import (
	"data_service/database"
	"data_service/handlers/results"
	"data_service/models"
	"sync"

	"github.com/gofiber/fiber/v2"
)

type GameHandler struct {
	db *database.DataBase
}

var instance *GameHandler
var gameRouterOnce sync.Once

func New(db *database.DataBase) *GameHandler {
	gameRouterOnce.Do(func() {
		instance = &GameHandler{db}
	})
	return instance
}

func (handler *GameHandler) Create(ctx *fiber.Ctx) error {
	game := new(models.Game)

	// Store the body in the Game and return error if encountered
	if err := ctx.BodyParser(game); err != nil {
		return results.BadRequestResult(ctx, "Bad request body", err)
	}

	// Create the Game and return error if encountered
	if err := handler.db.Create(&game).Error; err != nil {
		return results.ServerErrorResult(ctx, "Could not create game", err)
	}

	// Return the created note
	return results.OkResult(ctx, "Created game", game)
}

func (handler *GameHandler) GetAll(ctx *fiber.Ctx) error {
	var games []models.Game

	// find all games in the database
	if err := handler.db.DB.Preload("Gamers").Find(&games).Error; err != nil {
		return results.ServerErrorResult(ctx, "Cannot make a select opration", err)
	}
	// Else return games
	return results.OkResult(ctx, "Games Found", games)
}

func (handler *GameHandler) GetOne(ctx *fiber.Ctx) error {
	var game models.Game

	gameId := ctx.Params("gameId")

	if err := handler.db.DB.Preload("Gamers").First(&game, gameId).Error; err != nil {
		return results.ServerErrorResult(ctx, "Cannot make a select opration", err)
	}

	// Else return games
	return results.OkResult(ctx, "Game Found", game)
}

func (handler *GameHandler) AddGamersToGame(ctx *fiber.Ctx) error {
	payload := struct {
		GamerIds []int `json:"gamerIds"`
	}{}

	if err := ctx.BodyParser(&payload); err != nil {
		return results.BadRequestResult(ctx, "Bad request body", err)
	}

	var game models.Game
	gamers := make([]models.Gamer, len(payload.GamerIds))
	gameId := ctx.Params("gameId")

	if err := handler.db.DB.First(&game, gameId).Error; err != nil {
		return results.ServerErrorResult(ctx, "Cannot make a select opration", err)
	}

	for index, gamerId := range payload.GamerIds {
		gamers[index].ID = uint(gamerId)
	}

	assosiation := handler.db.DB.Model(&game).Association("Gamers")

	if err := assosiation.Append(gamers); err != nil {
		return results.ServerErrorResult(ctx, "Cannot make a select opration", err)
	}

	return results.OkResult(ctx, "Created game", game)
}
