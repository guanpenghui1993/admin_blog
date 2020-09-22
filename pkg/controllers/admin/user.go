package admin

import (
	"watt/pkg/services"
	"watt/pkg/utils"
	"watt/pkg/validation"

	"github.com/gin-gonic/gin"
)

type UserController struct {
	BaseController
}

// 登录用户
func (u *UserController) Login(c *gin.Context) {

	var userlogin validation.UserLogin

	if err := u.valid(c, &userlogin); err != nil {

		u.json(c, utils.ERROR, err.Error(), nil)

		return
	}

	token, err := services.UserService.Login(&userlogin)

	if err != nil {

		u.json(c, utils.ERROR, "令牌错误", nil)

		return
	}

	if token == "" {

		u.json(c, utils.ERROR, "用户名密码不匹配", nil)

		return
	}

	u.json(c, utils.SUCCESS, "登录成功", map[string]string{"token": token})
}

// 用户信息
func (u *UserController) Info(c *gin.Context) {

	list := services.UserService.Info(u.getuid(c))

	u.json(c, utils.SUCCESS, "获取成功", list)
}
