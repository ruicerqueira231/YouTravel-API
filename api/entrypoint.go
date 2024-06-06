package api

import (
	"net/http"
	initialzers "youtravel-api/api/initializers"
	"youtravel-api/api/routes"

	"github.com/gin-gonic/gin"
)

var (
	app *gin.Engine
)

func init() {
	//initializers
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
