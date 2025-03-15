package initializers

import (
	"log"

	"github.com/joho/godotenv"
)

func LoadEnvs() {
	err := godotenv.Load()

	if err != nil {
		log.Fatal("error loading .env file", err)
	}
}
