package controllers

import (
	"go-todo/common"
	"go-todo/service"

	"github.com/gin-gonic/gin"
)

var userService = service.UserService{}

// 定义登录/注册用的请求结构体
type AuthRequest struct {
    Username string `json:"username" binding:"required"`
    Password string `json:"password" binding:"required"`
}

// Register 用户注册
func Register(c *gin.Context) {
    var req AuthRequest
    if err := c.ShouldBindJSON(&req); err != nil {
        common.Error(c, 400, "参数验证失败")
        return
    }

    if err := userService.Register(req.Username, req.Password); err != nil {
        common.Error(c, 400, err.Error())
        return
    }

    common.Success(c, "注册成功")
}

// Login 用户登录
func Login(c *gin.Context) {
    var req AuthRequest
    if err := c.ShouldBindJSON(&req); err != nil {
        common.Error(c, 400, "参数验证失败")
        return
    }

    user, err := userService.Login(req.Username, req.Password)
    if err != nil {
        common.Error(c, 401, err.Error())
        return
    }

    // TODO: 这里需要生成 JWT Token 返回给前端
    // 目前我们先返回用户信息测试
    common.Success(c, gin.H{
        "user_id":  user.ID,
        "username": user.Username,
        "token":    "这里之后替换为生成的JWT-TOKEN", 
    })
}