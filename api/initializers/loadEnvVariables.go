package initialzers

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func LoadEnvVariables() {
	env := os.Getenv("ENV")
	if env == "" {
		log.Println("ENV variable not set or accessible")
	} else {
		log.Println("Current ENV:", env)
	}

	if env != "production" {
		err := godotenv.Load()
		if err != nil {
			log.Fatalf("Error loading .env file: %v", err)
		}
	}
}
