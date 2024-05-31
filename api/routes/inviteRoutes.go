package routes

import (
	"youtravel-api/api/controllers"

	"github.com/gin-gonic/gin"
)

func InviteRoutes(r *gin.RouterGroup) {
	r.POST("/invite", controllers.CreateInvite)
}
