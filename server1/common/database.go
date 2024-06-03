package common

import (
	"fmt"
	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"server1/model"
)

var DB *gorm.DB

func InitDB() *gorm.DB {
	//driverName := viper.GetString("datasource.driverName")

	// 获取命令行参数
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

	db.AutoMigrate(&model.Argument{})

	DB = db
	return db
}

// GetDB 定义一个方法获取DB实例
func GetDB() *gorm.DB {
	return DB
}
