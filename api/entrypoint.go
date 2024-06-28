package api

import (
	"context"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"gorm.io/gorm"
	"net/http"
	initializers "youtravel-api/api/initializers"
	"youtravel-api/api/routes"

	"github.com/gin-gonic/gin"
)

var (
	app *gin.Engine
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

	// Success response
	c.JSON(http.StatusOK, gin.H{"message": "Successfully uploaded file"})
}

func dropTravelData(db *gorm.DB) error {
	// Using GORM to delete all records from the travel table
	result := db.Exec("DELETE FROM travel")
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func init() {
	initializers.LoadEnvVariables()

	// Initialize database connection with GORM
	initializers.ConnectToDB() // Assuming this function handles errors internally

	// Optionally synchronize database schema
	initializers.SyncDatabase() // Assuming this function also handles errors internally

	// Attempt to drop travel data, handle error if one occurs
	if err := dropTravelData(initializers.DB); err != nil {
		panic(err) // Handle error according to your error handling policy
	}

	// Initialize other database-related configurations
	initializers.InitCategories(initializers.DB)
	initializers.InitLocationCategories(initializers.DB)

	// Set up the Gin application
	app = gin.New()
	r := app.Group("/api")

	// Set up API routes
	routes.UserRoutes(r)
	routes.TravelRoutes(r)
	routes.InviteRoutes(r)
	routes.CategoriesRoutes(r)
	routes.LocationRoutes(r)

	// Image upload and retrieval endpoints
	r.GET("/image", func(c *gin.Context) {
		c.File("index.html")
	})
	r.POST("/upload", ImageUpload)
}

func Handler(w http.ResponseWriter, r *http.Request) {
	app.ServeHTTP(w, r)
}
