package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

// Config func to get env value
func Config(envFilePath, key string) string {
	// load .env file
	if err := godotenv.Load(envFilePath); err != nil {
		fmt.Printf("Error loading %v file", envFilePath)
	}
	// Return the value of the variable
	return os.Getenv(key)
}
