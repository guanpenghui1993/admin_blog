package router

import (
	"api-admin/controllers"
	"api-admin/helpers"

	"github.com/gin-gonic/gin"
)

var Route *gin.Engine

func init() {

	// 初始化router引擎
	Route = gin.Default()

	// 运行模式
	gin.SetMode(helpers.ServerConfig.Mode)

	// 全局中间件登录cookie

	// 绑定路由
	Admin := Route.Group("/admin")
	{
		Admin.GET("/:id", controllers.Printdemo)
	}
}
