package router

import (
	"watt/pkg/controllers/admin"
	"watt/pkg/middleware"
	"watt/pkg/utils"

	"github.com/gin-gonic/gin"
)

var Route *gin.Engine

func init() {

	var mode string

	if utils.Setting.Common.Debug {
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

	user := new(admin.UserController)

	adminRoute.GET("/login", user.Login) // 后台登录

	// 验证token中间件
	adminRoute.Use(middleware.CheckLogin())
	{
		adminRoute.GET("/info", user.Info)
	}

}
