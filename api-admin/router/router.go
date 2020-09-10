package router

import (
	"api-admin/controllers"
	"api-admin/helpers"
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
)

var Route *gin.Engine

func init() {

	// 初始化router引擎
	Route = gin.New()

	// 运行模式
	gin.SetMode(helpers.ServerConfig.Mode)

	// 异常处理,授权登录
	Route.Use(LogFormat(), gin.Logger(), recoverHand(), checkLoginHand())

	// 绑定路由
	admin := Route.Group("/admin")
	{
		admin.GET("/:id", controllers.Login)
	}
}

// 异常处理
func recoverHand() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if r := recover(); r != nil {
				helpers.ErrLog(r)
				c.AbortWithStatusJSON(200, helpers.ResponseMap(helpers.SERVER_ERROR_CODE, 0, fmt.Sprintf("%v", r), nil))
				return
			}
		}()

		c.Next()
	}
}

// 登录检测
func checkLoginHand() gin.HandlerFunc {

	return func(c *gin.Context) {

		token := c.GetHeader(helpers.CommonConfig.HeaderToken)

		if token == "" {
			c.AbortWithStatusJSON(200, helpers.ResponseMap(helpers.ERROR_CODE, 0, "缺少token信息", nil))
			return
		}

		_, err := helpers.Parse(token)

		if err != nil {
			c.AbortWithStatusJSON(200, helpers.ResponseMap(helpers.ERROR_CODE, 0, fmt.Sprintf("%v", err), nil))
			return
		}

		c.Next()
	}
}

// 自定义日志格式
func LogFormat() gin.HandlerFunc {

	return gin.LoggerWithFormatter(func(param gin.LogFormatterParams) string {

		return fmt.Sprintf("%s - [%s] \"%s %s %s %d %s \"%s\" %s\"\n",
			param.ClientIP,
			param.TimeStamp.Format(time.RFC1123),
			param.Method,
			param.Path,
			param.Request.Proto,
			param.StatusCode,
			param.Latency,
			param.Request.UserAgent(),
			param.ErrorMessage,
		)
	})
}
