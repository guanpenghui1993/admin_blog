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

		u.json(c, utils.ERROR, err.Error(), nil)

		return
	}

	u.json(c, utils.SUCCESS, "登录成功", map[string]string{"token": token})
}

// 用户信息
func (u *UserController) Info(c *gin.Context) {

	list := services.UserService.Info(u.getuid(c))

	u.json(c, utils.SUCCESS, "获取成功", list)
}

// 用户列表
func (u *UserController) UserList(c *gin.Context) {

	var param validation.BaseValid

	if err := u.valid(c, &param); err != nil {

		u.json(c, utils.ERROR, err.Error(), nil)

		return
	}

	data := services.UserService.UserList(&param)

	u.json(c, utils.SUCCESS, "获取成功", data)
}

// 删除用户
func (u *UserController) UserDel(c *gin.Context) {

	var param validation.BaseID

	if err := u.valid(c, &param); err != nil {

		u.json(c, utils.ERROR, err.Error(), nil)

		return
	}

	if services.UserService.Delete(&param) {
		u.json(c, utils.SUCCESS, "删除成功", nil)
		return
	}

	u.json(c, utils.ERROR, "删除失败", nil)
}

// 添加用户
func (u *UserController) UserAdd(c *gin.Context) {

	var insertData validation.InsertUserData

	if err := u.valid(c, &insertData); err != nil {

		u.json(c, utils.ERROR, err.Error(), nil)

		return
	}

	if err := services.UserService.Insert(&insertData); err != nil {
		u.json(c, utils.ERROR, err.Error(), nil)
		return
	}

	u.json(c, utils.SUCCESS, "添加成功", nil)
}

// 更新用户状态
func (u *UserController) UserUpdate(c *gin.Context) {

	var updateData validation.BaseIdStatus

	if err := u.valid(c, &updateData); err != nil {

		u.json(c, utils.ERROR, err.Error(), nil)

		return
	}

	if err := services.UserService.Update(&updateData); err != nil {
		u.json(c, utils.ERROR, err.Error(), nil)
		return
	}

	u.json(c, utils.SUCCESS, "更新成功", nil)
}
