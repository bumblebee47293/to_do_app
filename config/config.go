package config

import "os"

func LoadConfig() {
	// Example: Load configurations from environment variables
	os.Setenv("DB_DRIVER", "sqlite")
	os.Setenv("DB_NAME", "test.db")
}
