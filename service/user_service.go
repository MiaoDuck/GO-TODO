package service

import (
	"errors"
	"go-todo/common"
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



// Login 登录逻辑
func (s *UserService) Login(username, password string) (string, error) {
	var user models.User
	
	// 1. 根据用户名找用户
	if err := config.DB.Where("username = ?", username).First(&user).Error; err != nil {
		return "", errors.New("用户不存在")
	}

	// 2. 验证密码 (核心！)
	//哪怕你拿到了数据库里的密码 user.Password (是乱码)，你也不能直接 == 对比
	// 必须用 bcrypt.CompareHashAndPassword
	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return "", errors.New("密码错误")
	}

	// 3. 密码正确，生成 JWT Token
	token, err := common.GenerateToken(user.ID)
	if err != nil {
		return "", err
	}

	return token, nil
}