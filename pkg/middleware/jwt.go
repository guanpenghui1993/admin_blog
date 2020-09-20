package middleware

import "github.com/gin-gonic/gin"

func CheckLogin() gin.HandlerFunc {
	return func(c *gin.Context) {

		c.Next()
	}
}
