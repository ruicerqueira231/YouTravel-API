package initialzers

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func LoadEnvVariables() {
	env := os.Getenv("ENV")
	if env == "" {
		log.Printf("ENV variable not set, defaulting to .env file loading")
		err := godotenv.Load()
		if err != nil {
			log.Fatalf("Error loading .env file: %v", err)
		}
	} else {
		log.Printf("Current ENV: %s", env)
		if env != "production" {
			err := godotenv.Load()
			if err != nil {
				log.Fatalf("Error loading .env file: %v", err)
			}
		}
	}
}
