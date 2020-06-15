package helpers

import (
	"github.com/joho/godotenv"
	"github.com/rezwanul-haque/History-Service/src/logger"
	"log"
	"os"
)

func IsInvalid(value string) bool {
	if value == "" {
		return true
	}
	return false
}

func GoDotEnvVariable(key string) string {
	// load .env file
	err := godotenv.Load(".env")

	if err != nil {
		logger.Error("Error loading .env file", err)
		foundKey := os.Getenv(key)
		if foundKey != "" {
			return foundKey
		} else {
			log.Fatalf("Required environment variables not found.")
		}
	}

	return os.Getenv(key)
}
