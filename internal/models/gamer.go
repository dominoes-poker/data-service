package models

type Gamer struct {
	ID    uint   `gorm:"primary_key;AUTO_INCREMENT"`
	Ident string `gorm:"uniqueIndex"`
	Nick  string
	Name  string
	Games []Game `gorm:"foreignKey:OwnerID"`
}
