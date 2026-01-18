package models

// Todo 任务模型
// @Description 任务信息结构体
type Todo struct {
	// 任务 ID
	ID uint `json:"id" gorm:"primaryKey" example:"1"`
	// 任务标题
	Title string `json:"title" example:"完成项目文档"`
	// 任务描述
	Description string `json:"description" example:"编写详细的 README 和 API 文档"`
	// 完成状态：true 完成, false 未完成
	Status bool `json:"status" example:"false"`
	// 所属用户 ID
	UserID uint `json:"user_id" example:"1"`
}

// 注意那个 `json:"title"`
// 这叫做 "Tag" (标签)。
// 它的作用是告诉 Go：把结构体转成 JSON 返回给前端时，这个字段叫 "title" (小写)，而不是 "Title"。