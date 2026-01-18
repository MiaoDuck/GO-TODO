package controllers

import (
	"go-todo/common"
	"go-todo/service"

	"github.com/gin-gonic/gin"
)

var userService = service.UserService{}

// AuthRequest 认证请求
// @Description 用户登录和注册请求结构体
type AuthRequest struct {
	// 用户名
	Username string `json:"username" binding:"required" example:"john_doe"`
	// 密码
	Password string `json:"password" binding:"required" example:"password123"`
}

// Register 用户注册
// @Summary 用户注册
// @Description 新用户注册，需要提供用户名和密码
// @Tags Auth
// @Accept json
// @Produce json
// @Param request body AuthRequest true "注册请求信息"
// @Success 200 {object} common.Response "注册成功"
// @Failure 400 {object} common.Response "参数验证失败或用户已存在"
// @Failure 500 {object} common.Response "服务器错误"
// @Router /auth/register [post]
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
// @Summary 用户登录
// @Description 用户使用用户名和密码登录，返回 JWT 令牌
// @Tags Auth
// @Accept json
// @Produce json
// @Param request body AuthRequest true "登录请求信息"
// @Success 200 {object} map[string]string "登录成功，返回 JWT 令牌"
// @Failure 400 {object} common.Response "参数验证失败"
// @Failure 401 {object} common.Response "用户不存在或密码错误"
// @Failure 500 {object} common.Response "服务器错误"
// @Router /auth/login [post]
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