package admin

import (
	"watt/pkg/services"
	"watt/pkg/utils"
	"watt/pkg/validation"

	"github.com/gin-gonic/gin"
)

type AccessController struct {
	BaseController
}

// 设置角色权限
func (a *AccessController) Execute(c *gin.Context) {

	var accessNode validation.AccessData

	if err := a.valid(c, &accessNode); err != nil {

		a.json(c, utils.ERROR, err.Error(), nil)

		return
	}

	if err := services.AccessService.SetRoleNode(&accessNode); err != nil {
		a.json(c, utils.ERROR, err.Error(), nil)
		return
	}

	a.json(c, utils.SUCCESS, "操作成功", nil)
}
