package admin

import (
	"strconv"
	"watt/pkg/services"
	"watt/pkg/utils"
	"watt/pkg/validation"

	"github.com/gin-gonic/gin"
)

type RoleController struct {
	BaseController
}

// 获取角色列表
func (r *RoleController) List(c *gin.Context) {

	var param validation.BaseValid

	if err := r.valid(c, &param); err != nil {

		r.json(c, utils.ERROR, err.Error(), nil)

		return
	}

	data := services.RoleService.RoleList(&param)

	r.json(c, utils.SUCCESS, "获取成功", data)
}

// 添加角色
func (r *RoleController) Create(c *gin.Context) {

	var insertData validation.RoleData

	if err := r.valid(c, &insertData); err != nil {

		r.json(c, utils.ERROR, err.Error(), nil)

		return
	}

	if err := services.RoleService.Insert(&insertData); err != nil {
		r.json(c, utils.ERROR, err.Error(), nil)
		return
	}

	r.json(c, utils.SUCCESS, "添加成功", nil)
}

// 删除角色
func (r *RoleController) Delete(c *gin.Context) {

	var param validation.BaseID

	if err := r.valid(c, &param); err != nil {

		r.json(c, utils.ERROR, err.Error(), nil)

		return
	}

	if err := services.RoleService.Delete(&param); err != nil {
		r.json(c, utils.ERROR, err.Error(), nil)
		return
	}

	r.json(c, utils.SUCCESS, "删除成功", nil)
}

// 编辑角色
func (r *RoleController) Edite(c *gin.Context) {

	var data validation.RoleData

	if err := r.valid(c, &data); err != nil {

		r.json(c, utils.ERROR, err.Error(), nil)

		return
	}

	id, _ := strconv.Atoi(c.PostForm("id"))

	if id <= 0 {
		r.json(c, utils.ERROR, "id 有误", nil)
		return
	}

	if err := services.RoleService.Update(id, &data); err != nil {
		r.json(c, utils.ERROR, err.Error(), nil)
		return
	}

	r.json(c, utils.SUCCESS, "更新成功", nil)
}
