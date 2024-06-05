package routes

import (
	"github.com/gin-gonic/gin"
	"youtravel-api/api/controllers"
)

func CategoriesRoutes(r *gin.RouterGroup) {
	r.GET("/categories", controllers.GetAllCategories)
}
