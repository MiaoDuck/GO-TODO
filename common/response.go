package common

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Response 统一响应结构
// @Description API 统一响应格式
type Response struct {
	// 业务状态码（200 表示成功，其他表示失败）
	Code int `json:"code" example:"200"`
	// 提示消息
	Msg string `json:"msg" example:"success"`
	// 响应数据（可以是任意类型）
	Data interface{} `json:"data"`
}

// 成功返回
func Success(c *gin.Context, data interface{}) {
	c.JSON(http.StatusOK, Response{
		Code: 200,
		Msg:  "success",
		Data: data,
	})
}

// 失败返回
func Error(c *gin.Context, code int, msg string) {
	c.JSON(http.StatusOK, Response{
		Code: code,
		Msg:  msg,
		Data: nil,
	})
}