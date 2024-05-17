package model

import "gorm.io/gorm"

type Candidate struct {
	gorm.Model
	Name   string `gorm:"varchar(10);not null"`
	Detail string `gorm:"varchar(200)"`
	Photo  string `gorm:"varchar(50)"`
}
