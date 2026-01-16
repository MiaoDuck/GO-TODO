package routes

import (
	"go-todo/controllers" // 导入控制器包
	"go-todo/middleware"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
    r := gin.New()
	r.Use(middleware.Logger())
	//初始化Gin引擎
	r.Use(gin.Recovery())

	auth := r.Group("api/v1/auth")
	{
		auth.POST("/register", controllers.Register)
        auth.POST("/login", controllers.Login)	
	}

    v1 := r.Group("/api/v1")//路由分组
	//前缀管理：在这个组下面定义的路由，都会自动带上/api/v1
	//版本控制
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