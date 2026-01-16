package middleware

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
)

// Logger 是一个中间件函数
func Logger() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 1. 请求开始时间
		startTime := time.Now()

		// 2. 处理请求 (这里会跳转到下一个中间件或 Controller)
		c.Next()

		// 3. 请求结束时间
		endTime := time.Now()
		
		// 4. 计算耗时
		latency := endTime.Sub(startTime)

		// 5. 获取请求信息
		method := c.Request.Method
		uri := c.Request.RequestURI
		statusCode := c.Writer.Status()
		clientIP := c.ClientIP()

		// 打印日志
		// 格式：[状态码] | 耗时 | IP | 方法 | 路径
		fmt.Printf("| %3d | %13v | %15s | %s | %s |\n",
			statusCode,
			latency,
			clientIP,
			method,
			uri,
		)
	}
}