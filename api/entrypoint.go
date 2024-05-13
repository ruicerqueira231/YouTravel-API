package main

import (
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

var (
	app *gin.Engine
)

func iniRoute(r *gin.RouterGroup) {
	r.GET("/hello", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "I'm in Vercel",
		})
	})
}

func init() {
	app = gin.New()
	r := app.Group("/api")
	iniRoute(r)
}

func main() {

	println("Starting the application...")
	app = gin.New()
	r := app.Group("/api")
	iniRoute(r)

	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	if os.Getenv("ENV") == "development" {
		println("Running in development mode on port 8080...")
		app.Run(":8080")
	} else {
		println("ENV not set to 'development', application will not start the server.")
	}
}

func Handler(w http.ResponseWriter, r *http.Request) {
	app.ServeHTTP(w, r)
}
