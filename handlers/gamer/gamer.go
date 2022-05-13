package gamerHandler

import (
	"data_service/database"
	"data_service/handlers/results"
	"data_service/models"

	"github.com/gofiber/fiber/v2"
)

func GetGamers(ctx *fiber.Ctx) error {
	db := database.GetInstance()
	var gamers []models.Gamer

	// find all gamers in the database

	if err := db.DB.Preload("OwnedGames").Find(&gamers).Error; err != nil {
		return results.BadRequestResult(ctx, "Cannot make a select opration", err)
	}

	// Else return gamers
	return results.OkResult(ctx, "Gamers Found", gamers)
}

func CreateGamer(ctx *fiber.Ctx) error {
	db := database.GetInstance()
	gamer := new(models.Gamer)

	// Store the body in the Gamer and return error if encountered
	if err := ctx.BodyParser(gamer); err != nil {
		return results.BadRequestResult(ctx, "Bad request body", err)
	}

	// Create the Gamer and return error if encountered
	if err := db.Create(&gamer).Error; err != nil {
		return results.ServerErrorResult(ctx, "Could not create a gamer", err)
	}

	// Return the created note
	return results.OkResult(ctx, "Created gamer", gamer)
}
