package api

import (
	"context"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"net/http"
	"os"
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

	c.JSON(http.StatusOK, gin.H{"message": "Successfully uploaded file", "filename": key})
}

func init() {

	env := os.Getenv("ENV")
	if env == "" {
		env = "production"
		fmt.Println("Running in %s environment", env)
	} else {
		fmt.Println("Running in %s environment", env)
	}

	initializers.ConnectToDB()

	initializers.SyncDatabase()

	initializers.InitCategories(initializers.DB)
	initializers.InitLocationCategories(initializers.DB)

	app = gin.New()
	r := app.Group("/api")

	routes.UserRoutes(r)
	routes.TravelRoutes(r)
	routes.InviteRoutes(r)
	routes.CategoriesRoutes(r)
	routes.LocationRoutes(r)

	r.POST("/upload", ImageUpload)
}

func Handler(w http.ResponseWriter, r *http.Request) {
	app.ServeHTTP(w, r)
}
