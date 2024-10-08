package main

import (
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"os"
	"test/common"
	"test/routes"
)

func main() {
	InitConfig() //在项目启动一开始就读取配置
	db := common.InitDB()
	defer func() {
		sqlDb, _ := db.DB()
		sqlDb.Close()
	}()

	r := gin.Default()
	r = routes.CollectRoute(r)
	port := viper.GetString("server.port")
	if port != "" {
		panic(r.Run(":" + port))
	}
	panic(r.Run()) // 监听并在 0.0.0.0:8080 上启动服务
}

func InitConfig() {
	workDir, _ := os.Getwd()                 //获取当前工作目录
	viper.SetConfigName("application_ea")    //设置要读取的文件名
	viper.SetConfigType("yml")               //设置要读取的文件的类型
	viper.AddConfigPath(workDir + "/config") //设置文件的路径
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}
}
