package models

type Gamer struct {
	UserId int64 `gorm:"primarykey"`
	Nick   string
	Name   string
}
