package models

import "time"

type Game struct {
	ID        uint      `json:"id"        gorm:"primarykey"`
	CreatedAt time.Time `json:"createdAt"`
	OwnerID   uint      `json:"ownerId"`
	Gamers    []*Gamer  `json:"gamers"    gorm:"many2many:gamers_per_games;foreignKey:ID;joinForeignKey:GameID;References:ID;joinReferences:GamerID"`
	Rounds    []*Round  `json:"rounds"    gorm:"many2many:gamers_per_games;foreignKey:ID;joinForeignKey:GameID;References:ID;joinReferences:GamerID"`
	IsOver    bool      `json:"isOver"`
}

type Gamer struct {
	ID            uint    `json:"id"            gorm:"primary_key;AUTO_INCREMENT"`
	Identificator string  `json:"identificator" gorm:"uniqueIndex"`
	Name          string  `json:"name"          gorm:"uniqueIndex"`
	Games         []*Game `json:"games"         gorm:"many2many:gamers_per_games;foreignKey:ID;joinForeignKey:GamerID;References:ID;joinReferences:GameID"`
}

type Round struct {
	ID           uint     `json:"id"           gorm:"primary_key;AUTO_INCREMENT"`
	GameID       uint     `json:"gameId"`
	NumberOfDice uint     `json:"numberOfDice"`
	Number       uint     `json:"number"`
	Stakes       []*Stake `json:"stakes"       gorm:"preload:true"`
}

type Stake struct {
	ID      uint `json:"id"      gorm:"primary_key;AUTO_INCREMENT"`
	RoundID uint `json:"roundId"`
	GamerID uint `json:"gamerId"`
	Bet     uint `json:"bet"`
	Get     uint `json:"get"`
}
