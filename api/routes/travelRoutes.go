package routes

import (
	"youtravel-api/api/controllers"

	"github.com/gin-gonic/gin"
)

func TravelRoutes(r *gin.RouterGroup) {
	r.POST("/travel", controllers.CreateTravel)
	r.GET("/travels", controllers.GetAllTravels)
	r.GET("/travel/:id", controllers.GetTravelById)
	r.GET("/travelByRating", controllers.GetTravelsByRating)
}