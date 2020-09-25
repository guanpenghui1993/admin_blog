package admin

import "github.com/gin-gonic/gin"

type AccessController struct {
	BaseController
}

// 获取权限列表
func (a *AccessController) List(c *gin.Context) {

	// var param validation.BaseValid

}

// 设置权限
func (a *AccessController) Execute(c *gin.Context) {

}
