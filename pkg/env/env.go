package env

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func Setup() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

func GET(key string) string {
	return os.Getenv(key)
}
