package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	initializers "youtravel-api/api/initializers"
	"youtravel-api/api/models"
)

func GetAllCategories(c *gin.Context) {
	var categories []models.Category

	result := initializers.DB.First(&categories)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to fetch categories",
		})
		return
	}

	c.JSON(http.StatusOK, categories)
}
