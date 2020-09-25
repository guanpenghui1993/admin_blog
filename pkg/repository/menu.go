package repository

import (
	"errors"
	"watt/pkg/models"
	"watt/pkg/utils"
	"watt/pkg/validation"

	"github.com/jinzhu/gorm"
)

type MenuRepository struct {
}

var MenuRep = newMenuRepository()

func newMenuRepository() *MenuRepository {
	return new(MenuRepository)
}

// 根据ID 获取菜单信息
func (m *MenuRepository) Info(id uint) models.Menu {

	var menu models.Menu

	err := models.Link.Where("id = ? AND status = 1", id).First(&menu).Error

	if err != nil && err != gorm.ErrRecordNotFound {
		utils.Error(err)
	}

	return menu
}

// 菜单列表
func (m *MenuRepository) List(page, size int) []models.Menu {

	var menu []models.Menu

	err := models.Link.Where("status = 1").Order("sort DESC").Offset((page - 1) * size).Limit(size).Find(&menu).Error

	if err != nil && err != gorm.ErrRecordNotFound {
		utils.Error(err)
	}

	return menu
}

// 添加菜单
func (m *MenuRepository) InsertMenu(param *validation.InsertMenuData) error {

	var menu models.Menu

	err := models.Link.Select("id").Where("menuname = ? AND status = 1", param.Menuname).First(&menu).Error

	if err != nil && err != gorm.ErrRecordNotFound {
		return errors.New("添加失败")
	}

	if menu.ID > 0 {
		return errors.New("菜单名重复")
	}

	menu.Pid = param.Pid
	menu.Menuname = param.Menuname
	menu.Icon = param.Icon
	menu.Router = param.Router
	menu.Sort = param.Sort

	if err = models.Link.Create(&menu).Error; err != nil {
		utils.Error(err)
		return errors.New("添加失败!")
	}

	return nil
}

// 删除菜单
func (m *MenuRepository) Delete(id int) bool {

	var menu models.Menu

	err := models.Link.Where("status = 1 AND id = ?", id).First(&menu).Error

	if err != nil {
		return false
	}

	menu.Status = -1

	err = models.Link.Save(&menu).Error

	if err != nil {
		return false
	}

	return true
}
