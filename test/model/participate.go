package model

import "gorm.io/gorm"

type Participate struct {
	gorm.Model
	VoteID      uint   `gorm:"default:0;not null"`
	CandidateID uint   `gorm:"default:0;not null"`
	Outs        string `gorm:"varchar(50)"`
}
