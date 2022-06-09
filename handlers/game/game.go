package playerHandler

import (
	"data_service/database"
	"data_service/handlers/results"
	"data_service/models"
	"sync"

	"github.com/gofiber/fiber/v2"

	"gorm.io/gorm/clause"
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

func (handler *GameHandler) getGame(gameId uint) (models.Game, error) {
	var game models.Game
	db := handler.db.DB

	err := db.Preload(clause.Associations).Preload("Rounds.Stakes").Find(&game, gameId).Error
	return game, err
}

func (handler *GameHandler) getAllGames() ([]models.Game, error) {
	var games []models.Game
	db := handler.db.DB

	err := db.Preload(clause.Associations).Preload("Rounds.Stakes").Find(&games).Error
	return games, err
}

func (handler *GameHandler) Create(context *fiber.Ctx) error {
	game := new(models.Game)

	// Store the body in the Game and return error if encountered
	if err := context.BodyParser(game); err != nil {
		return results.BadRequestResult(context, err)
	}

	// Create the Game and return error if encountered
	if err := handler.db.Create(&game).Error; err != nil {
		return results.ServerErrorResult(context, err)
	}

	if game, err := handler.getGame(game.ID); err != nil {
		return results.ServerErrorResult(context, err)
	} else {
		return results.OkResult(context, game)
	}
}

func (handler *GameHandler) GetAll(context *fiber.Ctx) error {
	if games, err := handler.getAllGames(); err != nil {
		return results.ServerErrorResult(context, err)
	} else {
		return results.OkResult(context, games)
	}
}

func (handler *GameHandler) GetOne(gameId uint, context *fiber.Ctx) error {

	if game, err := handler.getGame(gameId); err != nil {
		return results.ServerErrorResult(context, err)
	} else {
		return results.OkResult(context, game)
	}
}

func (handler *GameHandler) AddPlayersToGame(gameId uint, context *fiber.Ctx) error {
	payload := struct {
		PlayerIds []int `json:"playerIds"`
	}{}

	if err := context.BodyParser(&payload); err != nil {
		return results.BadRequestResult(context, err)
	}

	var game models.Game
	players := make([]models.Player, len(payload.PlayerIds))

	if err := handler.db.DB.Preload("Players").Find(&game, gameId).Error; err != nil {
		return results.ServerErrorResult(context, err)
	}

	for index, playerId := range payload.PlayerIds {
		players[index].ID = uint(playerId)
	}

	assosiation := handler.db.DB.Model(&game).Association("Players")

	if err := assosiation.Append(players); err != nil {
		return results.ServerErrorResult(context, err)
	}

	if err := handler.db.DB.Preload("Players").Find(&game, gameId).Error; err != nil {
		return results.ServerErrorResult(context, err)
	}

	if game, err := handler.getGame(gameId); err != nil {
		return results.ServerErrorResult(context, err)
	} else {
		return results.OkResult(context, game)
	}
}

func (handler *GameHandler) StartRound(gameId uint, context *fiber.Ctx) error {
	var round models.Round

	if err := context.BodyParser(&round); err != nil {
		return results.BadRequestResult(context, err)
	}
	var numberOfDoneRounds int64

	round.GameID = uint(gameId)

	if err := handler.db.DB.Model(&models.Round{}).Where("game_id = ?", gameId).Count(&numberOfDoneRounds).Error; err != nil {
		return results.ServerErrorResult(context, err)
	}

	round.Number = uint(numberOfDoneRounds + 1)

	if err := handler.db.Create(&round).Error; err != nil {
		return results.ServerErrorResult(context, err)
	}

	return results.OkResult(context, round)
}

func (handler *GameHandler) MakeBet(gameId, roundNumber uint, context *fiber.Ctx) error {
	db := handler.db.DB
	var stake models.Stake

	if err := context.BodyParser(&stake); err != nil {
		return results.BadRequestResult(context, err)
	}

	var round models.Round

	if err := db.Where("game_id = ?", gameId).Where("number = ?", roundNumber).Find(&round).Error; err != nil {
		return results.ServerErrorResult(context, err)
	}

	stake.RoundID = round.ID

	if err := db.Create(&stake).Error; err != nil {
		return results.ServerErrorResult(context, err)
	}

	game, err := handler.getGame(gameId)
	if err != nil {
		return results.ServerErrorResult(context, err)
	}

	return results.OkResult(context, game)
}
