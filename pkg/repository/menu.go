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

// 根据id批量获取菜单路由
func (m *MenuRepository) GetMenuAllById(ids *[]int) []models.Menu {

	var menu []models.Menu

	err := models.Link.Where("status = 1 AND id IN (?)", *ids).Find(&menu).Error

	if err != nil && err != gorm.ErrRecordNotFound {
		utils.Error(err)
	}

	return menu
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
func (m *MenuRepository) List() []models.Menu {

	var menu []models.Menu

	err := models.Link.Where("status = 1").Order("sort DESC").Find(&menu).Error

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

// 更新菜单
func (m *MenuRepository) UpdateMenu(id int, param *validation.InsertMenuData) error {

	var menu models.Menu

	err := models.Link.Where("id = ?", id).First(&menu).Error

	if err != nil {
		return errors.New("菜单有误")
	}

	menu.Pid = param.Pid
	menu.Menuname = param.Menuname
	menu.Icon = param.Icon
	menu.Router = param.Router
	menu.Sort = param.Sort

	err = models.Link.Model(&menu).Updates(map[string]interface{}{"pid": param.Pid, "menuname": param.Menuname, "icon": param.Icon, "router": param.Router, "sort": param.Sort}).Error

	if err != nil {
		return errors.New("更新失败")
	}

	return nil
}

// 删除菜单
func (m *MenuRepository) Delete(id int) error {

	var menu, pmenu models.Menu

	models.Link.Where("status = 1 AND pid = ?", id).First(&pmenu)

	if pmenu.ID > 0 {
		return errors.New("存在子级菜单")
	}

	err := models.Link.Where("status = 1 AND id = ?", id).First(&menu).Error

	if err != nil {
		return errors.New("菜单有误")
	}

	menu.Status = -1

	err = models.Link.Save(&menu).Error

	if err != nil {
		return errors.New("更新失败")
	}

	return nil
}
