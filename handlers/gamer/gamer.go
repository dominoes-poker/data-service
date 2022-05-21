package gamerHandler

import (
	"data_service/database"
	"data_service/handlers/results"
	"data_service/models"
	"sync"

	"github.com/gofiber/fiber/v2"
)

type GamerHandler struct {
	db *database.DataBase
}

var instance *GamerHandler
var gameRouterOnce sync.Once

func New(db *database.DataBase) *GamerHandler {
	gameRouterOnce.Do(func() {
		instance = &GamerHandler{db}
	})
	return instance
}

func (handler *GamerHandler) GetAll(ctx *fiber.Ctx) error {
	var gamers []models.Gamer

	identificator := ctx.Query("identificator")
	// find all gamers in the database

	if err := handler.db.DB.Where("Identificator = ?", identificator).Preload("OwnedGames").Find(&gamers).Error; err != nil {
		return results.BadRequestResult(ctx, "Cannot make a select opration", err)
	}

	// Else return gamers
	return results.OkResult(ctx, "Gamers Found", gamers)
}

func (handler *GamerHandler) GetOne(ctx *fiber.Ctx) error {
	var gamer models.Gamer

	gamerId := ctx.Params("gamerIdentificator")
	// find all gamers in the database

	if err := handler.db.DB.Preload("OwnedGames").First(&gamer, gamerId).Error; err != nil {
		return results.BadRequestResult(ctx, "Cannot make a select opration", err)
	}

	// Else return gamers
	return results.OkResult(ctx, "Gamer Found", gamer)
}

func (handler *GamerHandler) Create(ctx *fiber.Ctx) error {
	gamer := new(models.Gamer)

	// Store the body in the Gamer and return error if encountered
	if err := ctx.BodyParser(gamer); err != nil {
		return results.BadRequestResult(ctx, "Bad request body", err)
	}

	// Create the Gamer and return error if encountered
	if err := handler.db.Create(&gamer).Error; err != nil {
		return results.ServerErrorResult(ctx, "Could not create a gamer", err)
	}

	// Return the created note
	return results.OkResult(ctx, "Created gamer", gamer)
}
