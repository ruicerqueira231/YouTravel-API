package api

import (
	"net/http"
	"youtravel-api/api/controllers"
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
	app = gin.New()
	r := app.Group("/api")
	iniRoute(r)
	controllers.SignupAPI(r)
	r.POST("/login", controllers.Login)
	r.GET("/validate", middleware.RequireAuth, controllers.Validate)
	controllers.Message(r)
	r.POST("/logout", middleware.RequireAuth, controllers.Logout)
}

func Handler(w http.ResponseWriter, r *http.Request) {
	app.ServeHTTP(w, r)
}
