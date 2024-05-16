package models

import "gorm.io/gorm"

type Category struct {
	gorm.Model
	Description string
	Travel      []Travel `gorm:"foreignKey:CategoryID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}
