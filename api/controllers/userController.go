package controllers

import (
	"context"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	"log"
	"net/http"
	"os"
	"time"
	initialzers "youtravel-api/api/initializers"
	"youtravel-api/api/models"
)

func Signup(c *gin.Context) {
	var body struct {
		Nome        string `json:"nome"`
		Username    string `json:"username"`
		Email       string `json:"email"`
		Photo       string `json:"photo"`
		Password    string `json:"password"`
		Age         int    `json:"age"`
		Nationality string `json:"nationality"`
	}

	if err := c.Bind(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to read body",
		})
		return
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(body.Password), 10)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to hash the password",
		})
		return
	}

	defaultImage := "default_image.png"
	if body.Photo == "" {
		body.Photo = defaultImage
	}

	user := models.User{
		Nome:        body.Nome,
		Username:    body.Username,
		Email:       body.Email,
		Photo:       body.Photo,
		Password:    string(hash),
		Age:         body.Age,
		Nationality: body.Nationality,
	}

	if initialzers.DB == nil {
		log.Fatal("Database connection is not initialized")
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Database is not initialized"})
		return
	}

	result := initialzers.DB.Create(&user)
	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to create user",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "User created successfully",
	})
}

func Login(c *gin.Context) {
	var body struct {
		Email    string
		Password string
	}

	if err := c.Bind(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to read body"})
		return
	}

	// Look for the user in the database
	var user models.User
	initialzers.DB.First(&user, "email = ?", body.Email)
	if user.ID == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid email or password"})
		return
	}

	// Compare the password
	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(body.Password))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Password"})
		return
	}

	// Generate a token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": user.ID,
		"exp": time.Now().Add(time.Hour * 24 * 30).Unix(),
	})
	tokenString, err := token.SignedString([]byte(os.Getenv("SECRET_KEY")))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to create token"})
		return
	}

	isProduction := os.Getenv("ENV") == "production"
	secure := isProduction
	httpOnly := true

	c.SetSameSite(http.SameSiteLaxMode)
	c.SetCookie("Authorization", tokenString, 3600*24*30, "/", "", secure, httpOnly)

	c.JSON(http.StatusOK, gin.H{"message": tokenString})
}

func Validate(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "I'm logged in",
	})
}

func Logout(c *gin.Context) {
	c.SetCookie("Authorization", "", -1, "", "", false, true)

	c.JSON(http.StatusOK, gin.H{
		"message": "Logged out successfully",
	})
}

func GetAllUsers(c *gin.Context) {

	var users []models.User

	result := initialzers.DB.Find(&users)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to fetch users",
		})
		return
	}

	c.JSON(http.StatusOK, users)
}

func GetUserById(c *gin.Context) {

	id := c.Param("id")

	var user models.User

	if err := initialzers.DB.First(&user, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusNotFound, gin.H{
				"error": "User not found",
			})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": "Failed to fech user",
			})
		}
		return
	}

	c.JSON(http.StatusOK, user)
}

func GetUserByEmail(c *gin.Context) {

	var body struct {
		Email string `json:"email"`
	}

	if err := c.BindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	var user models.User

	if err := initialzers.DB.Where("email = ?", body.Email).First(&user).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusNotFound, gin.H{
				"error": "User not found",
			})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": "Failed to fetch user",
			})
		}
		return
	}

	c.JSON(http.StatusOK, user)
}

func GetUserByUsername(c *gin.Context) {

	var body struct {
		Username string `json:"username"`
	}

	if err := c.BindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	var user models.User

	if err := initialzers.DB.Where("username = ?", body.Username).First(&user).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusNotFound, gin.H{
				"error": "User not found",
			})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": "Failed to fetch user",
			})
		}
		return
	}
	c.JSON(http.StatusOK, user)
}

func UpdateUser(c *gin.Context) {
	userID := c.Param("id")
	if userID == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Missing user ID"})
		return
	}

	var body struct {
		Nome        *string `json:"nome,omitempty"`
		Username    *string `json:"username,omitempty"`
		Email       *string `json:"email,omitempty"`
		Photo       *string `json:"photo,omitempty"`
		Password    *string `json:"password,omitempty"`
		Age         *int    `json:"age,omitempty"`
		Nationality *string `json:"nationality,omitempty"`
	}

	if err := c.BindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	var user models.User
	result := initialzers.DB.First(&user, userID)
	if result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	if body.Nome != nil {
		user.Nome = *body.Nome
	}
	if body.Username != nil {
		user.Username = *body.Username
	}
	if body.Email != nil {
		user.Email = *body.Email
	}
	if body.Photo != nil {
		user.Photo = *body.Photo
	}
	if body.Password != nil {
		hash, err := bcrypt.GenerateFromPassword([]byte(*body.Password), 10)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to hash the password"})
			return
		}
		user.Password = string(hash)
	}
	if body.Age != nil {
		user.Age = *body.Age
	}
	if body.Nationality != nil {
		user.Nationality = *body.Nationality
	}

	if result = initialzers.DB.Save(&user); result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update user"})
		return
	}

	c.JSON(http.StatusOK, user)
}

func GetUserImage(c *gin.Context) {
	userID := c.Param("id")
	if userID == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Missing user ID"})
		return
	}

	var user models.User
	result := initialzers.DB.First(&user, userID)
	if result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	if user.Photo == "" {
		c.JSON(http.StatusNotFound, gin.H{"error": "No photo available for this user"})
		return
	}

	cfg, err := config.LoadDefaultConfig(context.TODO())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to load AWS config"})
		return
	}

	client := s3.NewFromConfig(cfg)
	resp, err := client.GetObject(context.TODO(), &s3.GetObjectInput{
		Bucket: aws.String("you-travel-storage"),
		Key:    aws.String(user.Photo),
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve file from S3"})
		return
	}
	defer resp.Body.Close()

	contentType := "application/octet-stream"
	if resp.ContentType != nil {
		contentType = *resp.ContentType
	}
	c.DataFromReader(http.StatusOK, *resp.ContentLength, contentType, resp.Body, nil)
}
