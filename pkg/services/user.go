package services

import (
	"errors"
	"html"
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

	return "", errors.New("用户名密码不匹配")
}

// 获取用户信息
func (u *userService) Info(uid int) models.User {

	return repository.UserRep.UserInfoById(uid)
}

// 获取用户列表
func (u *userService) UserList(param *validation.BaseValid) []models.User {

	return repository.UserRep.UserList(param.Page, param.Size)
}

// 删除某个用户
func (u *userService) Delete(uid *validation.BaseID) bool {

	return repository.UserRep.DeleteUser(uid.ID)
}

// 新增用户
func (u *userService) Insert(param *validation.InsertUserData) error {

	// html 标签转义
	param.Username = html.EscapeString(param.Username)

	param.Nickname = html.EscapeString(param.Nickname)

	param.Password = utils.Md5(param.Password)

	return repository.UserRep.InsertUser(param)
}

// 修改用户状态
