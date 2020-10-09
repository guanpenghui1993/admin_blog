package services

import (
	"watt/pkg/models"
	"watt/pkg/repository"
	"watt/pkg/validation"
)

var RoleService = newRoleService()

type roleService struct {
}

func newRoleService() *roleService {
	return new(roleService)
}

// 获取角色列表
func (r *roleService) RoleList(param *validation.BaseValid) []models.Role {

	return repository.RoleRep.RoleList(param.Page, param.Size)
}

// 添加角色
func (r *roleService) Insert(data *validation.RoleData) error {

	return repository.RoleRep.InsertRole(data)
}

// 删除角色
func (r *roleService) Delete(data *validation.BaseID) error {

	return repository.RoleRep.Delete(data.ID)
}

// 更新角色
func (r *roleService) Update(id int, data *validation.RoleData) error {

	return repository.RoleRep.Update(id, data)
}
