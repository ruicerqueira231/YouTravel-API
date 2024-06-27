package controllers

import (
	"context"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
	"youtravel-api/api/dto"
	initialzers "youtravel-api/api/initializers"
	"youtravel-api/api/models"
)

func ImageUpload(c *gin.Context) {
	cfg, err := config.LoadDefaultConfig(context.TODO())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to load AWS config"})
		return
	}

	client := s3.NewFromConfig(cfg)

	file, header, err := c.Request.FormFile("fileUpload")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "file retrieval error"})
		return
	}
	defer file.Close()

	bucket := "you-travel-storage"
	key := header.Filename

	_, err = client.PutObject(context.TODO(), &s3.PutObjectInput{
		Bucket: aws.String(bucket),
		Key:    aws.String(key),
		Body:   file,
		ACL:    "public-read",
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to upload to S3"})
		return
	}
}

func CreateTravel(c *gin.Context) {

	var body struct {
		UserIDAdmin uint   `json:"user_id_admin"`
		CategoryID  uint   `json:"category_id"`
		Title       string `json:"title"`
		Description string `json:"description"`
		Date        string `json:"date"`
		Rating      string `json:"rating"`
		PhotoURL    string `json:"photo_url"`
	}

	if err := c.BindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invaid request body",
		})
	}

	travel := models.Travel{
		UserIDAdmin: body.UserIDAdmin,
		CategoryID:  body.CategoryID,
		Title:       body.Title,
		Description: body.Description,
		Date:        body.Date,
		Rating:      body.Rating,
		Photo:       body.PhotoURL,
	}

	if err := initialzers.DB.Create(&travel).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to create travel",
		})
	}

	participation := models.Participation{
		UserID:   body.UserIDAdmin,
		TravelID: travel.ID,
	}

	if err := initialzers.DB.Create(&participation).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to create participation for admin",
		})
	}

	c.JSON(http.StatusOK, travel)
}

func GetAllTravels(c *gin.Context) {
	var travels []models.Travel

	result := initialzers.DB.Preload("User").Preload("Category").Find(&travels)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to fetch travels",
		})
		return
	}

	// converte travel em DTOs
	travelDTOs := make([]dto.TravelDTO, len(travels))
	for i, t := range travels {
		travelDTOs[i] = dto.TravelDTO{
			ID:          t.ID,
			UserIDAdmin: t.UserIDAdmin,
			CategoryID:  t.CategoryID,
			Title:       t.Title,
			Description: t.Description,
			Date:        t.Date,
			Rating:      t.Rating,
			Category:    t.Category.Description}
	}

	c.JSON(http.StatusOK, travelDTOs)
}

func GetTravelById(c *gin.Context) {
	id := c.Param("id")
	var travel models.Travel

	result := initialzers.DB.Preload("User").Preload("Category").First(&travel, id)
	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			c.JSON(http.StatusNotFound, gin.H{"error": "Travel Not Found!"})
			return
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch data"})
			return
		}
	}

	// Convert the travel model to DTO
	travelDTO := dto.TravelDTO{
		ID:          travel.ID,
		UserIDAdmin: travel.UserIDAdmin,
		CategoryID:  travel.CategoryID,
		Title:       travel.Title,
		Description: travel.Description,
		Date:        travel.Date,
		Rating:      travel.Rating,
		Category:    travel.Category.Description,
	}

	c.JSON(http.StatusOK, travelDTO)
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

func GetTravelsByUserId(c *gin.Context) {
	userId := c.Param("userId")
	var participations []models.Participation

	// Find all participations for the user
	result := initialzers.DB.Where("user_id = ?", userId).Preload("Travel").Find(&participations)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch participations"})
		return
	}

	var travelDTOs []dto.TravelDTO
	for _, p := range participations {
		travelDTOs = append(travelDTOs, dto.TravelDTO{
			ID:          p.Travel.ID,
			UserIDAdmin: p.Travel.UserIDAdmin,
			CategoryID:  p.Travel.CategoryID,
			Title:       p.Travel.Title,
			Description: p.Travel.Description,
			Date:        p.Travel.Date,
			Rating:      p.Travel.Rating,
			Category:    p.Travel.Category.Description,
		})
	}

	c.JSON(http.StatusOK, travelDTOs)
}

func DeleteTravel(c *gin.Context) {
	travelID := c.Param("id")

	result := initialzers.DB.Delete(&models.Travel{}, travelID)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete travel"})
		return
	}

	if result.RowsAffected == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "No travel found with given ID"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Travel deleted successfully"})
}
