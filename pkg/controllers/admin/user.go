package admin

import (
	"watt/pkg/services"

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
		u.Response(c, 200, "token error", nil)
		return
	}

	if token == "" {
		u.Response(c, 200, "token error", nil)
		return
	}

	u.Response(c, 200, "登录成功", map[string]string{"token": token})
}
