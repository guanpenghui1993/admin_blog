package middleware

import (
	"watt/pkg/utils"

	"github.com/gin-gonic/gin"
)

// 全局检测token有效性
func CheckLogin() gin.HandlerFunc {

	return func(c *gin.Context) {

		code := utils.SUCCESS

		token := c.GetHeader(utils.Setting.Jwt.Header)

		if token == "" {
			code = utils.HEADER_ERROR
		} else {
			_, err := utils.Parse(token)

			if err != nil {
				code = utils.ERROR
			}
		}

		if code != utils.SUCCESS {
			c.JSON(200, utils.Response{code, "令牌错误", nil})
			c.Abort()
			return
		}

		c.Next()
	}
}
