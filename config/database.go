package config

import (
	"fmt"
	"go-todo/models"

	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DB *gorm.DB

func InitConfig() {
	viper.SetConfigName("config") // 配置文件名 (不带后缀)
	viper.SetConfigType("yaml")   // 配置文件类型
	viper.AddConfigPath(".")      // 查找当前目录

	if err := viper.ReadInConfig(); err != nil {
		panic(fmt.Errorf("fatal error config file: %w", err))
	}
}

func ConnectDatabase() {
	// 从 Viper 获取配置
	username := viper.GetString("database.username")
	password := viper.GetString("database.password")
	host := viper.GetString("database.host")
	port := viper.GetInt("database.port")
	dbname := viper.GetString("database.dbname")

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		username, password, host, port, dbname)

	database, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})

	if err != nil {
		fmt.Printf("数据库连接失败详情: %v\n", err)
		panic("🔥 无法连接数据库！")
	}

	err = database.AutoMigrate(&models.User{},&models.Todo{})
    
    if err != nil {
        fmt.Printf("自动迁移失败: %v\n", err)
    }

    DB = database
    fmt.Println("✅ 数据库连接成功，表结构同步完成！")
}