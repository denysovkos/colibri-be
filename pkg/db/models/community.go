package models

import "gorm.io/gorm"

type Community struct {
	gorm.Model
	Name            string `json:"name" gorm:"not null;index"`
	BackgroundImage string `json:"backgroundImage"`
	OwnerID         uint   `json:"ownerId" gorm:"not null;index,foreignkey:OwnerID;references:ID"`
}
