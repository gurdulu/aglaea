package inits

import (
	"github.com/joho/godotenv"
	"log"
)

func LoadEnvs() {
	err := godotenv.Load()
	if err != nil {
		err = godotenv.Load(".env.dev")
		if err != nil {
			log.Fatal("Error loading .env file")
		}
	}
}
