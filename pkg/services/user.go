package services

import (
	"crypto/md5"
	"encoding/hex"
	"watt/pkg/models"
)

type userService struct {
}

var UserService = newUserService()

func newUserService() *userService {
	return new(userService)
}
func (u *userService) Create() {
	var user models.User
	models.Link.CreateTable(&user)
}

// 登录
func (u *userService) Login(username, password string) (bool, error) {

	var user models.User

	h := md5.New()

	h.Write([]byte(password))

	cipherStr := h.Sum(nil)

	err := models.Link.Select("uid").Where("username = ? AND password = ?", username, hex.EncodeToString(cipherStr)).First(&user).Error

	if err != nil {
		return false, err
	}

	return true, nil
}
