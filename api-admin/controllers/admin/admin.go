package admin

import (
	"api-admin/helpers"
	"api-admin/services"

	"github.com/gin-gonic/gin"
)

// 测试登录
func Login(c *gin.Context) {

	// username := c.Param("username")

	// password := c.Param("password")

	boolen, err := services.AdminServiceObject.Login("guanpenghui", "123456")

	if err != nil {
		c.JSON(200, helpers.ResponseMap(helpers.ERROR_CODE, 0, "用户名密码不匹配", nil))
		return
	}

	if boolen {
		c.JSON(200, helpers.ResponseMap(helpers.SUCCESS_CODE, 0, "登录成功", boolen))
		return
	}

	c.JSON(200, helpers.ResponseMap(helpers.ERROR_CODE, 0, "登录失败", boolen))
}
