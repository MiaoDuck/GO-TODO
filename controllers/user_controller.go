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
    // 1. 绑定并校验参数
    if err := c.ShouldBindJSON(&req); err != nil {
        common.Error(c, 400, "参数验证失败: "+err.Error())
        return
    }

    // 2. 调用 Service 进行登录验证并获取 Token
    // 这里的 token 变量接收的就是 service 返回的字符串
    token, err := userService.Login(req.Username, req.Password)
    if err != nil {
        // 登录失败（用户不存在或密码错误）返回 401
        common.Error(c, 401, err.Error())
        return
    }

    // 3. 登录成功，直接把 Token 返回给前端
    common.Success(c, gin.H{
        "token": token,
    })
}