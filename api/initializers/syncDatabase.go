package initialzers

import "youtravel-api/api/models"

func SyncDatabase() {
	DB.AutoMigrate(&models.User{},
		&models.Travel{},
		&models.Location{},
		&models.Photo{},
		&models.Invite{},
		&models.Participation{},
		&models.Category{},
		&models.LocationCategory{},
		&models.FavoritePhoto{})
}
