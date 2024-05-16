package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Nome          string
	Username      string `gorm:"unique"`
	Email         string `gorm:"unique"`
	Photo         string
	Password      string
	Age           int
	Nationality   string
	Travel        []Travel        `gorm:"foreignKey:UserIDAdmin;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	Inviter       []Invite        `gorm:"foreignKey:UserIDInviter;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	Invited       []Invite        `gorm:"foreignKey:UserIDInvited;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	Participation []Participation `gorm:"foreignKey:UserID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}
