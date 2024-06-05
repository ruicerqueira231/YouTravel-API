package controllers

import (
	"net/http"
	initialzers "youtravel-api/api/initializers"
	"youtravel-api/api/models"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func CreateTravel(c *gin.Context) {

	var body struct {
		UserIDAdmin uint   `json:"user_id_admin"`
		CategoryID  uint   `json:"category_id"`
		Description string `json:"description"`
		Date        string `json:"date"`
		Rating      string `json:"rating"`
	}

	if err := c.BindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invaid request body",
		})
	}

	travel := models.Travel{
		UserIDAdmin: body.UserIDAdmin,
		CategoryID:  body.CategoryID,
		Description: body.Description,
		Date:        body.Date,
		Rating:      body.Rating,
	}

	if err := initialzers.DB.Create(&travel).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to create travel",
		})
	}

	response := models.Travel{
		UserIDAdmin: travel.UserIDAdmin,
		CategoryID:  travel.CategoryID,
		Description: travel.Description,
		Date:        travel.Date,
		Rating:      travel.Rating,
	}

	c.JSON(http.StatusOK, response)
}

func GetAllTravels(c *gin.Context) {
	var travels []models.Travel

	result := initialzers.DB.Find(&travels)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to fetch travels",
		})
		return
	}

	c.JSON(http.StatusOK, travels)
}

func GetTravelById(c *gin.Context) {
	id := c.Param("id")
	var travel models.Travel

	result := initialzers.DB.First(&travel, id)
	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			c.JSON(http.StatusNotFound, gin.H{
				"error": "Travel Not Found!",
			})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": "Failed to fetch data",
			})
		}
	}

	c.JSON(http.StatusOK, travel)
}

func GetTravelsByRating(c *gin.Context) {
	rating := c.Query("rating")
	if rating == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Rating parameter is required"})
		return
	}

	var travels []models.Travel

	result := initialzers.DB.Where("rating = ?", rating).Find(&travels)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch travels"})
		return
	}

	c.JSON(http.StatusOK, travels)
}
