package repository

import (
	"errors"
	"watt/pkg/models"
	"watt/pkg/utils"
	"watt/pkg/validation"

	"github.com/jinzhu/gorm"
)

type RoleRepository struct {
}

var RoleRep = newRoleRepository()

func newRoleRepository() *RoleRepository {
	return new(RoleRepository)
}

// 角色列表
func (r *RoleRepository) RoleList(page, size int) []models.Role {

	var role []models.Role

	err := models.Link.Where("status = 1").Order("created_at DESC").Offset((page - 1) * size).Limit(size).Find(&role).Error

	if err != nil && err != gorm.ErrRecordNotFound {
		utils.Error(err)
	}

	return role
}

// 添加角色
func (r *RoleRepository) InsertRole(param *validation.RoleData) error {

	var role models.Role

	err := models.Link.Select("id").Where("rolename = ? AND status = 1", param.Rolename).First(&role).Error

	if err != nil && err != gorm.ErrRecordNotFound {
		return errors.New("添加失败")
	}

	if role.ID > 0 {
		return errors.New("角色名称重复")
	}

	role.Rolename = param.Rolename
	role.Roledesc = param.Roledesc

	if err = models.Link.Create(&role).Error; err != nil {
		utils.Error(err)
		return errors.New("添加失败!")
	}

	return nil
}

// 删除角色
func (r *RoleRepository) Delete(id int) error {

	var role models.Role

	err := models.Link.Where("status = 1 AND id = ?", id).First(&role).Error

	if err != nil {
		return errors.New("角色有误")
	}

	role.Status = -1

	err = models.Link.Save(&role).Error

	if err != nil {
		return errors.New("更新失败")
	}

	return nil
}

// 更新角色
func (r *RoleRepository) Update(id int, param *validation.RoleData) error {

	var role models.Role

	err := models.Link.Select("id").Where("id = ?", id).First(&role).Error

	if err != nil && err != gorm.ErrRecordNotFound {
		return errors.New("角色查询错误")
	}

	if role.ID == 0 {
		return errors.New("角色有误!")
	}

	if err = models.Link.Model(&role).Updates(map[string]interface{}{"roledesc": param.Roledesc, "rolename": param.Rolename}).Error; err != nil {
		utils.Error(err)
		return errors.New("更新失败!")
	}

	return nil
}
