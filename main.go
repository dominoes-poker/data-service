package main

import (
	"data_service/config"
)

func main() {
	db_url := config.Config("DB_URL")
	setup_database(db_url)

	app := create_app()
	app.Listen(":3000")
}
