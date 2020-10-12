package router

import (
	"watt/pkg/controllers/admin"
	"watt/pkg/middleware"
	"watt/pkg/utils"

	"github.com/gin-gonic/gin"
)

// 全局变量
var (
	Route  *gin.Engine
	menu   = new(admin.MenuController)   // 菜单类
	user   = new(admin.UserController)   // 用户类
	role   = new(admin.RoleController)   // 角色类
	access = new(admin.AccessController) // 权限类
)

func init() {

	// 初始化router引擎
	Route = gin.New()

	// 运行模式
	if utils.Setting.Common.Debug {
		gin.SetMode(gin.DebugMode)
		Route.Use(gin.Logger())
	} else {
		gin.SetMode(gin.ReleaseMode)
	}

	// 全局异常 日志 跨域
	Route.Use(middleware.Recovery(), middleware.Cors())

	// 后台路由
	adminRouter()

	// 前台路由
	apiRouter()
}
