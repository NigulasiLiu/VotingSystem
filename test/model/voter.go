package model

import "gorm.io/gorm"

// Voter 定义投票者模型
type Voter struct {
	gorm.Model
	Name   string `gorm:"type:varchar(100);not null"` // 姓名字段
	Phone  string `gorm:"type:varchar(15);not null"`  // 电话字段
	Email  string `gorm:"type:varchar(255);not null"` // 邮箱字段
	Key    string `gorm:"type:varchar(255);not null"` // 新增密钥字段
	VoteID uint   `gorm:"not null"`                   // 外键，指向 votes 表
}
