package config

import (
	"data_service/constants"
	"testing"
)

func TestConfig(test *testing.T) {
	const envFilePath = "../.env.test"
	tests := []struct {
		name string
		key  string
		want string
	}{
		{
			"DB_URL",
			constants.DB_URL_KEY_NAME,
			"data/test_database.db",
		},
	}
	for _, tt := range tests {
		test.Run(tt.name, func(test *testing.T) {
			if got := Config(envFilePath, tt.key); got != tt.want {
				test.Errorf("Config() = %v, want %v", got, tt.want)
			}
		})
	}
}
