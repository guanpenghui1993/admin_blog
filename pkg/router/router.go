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

	// 全局异常 日志 跨域
	Route.Use(middleware.Recovery(), gin.Logger(), middleware.Cors())

	// 运行模式
	gin.SetMode(mode)

	// 对象struct
	menu := new(admin.MenuController)     // 菜单类
	user := new(admin.UserController)     // 用户类
	role := new(admin.RoleController)     // 角色类
	access := new(admin.AccessController) // 权限类

	// 绑定后台路由
	adminRoute := Route.Group("/admin")

	// 后台登录
	adminRoute.POST("/login", user.Login)

	// 检测登录Token跟路由权限
	adminRoute.Use(middleware.CheckLogin(), middleware.Access())
	{
		// 后台用户接口
		adminRoute.GET("/user/info", user.Info)          // 用户信息
		adminRoute.GET("/user/list", user.UserList)      // 用户列表
		adminRoute.POST("/user/delete", user.UserDel)    // 删除用户
		adminRoute.POST("/user/create", user.UserAdd)    // 创建用户
		adminRoute.POST("/user/status", user.UserUpdate) // 更新状态

		// 后台菜单接口
		adminRoute.GET("/menu/list", menu.List)      // 菜单列表
		adminRoute.POST("/menu/create", menu.Create) // 创建菜单
		adminRoute.POST("/menu/delete", menu.Delete) // 删除菜单
		adminRoute.POST("menu/edite", menu.Edite)    // 编辑菜单

		// 后台角色接口
		adminRoute.GET("/role/list", role.List)      // 角色列表
		adminRoute.POST("/role/create", role.Create) // 创建角色
		adminRoute.POST("/role/delete", role.Delete) // 删除角色
		adminRoute.POST("role/edite", role.Edite)    // 编辑角色

		// 设置角色权限
		adminRoute.POST("access/set", access.Execute) // 设置权限
	}
}
