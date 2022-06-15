package playerHandler

import (
	"data_service/database"
	"data_service/handlers/results"
	"data_service/models"
	"sync"

	"github.com/gofiber/fiber/v2"
)

type PlayerHandler struct {
	db *database.DataBase
}

var instance *PlayerHandler
var gameRouterOnce sync.Once

func New(db *database.DataBase) *PlayerHandler {
	gameRouterOnce.Do(func() {
		instance = &PlayerHandler{db}
	})
	return instance
}

func (handler *PlayerHandler) GetAll(ctx *fiber.Ctx, username, identificator string) error {
	var players []models.Player

	query := handler.db.DB

	if len(username) > 0 {
		query = query.Where("Username = ?", username)
	}

	if len(identificator) > 0 {
		query = query.Where("Identificator = ?", identificator)
	}

	if err := query.Find(&players).Error; err != nil {
		return results.BadRequestResult(ctx, err)
	}

	// Else return players
	return results.OkResult(ctx, players)
}

func (handler *PlayerHandler) GetOne(playerId *uint, ctx *fiber.Ctx) error {
	var player models.Player

	if err := handler.db.DB.First(&player, playerId).Error; err != nil {
		return results.NotFoundResult(ctx, err)
	}

	// Else return players
	return results.OkResult(ctx, player)
}

func (handler *PlayerHandler) Create(context *fiber.Ctx) error {
	db := handler.db.DB
	player := new(models.Player)
	if len(player.Identificator) > 0 {
		var numberPlayers int64
		if err := db.Where("Identificator = ?", player.Identificator).First(&models.Player{}).Count(&numberPlayers).Error; err != nil {
			return results.ServerErrorResult(context, err)
		}
		if numberPlayers > 0 {
			return results.BadRequestResult(context, player)
		}

	}
	// Store the body in the Player and return error if encountered
	if err := context.BodyParser(player); err != nil {
		return results.BadRequestResult(context, err)
	}

	// Create the Player and return error if encountered
	if err := handler.db.Create(&player).Error; err != nil {
		return results.ServerErrorResult(context, err)
	}

	// Return the created note
	return results.OkResult(context, player)
}
