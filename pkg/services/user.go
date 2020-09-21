package services

import (
	"watt/pkg/dto"
	"watt/pkg/repository"
	"watt/pkg/utils"
)

type userService struct {
}

var UserService = newUserService()

func newUserService() *userService {
	return new(userService)
}

// 登录
func (u *userService) Login(dto dto.UserLoginDto) (string, error) {

	user := repository.UserRep.UserInfoByName(dto.Username, dto.Password)

	if user.ID > 0 {

		return utils.Token(user.ID)
	}

	return "", nil
}
