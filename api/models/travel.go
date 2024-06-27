package models

import "gorm.io/gorm"

type Travel struct {
	gorm.Model
	UserIDAdmin   uint
	CategoryID    uint
	Title         string
	Description   string
	Date          string
	Rating        string
	Photo         string
	User          User            `gorm:"foreignKey:UserIDAdmin;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	Location      []Location      `gorm:"foreignKey:TravelID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	Invite        []Invite        `gorm:"foreignKey:TravelID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	Participation []Participation `gorm:"foreignKey:TravelID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	Category      Category        `gorm:"foreignKey:CategoryID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}
