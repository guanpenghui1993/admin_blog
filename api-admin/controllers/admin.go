package controllers

import (
	"github.com/gin-gonic/gin"
)

func Printdemo(c *gin.Context) {

	c.JSON(200, gin.H{
		"code": 200,
		"data": nil,
	})
}
