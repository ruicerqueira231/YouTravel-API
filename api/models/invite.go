package models

import "gorm.io/gorm"

type Invite struct {
	gorm.Model
	UserIDInviter uint `json:"user_id_inviter"`
	UserIDInvited uint `json:"user_id_invited"`
	TravelID      uint `json:"travel_id"`
	Status        string
	Date          string
	Inviter       User   `gorm:"foreignKey:UserIDInviter;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	Invited       User   `gorm:"foreignKey:UserIDInvited;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	Travel        Travel `gorm:"foreignKey:TravelID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}
