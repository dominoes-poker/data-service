package gamerHandler

import (
	"data_service/database"
	"data_service/handlers/results"
	"data_service/models"
	"strings"
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

	query := handler.db.DB

	if identificator := ctx.Query("identificator"); len(identificator) > 0 {
		query = query.Where("Identificator = ?", identificator)
	}

	if name := ctx.Query("name"); len(name) > 0 {
		query = query.Or("Name IN ?", strings.Split(name, ","))
	}

	if err := query.Find(&gamers).Error; err != nil {
		return results.BadRequestResult(ctx, err)
	}

	// Else return gamers
	return results.OkResult(ctx, gamers)
}

func (handler *GamerHandler) GetOne(gamerIdentificator string, ctx *fiber.Ctx) error {
	var gamer models.Gamer

	if err := handler.db.DB.First(&gamer, "Identificator = ?", gamerIdentificator).Error; err != nil {
		return results.NotFoundResult(ctx, err)
	}

	// Else return gamers
	return results.OkResult(ctx, gamer)
}

func (handler *GamerHandler) Create(ctx *fiber.Ctx) error {
	gamer := new(models.Gamer)

	// Store the body in the Gamer and return error if encountered
	if err := ctx.BodyParser(gamer); err != nil {
		return results.BadRequestResult(ctx, err)
	}

	// Create the Gamer and return error if encountered
	if err := handler.db.Create(&gamer).Error; err != nil {
		return results.ServerErrorResult(ctx, err)
	}

	// Return the created note
	return results.OkResult(ctx, gamer)
}
