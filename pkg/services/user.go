package services

import (
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
func (u *userService) Login(username, password string) (string, error) {

	user := repository.UserRep.CheckUserByNamePwd(username, password)

	if user.ID > 0 {

		return utils.Token(user.ID)
	}

	return "", nil
}
