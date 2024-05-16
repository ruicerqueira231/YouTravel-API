package models

import "gorm.io/gorm"

type Photo struct {
	gorm.Model
	URL           string
	Description   string
	LocationID    uint
	Location      Location        `gorm:"foreignKey:LocationID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	FavoritePhoto []FavoritePhoto `gorm:"foreignKey:PhotoID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}
