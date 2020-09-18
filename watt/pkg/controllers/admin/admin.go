package admin

import (
	"watt/pkg/services"

	"github.com/gin-gonic/gin"
)

// 测试登录
func Login(c *gin.Context) {

	// username := c.Param("username")

	// password := c.Param("password")

	boolen, err := services.AdminServiceObject.Login("guanpenghui", "123456")

	if err != nil {
		c.JSON(200, gin.H{"message": "用户名密码不匹配"})
		return
	}

	if boolen {
		c.JSON(200, gin.H{"message": "登录成功"})
		return
	}

	c.JSON(200, gin.H{"message": "登录失败"})
}
