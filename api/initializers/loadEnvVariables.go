package initialzers

import (
	"fmt"
	"github.com/joho/godotenv"
	"os"
)

func LoadEnvVariables() {
	cwd, _ := os.Getwd()
	fmt.Println("Current working directory:", cwd)

	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error loading .env file")
	}
}
