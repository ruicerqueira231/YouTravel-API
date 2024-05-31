package routes

import (
	"youtravel-api/api/controllers"
	"youtravel-api/api/middleware"

	"github.com/gin-gonic/gin"
)

func UserRoutes(r *gin.RouterGroup) {
	r.POST("/signup", controllers.Signup)
	r.POST("/login", controllers.Login)
	r.GET("/validate", middleware.RequireAuth, controllers.Validate)
	r.POST("/logout", middleware.RequireAuth, controllers.Logout)
	r.GET("/users", controllers.GetAllUsers)
	r.GET("/user/:id", controllers.GetUserById)
	r.GET("/userByEmail", controllers.GetUserByEmail)
	r.GET("/userByUsername", controllers.GetUserByUsername)
}
