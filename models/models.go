package models

import "time"

type Game struct {
	ID        uint      `json:"id"        gorm:"primarykey"`
	CreatedAt time.Time `json:"createdAt"`
	Players   []*Player `json:"players"    gorm:"many2many:players_per_games;foreignKey:ID;joinForeignKey:GameID;References:ID;joinReferences:PlayerID"`
	Rounds    []*Round  `json:"rounds"`
	IsOver    bool      `json:"isOver"`
}

type Player struct {
	ID            uint    `json:"id"            gorm:"primary_key;AUTO_INCREMENT"`
	Identificator string  `json:"identificator"`
	Username      string  `json:"username"      gorm:"uniqueIndex"`
	Games         []*Game `json:"games"         gorm:"many2many:players_per_games;foreignKey:ID;joinForeignKey:PlayerID;References:ID;joinReferences:GameID"`
}

type Round struct {
	ID           uint     `json:"id"           gorm:"primary_key;AUTO_INCREMENT"`
	GameID       uint     `json:"gameId"`
	NumberOfDice uint     `json:"numberOfDice"`
	Number       uint     `json:"number"`
	Stakes       []*Stake `json:"stakes"       gorm:"preload:true"`
}

type Stake struct {
	ID       uint `json:"id"      gorm:"primary_key;AUTO_INCREMENT"`
	RoundID  uint `json:"roundId"`
	PlayerID uint `json:"playerId"`
	Bet      uint `json:"bet"`
	Get      uint `json:"get"`
}
