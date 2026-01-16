package controllers

import (
	"go-todo/common" // 导入你定义的通用响应包
	"go-todo/models"
	"go-todo/service"

	"github.com/gin-gonic/gin"
)

// 实例化 service
var todoService = service.TodoService{}

// GetTodos 获取所有任务
func GetTodos(c *gin.Context) {
	userID, _ := c.Get("userID")
	todos, err := todoService.GetAll(userID.(uint))
	if err != nil {
		common.Error(c, 500, "查询失败")
		return
	}
	common.Success(c, todos)
}

// CreateTask 创建任务
func CreateTask(c *gin.Context) {
	userID, _ := c.Get("userID")
	var todo models.Todo
	if err := c.ShouldBindJSON(&todo); err != nil {
		common.Error(c, 400, err.Error())
		return
	}

	if err := todoService.Create(userID.(uint), &todo); err != nil {
		common.Error(c, 500, "创建失败")
		return
	}
	common.Success(c, todo)
}

// GetTodo 获取单个任务
func GetTodo(c *gin.Context) {
	userID, _ := c.Get("userID")
	id := c.Param("id")
	todo, err := todoService.GetByID(userID.(uint), id)
	if err != nil {
		common.Error(c, 404, "任务没找到")
		return
	}
	common.Success(c, todo)
}

// UpdateTodo 更新任务
func UpdateTodo(c *gin.Context) {
	userID, _ := c.Get("userID")
	id := c.Param("id")
	// 1. 先查是否存在
	todo, err := todoService.GetByID(userID.(uint), id)
	if err != nil {
		common.Error(c, 404, "找不到该任务")
		return
	}

	// 2. 绑定新数据
	if err := c.ShouldBindJSON(&todo); err != nil {
		common.Error(c, 400, "参数格式错误")
		return
	}

	// 3. 调用 Service 更新
	if err := todoService.Update(userID.(uint), &todo); err != nil {
		common.Error(c, 500, "更新失败")
		return
	}
	common.Success(c, todo)
}

// DeleteTodo 删除任务
func DeleteTodo(c *gin.Context) {
	userID, _ := c.Get("userID")
	id := c.Param("id")
	if err := todoService.Delete(userID.(uint), id); err != nil {
		common.Error(c, 500, "删除失败")
		return
	}
	// 删除成功也可以返回一个简单的 map 或者 null
	common.Success(c, gin.H{"id": id})
}