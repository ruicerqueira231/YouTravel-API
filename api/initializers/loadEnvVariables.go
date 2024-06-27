package initialzers

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

func LoadEnvVariables() {
	env := os.Getenv("ENV")
	if env == "" {
		fmt.Println("ENV variable not set or accessible, defaulting to production for safety.")
		env = "production"
	}
	fmt.Println("Current ENV:", env)
	if env != "production" {
		err := godotenv.Load()
		if err != nil {
			log.Fatalf("Error loading .env file: %v", err)
		}
	}
}
