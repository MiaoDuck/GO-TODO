package common

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

// 定义 Token 的秘钥 (生产环境应该从 Config 读取，这里先演示)
// ⚠️ 注意：这个 key 绝对不能泄露，一旦泄露，别人就能伪造身份
var jwtKey = []byte("my_secret_key_todo_app") 

// 自定义 Claims (载荷)，这里我们只存 UserID
type MyCustomClaims struct {
	UserID uint `json:"user_id"`
	jwt.RegisteredClaims
}

// 1. 生成 Token
func GenerateToken(userID uint) (string, error) {
	// 设置有效期，比如 24 小时
	expirationTime := time.Now().Add(24 * time.Hour)

	claims := &MyCustomClaims{
		UserID: userID,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
			Issuer:    "go-todo",
		},
	}

	// 使用 HS256 算法签名
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtKey)
}

// 2. 解析 Token
func ParseToken(tokenString string) (*MyCustomClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &MyCustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})

	if err != nil {
		return nil, err
	}

	// 验证 token 是否有效，并转换成我们自定义的 Claims
	if claims, ok := token.Claims.(*MyCustomClaims); ok && token.Valid {
		return claims, nil
	}

	return nil, errors.New("invalid token")
}