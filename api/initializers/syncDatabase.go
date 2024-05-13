package initialzers

import "youtravel-api/api/models"

func SyncDatabase() {
	DB.AutoMigrate(&models.User{})
}
