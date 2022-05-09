package database

import (
	"fmt"
	"sync"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type DataBase struct {
	DB *gorm.DB
}

var instance *DataBase
var dbOnce sync.Once

func GetInstance() *DataBase {
	dbOnce.Do(func() {
		instance = &DataBase{}
	})
	return instance
}

func (db *DataBase) Connect(databaseConnection string) {
	var err error

	db.DB, err = gorm.Open(sqlite.Open(databaseConnection))

	if err != nil {
		panic("Failed to connect database")
	}

	fmt.Println("Connection Opened to Database")
}

func (db *DataBase) InitializeTables(models ...interface{}) error {
	return db.DB.AutoMigrate(models...)
}

func (db *DataBase) Create(model interface{}) (tx *gorm.DB) {
	return db.DB.Create(model)
}
