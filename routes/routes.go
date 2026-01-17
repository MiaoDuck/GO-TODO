package routes

import (
	"go-todo/controllers" // 导入控制器包
	"go-todo/middleware"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
    r := gin.New()
	r.Use(middleware.Logger())
	r.Use(middleware.Cors())
	//初始化Gin引擎
	r.Use(gin.Recovery())
	// 捕获和处理运行时发生的 panic 错误，防止程序因未捕获的 panic 而崩溃，
	// 转而返回一个标准的 HTTP 500 错误响应给客户端，常用于生产环境确保应用的健壮性。 

	//公开接口（注册 登录）
	auth := r.Group("/api/v1/auth")
	{
		auth.POST("/register", controllers.Register)
        auth.POST("/login", controllers.Login)	
	}

    v1 := r.Group("/api/v1")//路由分组
	//前缀管理：在这个组下面定义的路由，都会自动带上/api/v1
	//版本控制
	v1.Use(middleware.AuthMiddleware())
    {
        // 这里的 controllers.GetTodos 对应上面定义的函数
        v1.POST("/todos", controllers.CreateTask)
		v1.GET("/todos", controllers.GetTodos)
		v1.GET("/todos/:id", controllers.GetTodo)     // 查询单个
    	v1.DELETE("/todos/:id", controllers.DeleteTodo) // 删除
		v1.PUT("/todos/:id",controllers.UpdateTodo)

    }

    return r
}