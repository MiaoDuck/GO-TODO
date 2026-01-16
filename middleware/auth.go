package middleware

import (
	"go-todo/common"
	"strings"

	"github.com/gin-gonic/gin"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 1. 获取 Authorization Header
		// 格式通常是: "Bearer <token>"
		tokenString := c.GetHeader("Authorization")

		// 2. 校验格式
		if tokenString == "" || !strings.HasPrefix(tokenString, "Bearer ") {
			common.Error(c, 401, "未登录或非法访问")
			c.Abort() // 🔥 核心：终止请求，不再往下执行
			return
		}

		// 去掉 "Bearer " 前缀
		tokenString = tokenString[7:]

		// 3. 解析 Token
		claims, err := common.ParseToken(tokenString)
		if err != nil {
			common.Error(c, 401, "Token 已过期或无效")
			c.Abort()
			return
		}

		// 4. 🔥 关键点：把解析出来的 UserID 塞进上下文 (Context)
		// 这样后续的 Controller 就能通过 c.Get("userID") 知道是谁在发请求了！
		c.Set("userID", claims.UserID)

		c.Next() // 放行
	}
}