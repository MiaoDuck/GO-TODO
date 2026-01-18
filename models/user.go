package models

import "gorm.io/gorm"

// User 用户模型
// @Description 用户信息结构体
type User struct {
	// 数据库标准字段：ID, CreatedAt, UpdatedAt, DeletedAt
	gorm.Model
	// 用户名，全局唯一
	Username string `gorm:"unique" json:"username" example:"john_doe"`
	// 密码（JSON 返回时忽略，防止泄露）
	Password string `json:"-"`
	// 该用户的所有任务
	Todos []Todo `json:"todos"`
}