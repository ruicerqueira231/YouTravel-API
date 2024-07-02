package initialzers

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

func LoadEnvVariables() {
	cwd, _ := os.Getwd()
	fmt.Println("Current working directory:", cwd)

	env := os.Getenv("ENV")
	fmt.Println("Current ENV:", env)
	if env != "production" {
		err := godotenv.Load()
		if err != nil {
			log.Fatalf("Error loading .env file: %v", err)
		} else {
			fmt.Println(".env file loaded successfully")
		}
	} else {
		fmt.Println("Running in production environment, not loading .env file.")
	}
}
