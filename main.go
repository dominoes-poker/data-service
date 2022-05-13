package main

import (
	"data_service/config"
	"data_service/constants"
	"data_service/database"
)

func main() {
	db_url := config.Config(constants.ENV_FILE, constants.DB_URL_KEY_NAME)
	db := database.SetupDatabase(db_url)

	app := CreateApp(db)
	if err := app.Listen(":3000"); err != nil {
		panic("Cannot run application")
	}
}
