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
	//initializers
	initialzers.ConnectToDB()
	initialzers.SyncDatabase()
	initialzers.InitCategories(initialzers.DB)

	//routes
	app = gin.New()
	r := app.Group("/api")
	iniRoute(r)

	//User routes
	r.POST("/signup", controllers.Signup)
	r.POST("/login", controllers.Login)
	r.GET("/validate", middleware.RequireAuth, controllers.Validate)
	r.POST("/logout", middleware.RequireAuth, controllers.Logout)
	r.GET("/getUsers", controllers.GetAllUsers)
	r.GET("/getUserById/:id", controllers.GetUserById)
}

func Handler(w http.ResponseWriter, r *http.Request) {
	app.ServeHTTP(w, r)
}
