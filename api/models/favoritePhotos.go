package models

import "gorm.io/gorm"

type FavoritePhoto struct {
	gorm.Model
	ParticipationID uint
	PhotoID         uint
	Participation   Participation `gorm:"foreignKey:ParticipationID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	Photo           Photo         `gorm:"foreignKey:PhotoID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}
