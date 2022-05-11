package models

type Game struct {
	ID      uint `gorm:"primarykey;AUTO_INCREMENT"`
	OwnerID uint
}
