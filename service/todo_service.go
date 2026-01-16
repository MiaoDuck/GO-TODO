package service

import (
	"go-todo/config"
	"go-todo/models"
)

// 定义一个结构体，方便以后扩展（比如注入不同的 DB）
type TodoService struct{}

func (s *TodoService) GetAll() ([]models.Todo, error) {
    var todos []models.Todo
    err := config.DB.Find(&todos).Error
    return todos, err
}

func (s *TodoService) Create(todo *models.Todo) error {
    return config.DB.Create(todo).Error
}

func (s *TodoService) GetByID(id string) (models.Todo, error) {
    var todo models.Todo
    err := config.DB.First(&todo, id).Error
    return todo, err
}

func (s *TodoService) Update(todo *models.Todo) error {
    return config.DB.Save(todo).Error
}

func (s *TodoService) Delete(id string) error {
    return config.DB.Delete(&models.Todo{}, id).Error
}