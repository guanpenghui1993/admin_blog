package controllers

import (
	"api-admin/helpers"
	"api-admin/services"

	"github.com/gin-gonic/gin"
)

// 测试登录
func Login(c *gin.Context) {

	// username := c.Param("username")
	// password := c.Param("password")

	boolen := services.AdminServiceObject.Login("guanpenghui", "123456")

	if boolen {
		c.JSON(200, helpers.ResponseMap(helpers.SUCCESS_CODE, 0, "测试登录成功", boolen))
		return
	}

	c.JSON(200, helpers.ResponseMap(helpers.ERROR_CODE, 0, "测试登录失败", boolen))
}
