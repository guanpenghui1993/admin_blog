package controllers

import (
	"api-admin/models"

	"github.com/gin-gonic/gin"
)

func Printdemo(c *gin.Context) {

	list := models.GetUserList()

	c.JSON(200, gin.H{
		"code": 200,
		"data": list,
	})
}
