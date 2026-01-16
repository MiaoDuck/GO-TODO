package models

// Todo 结构体，对应数据库中的 todos 表
type Todo struct {
	ID     uint   `json:"id" gorm:"primaryKey"` // 主键
	Title  string `json:"title"`                // 任务标题
	Status bool   `json:"status"`               // 完成状态：true完成, false未完成

	// 🔥 新增：外键关联
	UserID uint `json:"user_id"` // 属于哪个用户
}

// 注意那个 `json:"title"`
// 这叫做 "Tag" (标签)。
// 它的作用是告诉 Go：把结构体转成 JSON 返回给前端时，这个字段叫 "title" (小写)，而不是 "Title"。