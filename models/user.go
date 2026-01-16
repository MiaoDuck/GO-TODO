package models

import "gorm.io/gorm"

type User struct {
	gorm.Model        // gorm.Model 自带 ID, CreatedAt, UpdatedAt, DeletedAt
	Username   string `gorm:"unique" json:"username"` // 用户名唯一
	Password   string `json:"-"`                      // json:"-" 表示返回 JSON 时忽略该字段，防止密码泄露！
	Todos      []Todo `json:"todos"`                  // 一对多关联：一个用户有多个 Todo
}