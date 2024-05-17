package controllers

import (
	"net/http"
	initialzers "youtravel-api/api/initializers"
	"youtravel-api/api/models"

	"github.com/gin-gonic/gin"
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

	c.JSON(http.StatusOK, travel)
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
