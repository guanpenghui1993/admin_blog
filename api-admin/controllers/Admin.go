package controllers

import (
	"api-admin/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Printdemo(c *gin.Context) {

	// id := c.Param("id")

	var admin = &models.Admin{}

	list, err := admin.Select()

	if err != nil {
		c.JSON(204, err)
	}

	c.JSON(http.StatusOK, returnJson(200, "获取成功", list))
}

func returnJson(code int, message string, data interface{}) gin.H {

	var result gin.H

	result = make(map[string]interface{})

	result["code"] = code
	result["message"] = message
	result["data"] = data

	return result
}
