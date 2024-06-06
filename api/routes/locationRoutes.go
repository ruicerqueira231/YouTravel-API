package routes

import (
	"github.com/gin-gonic/gin"
	"youtravel-api/api/controllers"
)

func LocationRoutes(r *gin.RouterGroup) {
	r.POST("/location", controllers.CreateLocation)
	r.GET("/locations", controllers.GetAllLocations)
	r.GET("/locations/:id", controllers.GetLocationById)
	r.GET("/locations/travel/:travel_id", controllers.GetLocationsByTravelID)
}
