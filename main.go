package main

import (
	"go-todo/config"
	"go-todo/routes"

	"github.com/spf13/viper"

	_ "go-todo/docs"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title           Go Todo API
// @version         1.0
// @description     这是一个基于 Gin + GORM 的任务管理后端。
// @host            localhost:8080
// @BasePath        /api/v1

func main() {
	config.InitConfig()      // 先加载配置
	config.ConnectDatabase() // 再连接数据库

	r := routes.SetupRouter()

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	port := viper.GetString("server.port")
	r.Run(":" + port) // 动态端口
}