package main

import (
	"data_service/config"
	"data_service/constants"
	"data_service/database"
)

func main() {
	db_url := config.Config(constants.ENV_FILE, constants.DB_URL_KEY_NAME)
	database.SetupDatabase(db_url)

	app := CreateApp()
	app.Listen(":3000")
}
