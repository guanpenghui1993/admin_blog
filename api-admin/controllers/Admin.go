package controllers

import (
	"api-admin/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

var result = make(map[string]interface{})

func Printdemo(c *gin.Context) {

	// id := c.Param("id")

	var admin = new(models.Admin)

	list, err := admin.Select()

	if err != nil {
		c.JSON(204, err)
	}

	result["code"] = 200
	result["message"] = "xxxresult"
	result["data"] = list

	c.JSON(http.StatusOK, result)
}
