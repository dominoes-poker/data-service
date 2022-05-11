package models

type Game struct {
	ID      uint  `json:"id" gorm:"primarykey;AUTO_INCREMENT"`
	OwnerID uint  `json:"owner_id"`
	Owner   Gamer `json:"owner" gorm:"foreignkey:OwnerID;`
}

type Gamer struct {
	ID    uint   `json:"id" gorm:"primary_key;AUTO_INCREMENT"`
	Ident string `json:"identificator" gorm:"uniqueIndex"`
	Nick  string `json:"nick"`
	Name  string `json:"name"`
	Games []Game `json:"games" gorm:"foreignKey:OwnerID"`
}
