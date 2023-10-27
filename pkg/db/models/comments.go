package models

import "gorm.io/gorm"

type Comments struct {
	gorm.Model
	Message string `json:"message" gorm:"not null"`
	UserID  uint   `json:"userId" gorm:"not null;index,foreignkey:UserID;references:ID"`
	TopicID uint   `json:"topicId" gorm:"not null;index,foreignkey:TopicID;references:ID"`
}
