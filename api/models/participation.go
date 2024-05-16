package models

import "gorm.io/gorm"

type Participation struct {
	gorm.Model
	UserID        uint
	TravelID      uint
	User          User            `gorm:"foreignKey:UserID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	Travel        Travel          `gorm:"foreignKey:TravelID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	FavoritePhoto []FavoritePhoto `gorm:"foreignKey:ParticipationID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}
