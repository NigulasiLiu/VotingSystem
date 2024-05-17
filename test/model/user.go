package model

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Telephone string `gorm:"varchar(11);not null;unique"`
	Password  string `gorm:"varchar(50);not null"`
	Role      uint   `gorm:"default:0;not null"`
}
