package controllers

import (
	"fmt"
	"go-todo/common" // 导入你定义的通用响应包
	"go-todo/models"
	"go-todo/service"

	"github.com/gin-gonic/gin"
)

// 实例化 service
var todoService = service.TodoService{}

// GetTodos 获取所有任务（支持分页）
// @Summary 获取所有任务
// @Description 获取当前用户的所有任务，支持分页查询
// @Tags Todos
// @Accept json
// @Produce json
// @Param Authorization header string true "Bearer Token"
// @Param page query int false "页码，默认为 1"
// @Param pageSize query int false "每页数量，默认为 10"
// @Success 200 {object} map[string]interface{} "返回任务列表和分页信息"
// @Failure 400 {object} common.Response "请求参数错误"
// @Failure 500 {object} common.Response "服务器错误"
// @Router /todos [get]
func GetTodos(c *gin.Context) {
	userID, _ := c.Get("userID")
	
	// 从查询参数获取分页信息
	page := 1
	pageSize := 10
	
	if p := c.Query("page"); p != "" {
		if _, err := fmt.Sscanf(p, "%d", &page); err != nil {
			common.Error(c, 400, "page 参数格式错误")
			return
		}
	}
	
	if ps := c.Query("pageSize"); ps != "" {
		if _, err := fmt.Sscanf(ps, "%d", &pageSize); err != nil {
			common.Error(c, 400, "pageSize 参数格式错误")
			return
		}
	}
	
	// 调用 service 获取分页数据
	todos, total, err := todoService.GetAll(userID.(uint), page, pageSize)
	if err != nil {
		common.Error(c, 500, "查询失败")
		return
	}
	
	// 返回分页数据
	common.Success(c, gin.H{
		"data":      todos,
		"page":      page,
		"pageSize":  pageSize,
		"total":     total,
	})
}

// CreateTask 创建任务
// @Summary 创建一个新任务
// @Description 创建一个新的任务，需要传递 title 字段，user_id 会自动从 Token 获取
// @Tags Todos
// @Accept json
// @Produce json
// @Param Authorization header string true "Bearer Token"
// @Param todo body models.Todo true "任务信息"
// @Success 201 {object} models.Todo "创建成功"
// @Failure 400 {object} common.Response "请求参数错误"
// @Failure 500 {object} common.Response "服务器错误"
// @Router /todos [post]
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
// @Summary 获取单个任务
// @Description 根据任务 ID 获取单个任务详情
// @Tags Todos
// @Accept json
// @Produce json
// @Param Authorization header string true "Bearer Token"
// @Param id path string true "任务 ID"
// @Success 200 {object} models.Todo "获取成功"
// @Failure 404 {object} common.Response "任务不存在"
// @Failure 500 {object} common.Response "服务器错误"
// @Router /todos/{id} [get]
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
// @Summary 更新任务
// @Description 更新指定 ID 的任务信息
// @Tags Todos
// @Accept json
// @Produce json
// @Param Authorization header string true "Bearer Token"
// @Param id path string true "任务 ID"
// @Param todo body models.Todo true "更新的任务信息"
// @Success 200 {object} models.Todo "更新成功"
// @Failure 400 {object} common.Response "请求参数错误"
// @Failure 404 {object} common.Response "任务不存在"
// @Failure 500 {object} common.Response "服务器错误"
// @Router /todos/{id} [put]
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
// @Summary 删除任务
// @Description 删除指定 ID 的任务
// @Tags Todos
// @Accept json
// @Produce json
// @Param Authorization header string true "Bearer Token"
// @Param id path string true "任务 ID"
// @Success 200 {object} map[string]string "删除成功"
// @Failure 500 {object} common.Response "服务器错误"
// @Router /todos/{id} [delete]
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