package router

import (
	"watt/pkg/middleware"
)

func adminRouter() {

	// 绑定后台路由
	var adminRoute = Route.Group("/admin")

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
