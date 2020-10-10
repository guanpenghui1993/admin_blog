package repository

import (
	"errors"
	"watt/pkg/models"
	"watt/pkg/utils"

	"github.com/jinzhu/gorm"
)

type AccessRepository struct {
}

var AccessRep = newAccessRepository()

func newAccessRepository() *AccessRepository {
	return new(AccessRepository)
}

// 获取所有权限节点
func (a *AccessRepository) GetAccessAll(roleid uint) []models.Menu {

	var (
		access   models.Access
		menuList []models.Menu
		menuid   []int
	)

	models.Link.Model(&access).Where("roleid = ?", roleid).Pluck("menuid", &menuid)

	if len(menuid) <= 0 {
		return menuList
	}

	return MenuRep.GetMenuAllById(&menuid)
}

// 删除角色下的节点
func (a *AccessRepository) DeleteRoleNode(roleid uint) error {

	var access models.Access

	err := models.Link.Where("roleid = ?", roleid).Delete(&access).Error

	if err != nil {
		utils.Error(err)
		return errors.New("删除失败")
	}

	return nil
}

// 添加角色权限节点
func (a *AccessRepository) InsertRoleNode(data *[]models.Access) error {

	err := models.Link.Create(data).Error

	if err != nil {
		utils.Error(err)
		return errors.New("添加失败")
	}

	return nil
}

// 设置角色权限(事务方式)
func (a *AccessRepository) SetRoleNode(roleid uint, data *[]models.Access) error {

	err := models.Link.Transaction(func(tx *gorm.DB) error {

		var access models.Access

		if err := tx.Where("roleid = ?", roleid).Delete(&access).Error; err != nil {
			utils.Error(err)
			utils.Error("----------------------------------")
			return err
		}

		if erradd := tx.Create(data).Error; erradd != nil {
			utils.Error(erradd)
			utils.Error("+++++++++++++++++++++++++++++++++++++")
			return erradd
		}

		// 返回 nil 提交事务
		return nil
	})

	return err
}
