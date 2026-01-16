package service

import (
	"go-todo/config"
	"go-todo/models"
)

// 定义一个结构体，方便以后扩展（比如注入不同的 DB）
type TodoService struct{}

func (s *TodoService) GetAll(userID uint) ([]models.Todo, error) {
    var todos []models.Todo
    err := config.DB.Where("user_id = ?", userID).Find(&todos).Error
    return todos, err
}

func (s *TodoService) Create(userID uint, todo *models.Todo) error {
    // 确保设置正确的用户ID
    todo.UserID = userID
    return config.DB.Create(todo).Error
}

func (s *TodoService) GetByID(userID uint, id string) (models.Todo, error) {
    var todo models.Todo
    // 添加 user_id 条件，确保只能访问自己的 todo
    err := config.DB.Where("user_id = ?", userID).First(&todo, id).Error
    return todo, err
}

func (s *TodoService) Update(userID uint, todo *models.Todo) error {
    // 确保 user_id 不被篡改
    todo.UserID = userID
    // 使用 Where 条件确保只能更新自己的 todo
    return config.DB.Where("user_id = ?", userID).Save(todo).Error
}

func (s *TodoService) Delete(userID uint, id string) error {
    // 添加 user_id 条件，确保只能删除自己的 todo
    return config.DB.Where("user_id = ?", userID).Delete(&models.Todo{}, id).Error
}