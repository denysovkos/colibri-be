package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	ID              uint        `gorm:"primaryKey;autoIncrement;unique;index"`
	Email           string      `json:"email" gorm:"unique;not null;index"`
	FirstName       string      `json:"firstName" gorm:"not null"`
	LastName        string      `json:"lastName"`
	NameHandler     string      `json:"nameHandler"`
	Password        string      `json:"password" gorm:"not null"`
	Avatar          string      `json:"avatar"`
	BackgroundImage string      `json:"backgroundImage"`
	Communities     []Community `gorm:"foreignkey:OwnerID;references:ID"`
	Comments        []Comments  `gorm:"foreignkey:UserID;references:ID"`
	Topics          []Topic     `gorm:"foreignkey:OwnerID;references:ID"`
}
