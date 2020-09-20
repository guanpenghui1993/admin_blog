package admin

import (
	"github.com/gin-gonic/gin"
)

type BaseController struct {
}

// 统一返回json问题
func (*BaseController) Response(c *gin.Context, code int, message string, data interface{}) {
	c.JSON(200, gin.H{"code": code, "message": message, "data": data})
}
