package gamerHandler

import (
	"data_service/database"
	"data_service/handlers/results"
	"data_service/models"

	"github.com/gofiber/fiber/v2"
)

func CreateGame(ctx *fiber.Ctx) error {
	db := database.GetInstance()
	game := new(models.Game)

	// Store the body in the Game and return error if encountered
	if err := ctx.BodyParser(game); err != nil {
		return results.BadRequestResult(ctx, "Bad request body", err)
	}

	// Create the Game and return error if encountered
	if err := db.Create(&game).Error; err != nil {
		return results.ServerErrorResult(ctx, "Could not create game", err)
	}

	// Return the created note
	return results.OkResult(ctx, "Created game", game)
}

func GetAllGames(ctx *fiber.Ctx) error {
	db := database.GetInstance()
	var games []models.Game

	// find all games in the database
	if err := db.DB.Preload("Gamers").Find(&games).Error; err != nil {
		return results.ServerErrorResult(ctx, "Cannot make a select opration", err)
	}
	// Else return games
	return results.OkResult(ctx, "Games Found", games)
}

func GetGame(ctx *fiber.Ctx) error {
	db := database.GetInstance()
	var game models.Game

	gameId := ctx.Params("gameId")

	if err := db.DB.First(&game, gameId); err != nil {
		return results.ServerErrorResult(ctx, "Cannot make a select opration", err)
	}

	// Else return games
	return results.OkResult(ctx, "Game Found", game)
}

func AddGamersToGame(ctx *fiber.Ctx) error {
	payload := struct {
		GamerIds []int `json:"gamerIds"`
	}{}

	if err := ctx.BodyParser(&payload); err != nil {
		return results.BadRequestResult(ctx, "Bad request body", err)
	}

	var game models.Game
	gamers := make([]models.Gamer, len(payload.GamerIds))
	gameId := ctx.Params("gameId")
	db := database.GetInstance()

	if err := db.DB.First(&game, gameId).Error; err != nil {
		return results.ServerErrorResult(ctx, "Cannot make a select opration", err)
	}

	for index, gamerId := range payload.GamerIds {
		gamers[index].ID = uint(gamerId)
	}

	assosiation := db.DB.Model(&game).Association("Gamers")

	if err := assosiation.Append(gamers).Error; err != nil {
		return results.ServerErrorResult(ctx, "Cannot make a select opration", err)
	}

	return results.OkResult(ctx, "Created game", game)
}
