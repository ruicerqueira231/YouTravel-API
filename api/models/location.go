package models

import "gorm.io/gorm"

type Location struct {
	gorm.Model
	TravelID           uint
	LocationCategoryID uint
	Nome               string
	Coordinates        string
	Latitude           string
	Longitude          string
	Travel             Travel           `gorm:"foreignKey:TravelID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	Photos             []Photo          `gorm:"foreignKey:LocationID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	LocationCategory   LocationCategory `gorm:"foreignKey:LocationCategoryID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}
