package service

import (
	"strconv"
	"testing"

	"go-todo/config"
	"go-todo/models"

	"github.com/glebarez/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

// 初始化测试用的数据库
func setupTestDB() *gorm.DB {
	db, _ := gorm.Open(sqlite.Open("file::memory:"), &gorm.Config{
        Logger: logger.Default.LogMode(logger.Silent), 
    })
    db.AutoMigrate(&models.Todo{})
    return db
}

// TestGetAll 测试获取用户的所有 todo（支持分页）
func TestGetAll(t *testing.T) {
	db := setupTestDB()
	config.DB = db
	s := &TodoService{}

	// 为用户 1 创建测试数据
	db.Create(&models.Todo{Title: "任务1", Status: false, UserID: 1})
	db.Create(&models.Todo{Title: "任务2", Status: true, UserID: 1})
	db.Create(&models.Todo{Title: "任务3", Status: false, UserID: 2}) // 其他用户的任务

	// 测试获取用户 1 的所有任务（第1页，每页10条）
	todos, total, err := s.GetAll(1, 1, 10)
	if err != nil {
		t.Errorf("期望没有错误，但得到了: %v", err)
	}

	// 应该只返回 2 条属于用户 1 的任务
	if len(todos) != 2 {
		t.Errorf("期望返回 2 条任务，但得到了 %d 条", len(todos))
	}

	// 验证总数是正确的
	if total != 2 {
		t.Errorf("期望总数是 2，但得到了 %d", total)
	}

	// 验证所有返回的任务都属于用户 1
	for _, todo := range todos {
		if todo.UserID != 1 {
			t.Errorf("期望所有任务的 UserID 都是 1，但得到了 %d", todo.UserID)
		}
	}
}

// TestGetAll_Pagination 测试分页功能
func TestGetAll_Pagination(t *testing.T) {
	db := setupTestDB()
	config.DB = db
	s := &TodoService{}

	// 为用户 1 创建 15 个任务
	for i := 1; i <= 15; i++ {
		title := "task" + strconv.Itoa(i)
		db.Create(&models.Todo{Title: title, Status: false, UserID: 1})
	}

	// 测试第 1 页，每页 10 条
	todos, total, err := s.GetAll(1, 1, 10)
	if err != nil {
		t.Errorf("期望没有错误，但得到了: %v", err)
	}
	if len(todos) != 10 {
		t.Errorf("期望第 1 页返回 10 条，但得到了 %d 条", len(todos))
	}
	if total != 15 {
		t.Errorf("期望总数是 15，但得到了 %d", total)
	}

	// 测试第 2 页，每页 10 条
	todos, total, err = s.GetAll(1, 2, 10)
	if err != nil {
		t.Errorf("期望没有错误，但得到了: %v", err)
	}
	if len(todos) != 5 {
		t.Errorf("期望第 2 页返回 5 条，但得到了 %d 条", len(todos))
	}
	if total != 15 {
		t.Errorf("期望总数是 15，但得到了 %d", total)
	}
}

// TestGetAll_DefaultPageSize 测试分页参数维整
func TestGetAll_DefaultPageSize(t *testing.T) {
	db := setupTestDB()
	config.DB = db
	s := &TodoService{}

	// 为用户 1 创建 25 个任务
	for i := 1; i <= 25; i++ {
		title := "task" + strconv.Itoa(i)
		db.Create(&models.Todo{Title: title, Status: false, UserID: 1})
	}

	// 测试画幅参数（page < 1 时默认至 1，pageSize < 1 时默认至10）
	todos, total, err := s.GetAll(1, 0, 0)
	if err != nil {
		t.Errorf("期望没有错误，但得到了: %v", err)
	}
	if len(todos) != 10 {
		t.Errorf("期望默认每页 10 条，但得到了 %d 条", len(todos))
	}
	if total != 25 {
		t.Errorf("期望总数是 25，但得到了 %d", total)
	}
}

// TestCreate 测试创建 todo
func TestCreate(t *testing.T) {
	db := setupTestDB()
	config.DB = db
	s := &TodoService{}

	// 创建一个新的 todo
	todo := &models.Todo{Title: "新任务", Status: false}
	err := s.Create(1, todo)
	if err != nil {
		t.Errorf("期望没有错误，但得到了: %v", err)
	}

	// 验证 UserID 被正确设置
	if todo.UserID != 1 {
		t.Errorf("期望 UserID 是 1，但得到了 %d", todo.UserID)
	}

	// 验证数据确实被保存到数据库
	var savedTodo models.Todo
	db.First(&savedTodo, todo.ID)
	if savedTodo.Title != "新任务" {
		t.Errorf("期望保存的任务标题是 '新任务'，但得到了 '%s'", savedTodo.Title)
	}
}

// TestGetByID 测试获取单个 todo
func TestGetByID(t *testing.T) {
	db := setupTestDB()
	config.DB = db
	s := &TodoService{}

	// 为用户 1 创建一个任务
	todo := &models.Todo{Title: "任务1", Status: false, UserID: 1}
	db.Create(todo)

	// 为用户 2 创建一个任务
	todo2 := &models.Todo{Title: "任务2", Status: false, UserID: 2}
	db.Create(todo2)

	// 用户 1 可以获取自己的任务
	retrievedTodo, err := s.GetByID(1, toString(todo.ID))
	if err != nil {
		t.Errorf("期望没有错误，但得到了: %v", err)
	}
	if retrievedTodo.Title != "任务1" {
		t.Errorf("期望任务标题是 '任务1'，但得到了 '%s'", retrievedTodo.Title)
	}

	// 用户 1 不能获取用户 2 的任务
	_, err = s.GetByID(1, toString(todo2.ID))
	if err == nil {
		t.Error("期望用户 1 无法访问用户 2 的任务，但没有返回错误")
	}
}

// TestUpdate 测试更新 todo
func TestUpdate(t *testing.T) {
	db := setupTestDB()
	config.DB = db
	s := &TodoService{}

	// 为用户 1 创建一个任务
	todo := &models.Todo{Title: "原始标题", Status: false, UserID: 1}
	db.Create(todo)

	// 更新任务
	todo.Title = "新标题"
	todo.Status = true
	err := s.Update(1, todo)
	if err != nil {
		t.Errorf("期望没有错误，但得到了: %v", err)
	}

	// 验证数据库中的数据已更新
	var updatedTodo models.Todo
	db.First(&updatedTodo, todo.ID)
	if updatedTodo.Title != "新标题" {
		t.Errorf("期望更新后的标题是 '新标题'，但得到了 '%s'", updatedTodo.Title)
	}
	if updatedTodo.Status != true {
		t.Error("期望更新后的状态是 true，但得到了 false")
	}
}

// TestUpdate_UserIDProtection 测试更新时 UserID 无法被篡改
func TestUpdate_UserIDProtection(t *testing.T) {
	db := setupTestDB()
	config.DB = db
	s := &TodoService{}

	// 为用户 1 创建一个任务
	todo := &models.Todo{Title: "任务", Status: false, UserID: 1}
	db.Create(todo)

	// 尝试将任务的 UserID 改为 2
	todo.UserID = 2
	err := s.Update(1, todo)
	if err != nil {
		t.Errorf("期望没有错误，但得到了: %v", err)
	}

	// 验证 UserID 没有被改动（应该仍然是 1）
	var updatedTodo models.Todo
	db.First(&updatedTodo, todo.ID)
	if updatedTodo.UserID != 1 {
		t.Errorf("期望 UserID 仍然是 1，但得到了 %d（UserID 不应该被篡改）", updatedTodo.UserID)
	}
}

// TestDelete 测试删除 todo
func TestDelete(t *testing.T) {
	db := setupTestDB()
	config.DB = db
	s := &TodoService{}

	// 为用户 1 创建一个任务
	todo := &models.Todo{Title: "待删除任务", Status: false, UserID: 1}
	db.Create(todo)

	// 删除任务
	err := s.Delete(1, toString(todo.ID))
	if err != nil {
		t.Errorf("期望没有错误，但得到了: %v", err)
	}

	// 验证任务已被删除
	var deletedTodo models.Todo
	result := db.First(&deletedTodo, todo.ID)
	if result.Error == nil {
		t.Error("期望任务已被删除，但仍然存在于数据库中")
	}
}

// TestDelete_UserIDProtection 测试删除时用户隔离
func TestDelete_UserIDProtection(t *testing.T) {
	db := setupTestDB()
	config.DB = db
	s := &TodoService{}

	// 为用户 1 创建一个任务
	todo := &models.Todo{Title: "用户1的任务", Status: false, UserID: 1}
	db.Create(todo)

	// 用户 2 尝试删除用户 1 的任务
	err := s.Delete(2, toString(todo.ID))
	if err != nil {
		t.Errorf("期望没有错误（因为 Where 条件不匹配，不会影响任何行），但得到了: %v", err)
	}

	// 验证任务仍然存在
	var existingTodo models.Todo
	result := db.First(&existingTodo, todo.ID)
	if result.Error != nil {
		t.Error("期望任务仍然存在，但已被删除")
	}
}

// 辅助函数：uint 转 string
func toString(id uint) string {
    return strconv.FormatUint(uint64(id), 10)
}