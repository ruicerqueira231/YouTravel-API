package initialzers

import (
	"log"
	"youtravel-api/api/models"

	"gorm.io/gorm"
)

var categories = []string{
	"Business",
	"Lounge",
	"Culture",
}

var locationCategories = []string{
	"Nature",
	"Urban",
	"Cultural",
	"Recreation",
	"Food and Drink",
	"Education",
	"Transportation",
	"Health and Wellness",
	"Business",
	"Residential",
}

func InitCategories(db *gorm.DB) {
	for _, desc := range categories {
		var category models.Category
		result := db.Where("description = ?", desc).First(&category)

		if result.Error != nil && result.Error == gorm.ErrRecordNotFound {
			category = models.Category{Description: desc}
			if err := db.Create(&category).Error; err != nil {
				log.Printf("Error creating category %s: %v", desc, err)
			} else {
				log.Printf("Category %s created successfully", desc)
			}
		}
	}
}

func InitLocationCategories(db *gorm.DB) {
	for _, desc := range locationCategories {
		var category models.LocationCategory
		result := db.Where("description = ?", desc).First(&category)

		if result.Error != nil && result.Error == gorm.ErrRecordNotFound {
			category = models.LocationCategory{Description: desc}
			if err := db.Create(&category).Error; err != nil {
				log.Printf("Error creating category %s: %v", desc, err)
			} else {
				log.Printf("Category %s created successfully", desc)
			}
		}
	}
}
