package model

import "gorm.io/gorm"

type Voted struct {
	gorm.Model
	VoteID uint `gorm:"default:0;not null"`
	UserID uint `gorm:"default:0;not null"`
}
