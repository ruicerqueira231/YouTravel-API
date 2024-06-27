package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"youtravel-api/api/dto"
	initializers "youtravel-api/api/initializers"
	"youtravel-api/api/models"
)

func GetAllCategories(c *gin.Context) {
	var categories []models.Category
	result := initializers.DB.Find(&categories)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to fetch categories",
		})
		return
	}

	// Transform categories to CategoryDTO
	categoryDTOs := make([]dto.CategoryDTO, len(categories))
	for i, category := range categories {
		categoryDTOs[i] = dto.CategoryDTO{
			ID:          category.ID,
			Description: category.Description,
		}
	}

	// Return DTOs as JSON
	c.JSON(http.StatusOK, categoryDTOs)
}
