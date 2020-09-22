package services

import (
	"watt/pkg/models"
	"watt/pkg/repository"
	"watt/pkg/utils"
	"watt/pkg/validation"
)

type userService struct {
}

var UserService = newUserService()

func newUserService() *userService {
	return new(userService)
}

// 登录
func (u *userService) Login(dto *validation.UserLogin) (string, error) {

	user := repository.UserRep.UserInfoByName(dto.Username, dto.Password)

	if user.ID > 0 {

		return utils.Token(user.ID)
	}

	return "", nil
}

// 获取用户信息
func (u *userService) Info(uid int) models.User {

	return repository.UserRep.UserInfoById(uid)
}
