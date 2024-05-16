package models

import "gorm.io/gorm"

type LocationCategory struct {
	gorm.Model
	Description string
	Location    []Location `gorm:"foreignKey:LocationCategoryID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}
