package service

import (
	"errors"
	"go-todo/config"
	"go-todo/models"

	"golang.org/x/crypto/bcrypt"
)

type UserService struct{}

func (s *UserService) Register(username, password string) error {
	// 1. 检查用户名是否存在
	var count int64
	config.DB.Model(&models.User{}).Where("username = ?", username).Count(&count)
	if count > 0 {
		return errors.New("用户名已存在")
	}

	// 2. 密码加密 (Hash)
	// Cost 设为 14 左右比较安全，但计算较慢；10 是默认值
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	// 3. 创建用户
	user := models.User{
		Username: username,
		Password: string(hashedPassword), // 存入的是加密后的乱码
	}

	return config.DB.Create(&user).Error
}



func (s *UserService) Login(username, password string) (models.User, error) {
    var user models.User
    // 1. 查找用户是否存在
    if err := config.DB.Where("username = ?", username).First(&user).Error; err != nil {
        return user, errors.New("用户不存在")
    }

    // 2. 比较加密后的密码
    // bcrypt.CompareHashAndPassword 会自动处理盐值并对比
    err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
    if err != nil {
        return user, errors.New("密码错误")
    }

    return user, nil
}