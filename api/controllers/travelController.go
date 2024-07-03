package controllers

import (
	"context"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"log"
	"net/http"
	"youtravel-api/api/dto"
	initialzers "youtravel-api/api/initializers"
	"youtravel-api/api/models"
)

func CreateTravel(c *gin.Context) {

	var body struct {
		UserIDAdmin uint   `json:"user_id_admin"`
		CategoryID  uint   `json:"category_id"`
		Title       string `json:"title"`
		Description string `json:"description"`
		Date        string `json:"date"`
		Rating      string `json:"rating"`
		Photo       string `json:"photo"`
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
		Rating:      body.Rating,
		Photo:       body.Photo,
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

	for _, t := range travels {
		log.Printf("Travel: %v, User: %v\n\n", t.Title, t.User)
	}

	travelDTOs := make([]dto.TravelDTO, len(travels))
	for i, t := range travels {
		travelDTOs[i] = dto.TravelDTO{
			ID:          t.ID,
			Title:       t.Title,
			Description: t.Description,
			Date:        t.Date,
			PhotoURL:    t.Photo,
			Rating:      t.Rating,
			Category:    t.Category.Description,
			User:        t.User.Nome,
			UserPhoto:   t.User.Photo,
		}
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

	travelDTO := dto.TravelDTO{
		ID:          travel.ID,
		Title:       travel.Title,
		Description: travel.Description,
		Date:        travel.Date,
		PhotoURL:    travel.Photo,
		Rating:      travel.Rating,
		Category:    travel.Category.Description,
		User:        travel.User.Nome,
		UserPhoto:   travel.User.Photo,
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
	userId := c.Param("id")
	var participations []models.Participation

	result := initialzers.DB.Where("user_id = ?", userId).
		Preload("Travel").
		Preload("Travel.User").
		Preload("Travel.Category").
		Find(&participations)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch participations"})
		return
	}

	var travelDTO []dto.TravelDTO

	for _, p := range participations {
		travelDTO = append(travelDTO, dto.TravelDTO{
			ID:          p.Travel.ID,
			Title:       p.Travel.Title,
			Description: p.Travel.Description,
			Date:        p.Travel.Date,
			PhotoURL:    p.Travel.Photo,
			Rating:      p.Travel.Rating,
			Category:    p.Travel.Category.Description,
			User:        p.Travel.User.Nome,
			UserPhoto:   p.Travel.User.Photo,
		})
	}
	c.JSON(http.StatusOK, travelDTO)
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

func GetTravelImage(c *gin.Context) {
	travelID := c.Param("id")
	if travelID == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Missing travel ID"})
		return
	}

	var travel models.Travel
	result := initialzers.DB.First(&travel, travelID)
	if result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Travel not found"})
		return
	}

	if travel.Photo == "" {
		c.JSON(http.StatusNotFound, gin.H{"error": "No photo available for this travel"})
		return
	}

	cfg, err := config.LoadDefaultConfig(context.TODO())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to load AWS config"})
		return
	}

	client := s3.NewFromConfig(cfg)
	bucket := "you-travel-storage"
	key := travel.Photo

	resp, err := client.GetObject(context.TODO(), &s3.GetObjectInput{
		Bucket: aws.String(bucket),
		Key:    aws.String(key),
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("Failed to retrieve file from S3: %v", err)})
		return
	}
	defer resp.Body.Close()

	contentType := "application/octet-stream"
	if resp.ContentType != nil {
		contentType = *resp.ContentType
	}

	c.DataFromReader(http.StatusOK, *resp.ContentLength, contentType, resp.Body, nil)
}
