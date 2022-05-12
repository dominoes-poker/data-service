package main

import (
	"data_service/config"
	"data_service/database"
)

func main() {
	db_url := config.Config("DB_URL")
	database.SetupDatabase(db_url)

	app := CreateApp()
	app.Listen(":3000")
}
