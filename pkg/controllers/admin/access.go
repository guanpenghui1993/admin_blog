package admin

import (
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

	if err := m.valid(c, &accessNode); err != nil {

		m.json(c, utils.ERROR, err.Error(), nil)

		return
	}

}
