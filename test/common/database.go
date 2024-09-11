package common

import (
	"fmt"
	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"test/model"
)

var DB *gorm.DB

func InitDB() *gorm.DB {
	username := viper.GetString("datasource.username")
	password := viper.GetString("datasource.password")
	host := viper.GetString("datasource.host")
	port := viper.GetString("datasource.port")
	database := viper.GetString("datasource.database")
	charset := viper.GetString("datasource.charset")

	args := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s&parseTime=true",
		username,
		password,
		host,
		port,
		database,
		charset)

	db, err := gorm.Open(mysql.Open(args), &gorm.Config{})
	if err != nil {
		panic("failed to connect database,err: " + err.Error())
	}

	// 自动迁移所有模型
	db.AutoMigrate(&model.User{})
	db.AutoMigrate(&model.Vote{})
	db.AutoMigrate(&model.Candidate{})
	db.AutoMigrate(&model.Participate{})
	db.AutoMigrate(&model.Voted{})
	db.AutoMigrate(&model.Voter{}) // 新增 Voter 模型

	DB = db
	return db
}

// GetDB 定义一个方法获取DB实例
func GetDB() *gorm.DB {
	return DB
}
