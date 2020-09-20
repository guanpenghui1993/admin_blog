package repository

import (
	"crypto/md5"
	"encoding/hex"
	"watt/pkg/models"
)

type UserRepository struct {
}

var UserRep = newUserRepository()

func newUserRepository() *UserRepository {
	return new(UserRepository)
}

// 获取用户信息
func (u *UserRepository) CheckUserByNamePwd(username, password string) models.User {

	var user models.User

	h := md5.New()

	h.Write([]byte(password))

	cipherStr := h.Sum(nil)

	models.Link.Where("username = ? AND password = ? AND status = 1", username, hex.EncodeToString(cipherStr)).First(&user)

	return user
}
