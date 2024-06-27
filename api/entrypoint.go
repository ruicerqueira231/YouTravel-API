package api

import (
	"fmt"
	"net/http"
	"os"
	initialzers "youtravel-api/api/initializers"
	"youtravel-api/api/routes"

	"github.com/gin-gonic/gin"
)

var (
	app *gin.Engine
)

func PrintAllEnvVariables() {
	for _, env := range os.Environ() {
		fmt.Println(env)
	}
}

func init() {
	PrintAllEnvVariables()

	initialzers.LoadEnvVariables()

	initialzers.ConnectToDB()
	initialzers.SyncDatabase()
	initialzers.InitCategories(initialzers.DB)
	initialzers.InitLocationCategories(initialzers.DB)

	app = gin.New()
	r := app.Group("/api")

	//routes
	routes.UserRoutes(r)
	routes.TravelRoutes(r)
	routes.InviteRoutes(r)
	routes.CategoriesRoutes(r)
	routes.LocationRoutes(r)
}

func Handler(w http.ResponseWriter, r *http.Request) {
	app.ServeHTTP(w, r)
}
