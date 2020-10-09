package admin

import (
	"strconv"
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

	data := services.MenuService.MenuList()

	m.json(c, utils.SUCCESS, "获取成功", data)
}

// 添加菜单
func (m *MenuController) Create(c *gin.Context) {

	var insertData validation.InsertMenuData

	if err := m.valid(c, &insertData); err != nil {

		m.json(c, utils.ERROR, err.Error(), nil)

		return
	}

	if err := services.MenuService.Insert(&insertData); err != nil {
		m.json(c, utils.ERROR, err.Error(), nil)
		return
	}

	m.json(c, utils.SUCCESS, "添加成功", nil)
}

// 删除菜单
func (m *MenuController) Delete(c *gin.Context) {

	var param validation.BaseID

	if err := m.valid(c, &param); err != nil {

		m.json(c, utils.ERROR, err.Error(), nil)

		return
	}

	if err := services.MenuService.MenuDel(&param); err != nil {
		m.json(c, utils.ERROR, err.Error(), nil)
		return
	}

	m.json(c, utils.SUCCESS, "删除成功", nil)
}

// 编辑菜单
func (m *MenuController) Edite(c *gin.Context) {

	var updateData validation.InsertMenuData

	if err := m.valid(c, &updateData); err != nil {

		m.json(c, utils.ERROR, err.Error(), nil)

		return
	}

	id := strconv.Atoi(c.PostForm("id"))

	if id <= 0 {
		m.json(c, utils.ERROR, "id 有误", nil)
		return
	}

	// 编辑菜单信息
	if err := services.MenuService.MenuEdite(id, updateData); err != nil {
		m.json(c, utils.ERROR, err.Error(), nil)
		return
	}

	m.json(c, utils.SUCCESS, "更新成功", nil)
}
