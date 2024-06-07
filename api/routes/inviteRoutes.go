package routes

import (
	"youtravel-api/api/controllers"

	"github.com/gin-gonic/gin"
)

func InviteRoutes(r *gin.RouterGroup) {
	r.POST("/invite", controllers.CreateInvite)
	r.POST("/invite/changeStatusToAccepted/:id", controllers.ChangeStatusAcceptedInvited)
	r.POST("/invite/changeStatusToDeclined/:id", controllers.ChangeStatusDeclinedInvited)
}
