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
	Route = gin.New()

	// 全局异常
	Route.Use(middleware.Recovery())

	// 运行模式
	gin.SetMode(mode)

	// 绑定后台路由
	adminRoute := Route.Group("/admin")

	// 用户类
	user := new(admin.UserController)

	// 后台登录
	adminRoute.GET("/login", user.Login)

	// 后台接口
	adminRoute.Use(middleware.CheckLogin())
	{
		adminRoute.GET("/userinfo", user.Info)
		adminRoute.GET("/userlist", user.UserList)
		adminRoute.GET("/userdel", user.UserDel)
		adminRoute.GET("/useradd", user.UserAdd)
	}
}
