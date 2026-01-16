package common

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// 定义通用的响应结构
type Response struct {
	Code int         `json:"code"` // 业务状态码，200表示成功
	Msg  string      `json:"msg"`  // 提示信息
	Data interface{} `json:"data"` // 数据，用 interface{} 表示可以是任意类型
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