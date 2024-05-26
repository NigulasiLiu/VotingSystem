package model

import "gorm.io/gorm"

type Argument struct {
	gorm.Model
	VoteID   uint   `gorm:"default:0;not null"`
	VW0      []byte `gorm:"type:blob;size:-"`
	VW1      []byte `gorm:"type:blob;size:-"`
	Outputs0 string `gorm:"type:json;size:-"` // Changed to string for JSON storage
	Outputs1 string `gorm:"type:json;size:-"` // Changed to string for JSON storage
}
