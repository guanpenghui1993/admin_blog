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

	// if err != nil {
	// 	c.String(201, err)
	// }

	c.JSON(http.StatusOK, gin.H{
		"code":    1,
		"message": err,
		"data":    list,
	})

	// c.String(200, "hello admin controllers param is "+id)
}
