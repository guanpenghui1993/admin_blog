package router

import (
	"watt/pkg/config"
	"watt/pkg/controllers/admin"

	"github.com/gin-gonic/gin"
)

var Route *gin.Engine

func init() {

	var mode string

	if config.Conf.Common.Debug {
		mode = "debug"
	} else {
		mode = "release"
	}

	// 初始化router引擎
	Route = gin.Default()

	// 运行模式
	gin.SetMode(mode)

	// 绑定后台路由
	adminRoute := Route.Group("/admin")

	// 后台登录
	adminRoute.GET("/login", admin.Login)
}