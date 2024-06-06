package controllers

import (
	"errors"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
	initializers "youtravel-api/api/initializers"
	"youtravel-api/api/models"
)

func CreateLocation(c *gin.Context) {
	var body struct {
		TravelID           uint   `json:"travel_id"`
		LocationCategoryID uint   `json:"location_category_id"`
		Nome               string `json:"nome"`
		Coordinates        string `json:"coordinates"`
		Latitude           string `json:"latitude"`
		Longitude          string `json:"longitude"`
	}

	if err := c.BindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	location := models.Location{
		TravelID:           body.TravelID,
		LocationCategoryID: body.LocationCategoryID,
		Nome:               body.Nome,
		Coordinates:        body.Coordinates,
		Latitude:           body.Latitude,
		Longitude:          body.Longitude,
	}

	if err := initializers.DB.Create(&location).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create location"})
		return
	}

	c.JSON(http.StatusOK, location)
}

func GetAllLocations(c *gin.Context) {
	var locations []models.Location
	if err := initializers.DB.Find(&locations).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch locations"})
		return
	}

	c.JSON(http.StatusOK, locations)
}

func GetLocationById(c *gin.Context) {
	id := c.Param("id")
	var location models.Location
	if err := initializers.DB.First(&location, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusNotFound, gin.H{"error": "Location not found"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch data"})
		}
		return
	}

	c.JSON(http.StatusOK, location)
}

func GetLocationsByTravelID(c *gin.Context) {
	travelID := c.Param("travel_id")
	if travelID == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Travel ID parameter is required"})
		return
	}

	var locations []models.Location
	if err := initializers.DB.Where("travel_id = ?", travelID).Find(&locations).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch locations"})
		return
	}

	if len(locations) == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "No locations found for the given travel ID"})
		return
	}

	c.JSON(http.StatusOK, locations)
}
