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

	if res := db.DB.Exec("PRAGMA foreign_keys = ON", nil); res.Error != nil {
		panic("Failed to enable foreign keys")
	}

	fmt.Println("Connection Opened to Database")
}

func (db *DataBase) InitializeTables(models ...interface{}) {
	if err := db.DB.AutoMigrate(models...); err != nil {
		panic("Migration Failed")
	}
}

func (db *DataBase) Create(model interface{}) *gorm.DB {
	return db.DB.Create(model)
}
