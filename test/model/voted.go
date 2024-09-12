package model

import "gorm.io/gorm"

type Voted struct {
	gorm.Model
	VoteID    uint   `gorm:"default:0;not null"`
	VoteKey   string `gorm:"default:0;not null"`
	VoteIndex uint   `gorm:"default:0;not null"`
}
