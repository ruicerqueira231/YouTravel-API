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

	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	} else {
		fmt.Println(".env file loaded successfully")
	}

	env := os.Getenv("ENV")
	fmt.Println("Current ENV:", env)
}
