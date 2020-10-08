package middleware

import (
	"watt/pkg/services"
	"watt/pkg/utils"

	"github.com/gin-gonic/gin"
)

// 全局检测token有效性
func CheckLogin() gin.HandlerFunc {

	return func(c *gin.Context) {

		code := utils.SUCCESS

		msg := ""

		token := c.GetHeader(utils.Setting.Jwt.Header)

		if token == "" {

			code = utils.HEADER_ERROR
			msg = "缺少token信息"

		} else {

			uid, err := utils.Parse(token)

			if err != nil {
				code = utils.ERROR
				msg = "token error"
			} else {
				if list := services.UserService.Info(uid); list.ID <= 0 {
					msg = "用户不存在或已禁用"
					code = utils.ERROR
				}
			}
		}

		if code != utils.SUCCESS {
			c.JSON(200, utils.Response{code, msg, nil})
			c.Abort()
			return
		}

		c.Next()
	}
}

// 全局异常信息
func Recovery() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if r := recover(); r != nil {
				utils.Error(r)
				c.JSON(200, utils.Response{utils.SERVER_ERROR, "服务器异常，请稍后再试", nil})
				c.Abort()
				return
			}
		}()
		c.Next()
	}
}

// 当前用户操作权限（权限菜单路由权限）
