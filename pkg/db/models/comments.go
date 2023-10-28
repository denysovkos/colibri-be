package models

import "gorm.io/gorm"

type Comments struct {
	gorm.Model
	Message string `json:"message" gorm:"not null"`
	TopicID uint   `json:"topicId" gorm:"not null;index,foreignkey:TopicID;references:ID"`
	UserID  uint   `json:"userId" gorm:"not null;index,foreignkey:UserID;references:ID"`
	User    User   `gorm:"foreignkey:UserID"`
}
