package admin

import (
	"watt/pkg/services"
	"watt/pkg/utils"
	"watt/pkg/validation"

	"github.com/gin-gonic/gin"
)

type MenuController struct {
	BaseController
}

// 获取菜单列表
func (m *MenuController) List(c *gin.Context) {

	// var param validation.BaseValid

}

// 添加菜单
func (m *MenuController) Create(c *gin.Context) {

	var insertData validation.InsertMenuData

	if err := u.valid(c, &insertData); err != nil {

		u.json(c, utils.ERROR, err.Error(), nil)

		return
	}

	if err := services.MenuService.Insert(&insertData); err != nil {
		u.json(c, utils.ERROR, err.Error(), nil)
		return
	}

	u.json(c, utils.SUCCESS, "添加成功", nil)
}

// 删除菜单
func (m *MenuController) Delete(c *gin.Context) {

}

// 编辑菜单
func (m *MenuController) Edite(c *gin.Context) {

}
