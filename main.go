package main

import (
	"go-todo/config"
	"go-todo/routes"

	"github.com/spf13/viper"
)

func main() {
	config.InitConfig()      // 先加载配置
	config.ConnectDatabase() // 再连接数据库

	r := routes.SetupRouter()
	
	port := viper.GetString("server.port")
	r.Run(":" + port) // 动态端口
}