package router

import (
	"api-admin/controllers/admin"
	"api-admin/handler"
	"api-admin/helpers"

	"github.com/gin-gonic/gin"
)

var Route *gin.Engine

func init() {

	// 初始化router引擎
	Route = gin.New()

	// 运行模式
	gin.SetMode(helpers.ServerConfig.Mode)

	// 异常处理,授权登录
	Route.Use(handler.LogFormat(), gin.Logger(), handler.RecoverHand())

	// 绑定后台路由
	adminRoute := Route.Group("/admin")

	// 绑定前台路由
	apiRoute := Route.Group("/api")

	// 后台登录
	adminRoute.GET("/login", admin.Login)

	// 后台验证token
	adminRoute.Use(handler.CheckLoginHand())
	{

	}

	// 前台路由
	apiRoute.Use(handler.Cors())
	{
		apiRoute.GET("/demo", func(c *gin.Context) {
			c.JSON(200, gin.H{"messge": "卧槽昂围殴那狗翁哦澳网"})
		})
	}
}
