package api

import (
	"net/http"
	"youtravel-api/api/controllers"
	initialzers "youtravel-api/api/initializers"
	"youtravel-api/api/middleware"

	"github.com/gin-gonic/gin"
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
	initialzers.ConnectToDB()
	initialzers.SyncDatabase()
	app = gin.New()
	r := app.Group("/api")
	iniRoute(r)
	r.POST("/signup", controllers.Signup)
	r.POST("/login", controllers.Login)
	r.GET("/validate", middleware.RequireAuth, controllers.Validate)
	controllers.Message(r)
	r.POST("/logout", middleware.RequireAuth, controllers.Logout)
}

func Handler(w http.ResponseWriter, r *http.Request) {
	app.ServeHTTP(w, r)
}
