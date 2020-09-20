package admin

import (
	"watt/pkg/services"
	"watt/pkg/utils"

	"github.com/gin-gonic/gin"
)

type UserController struct {
	BaseController
}

// 登录用户
func (u *UserController) Login(c *gin.Context) {

	// username := c.Param("username")

	// password := c.Param("password")
	token, err := services.UserService.Login("guanpenghui", "123456")

	if err != nil {
		u.Response(c, utils.ERROR, "令牌错误", nil)
		return
	}

	if token == "" {
		u.Response(c, utils.ERROR, "验证失败", nil)
		return
	}

	u.Response(c, utils.SUCCESS, "登录成功", map[string]string{"token": token})
}
