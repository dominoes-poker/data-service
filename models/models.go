package models

type Game struct {
	ID      uint     `json:"id" gorm:"primarykey;AUTO_INCREMENT"`
	OwnerID uint     `json:"ownerId"`
	Gamers  []*Gamer `json:"gamers" gorm:"many2many:gamers_per_games;foreignKey:ID;joinForeignKey:GameID;References:ID;joinReferences:GamerID"`
}

type Gamer struct {
	ID         uint    `json:"id" gorm:"primary_key;AUTO_INCREMENT"`
	Ident      string  `json:"identificator" gorm:"uniqueIndex"`
	Nick       string  `json:"nick"`
	Name       string  `json:"name"`
	OwnedGames []Game  `json:"ownedGames" gorm:"foreignKey:OwnerID"`
	Games      []*Game `json:"games" gorm:"many2many:gamers_per_games;foreignKey:ID;joinForeignKey:GamerID;References:ID;joinReferences:GameID"`
}
