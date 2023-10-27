package models

import "gorm.io/gorm"

type Topic struct {
	gorm.Model
	Name        string     `json:"name" gorm:"not null;index"`
	Description string     `json:"description" gorm:"not null;index"`
	CommunityId uint       `json:"communityId" gorm:"not null;index,foreignkey:CommunityId;references:ID"`
	Comments    []Comments `gorm:"foreignkey:TopicID"`
}
