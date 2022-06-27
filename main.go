package main

import (
	"data_service/config"
	"data_service/constants"
	"data_service/database"
)

func main() {
	app := CreateApp()

	var db_url string = config.Config(constants.ENV_FILE, constants.DB_URL_KEY_NAME)
	var db *database.DataBase = database.SetupDatabase(db_url)
	SetupRoutes(app, "api", db)

	if err := app.Listen(":3000"); err != nil {
		panic("Cannot run application")
	}
}
