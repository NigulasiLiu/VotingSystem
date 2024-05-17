package model

import "gorm.io/gorm"

type Vote struct {
	gorm.Model
	Name     string `gorm:"varchar(50);not null"`
	Deadline string `gorm:"varchar(50);not null"`
	Num      int    `gorm:"default:0;not null"`
	State    int    `gorm:"default:0;not null"`
}
