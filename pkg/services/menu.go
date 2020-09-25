package services

import (
	"errors"
	"html"
	"strings"
	"watt/pkg/models"
	"watt/pkg/repository"
	"watt/pkg/validation"
)

type menuService struct {
}

var MenuService = newMenuService()

func newMenuService() *menuService {
	return new(menuService)
}

// 添加菜单
func (m *menuService) Insert(param *validation.InsertMenuData) error {

	if param.Pid > 0 {

		menu := repository.MenuRep.Info(param.Pid)

		if menu.ID <= 0 {
			return errors.New("菜单有误")
		}
	}

	param.Menuname = html.EscapeString(strings.TrimSpace(param.Menuname))

	param.Router = html.EscapeString(strings.TrimSpace(param.Router))

	param.Icon = html.EscapeString(strings.TrimSpace(param.Icon))

	return repository.MenuRep.InsertMenu(param)
}

// 菜单列表
func (m *menuService) MenuList(param *validation.BaseValid) []models.Menu {

	data := repository.MenuRep.List(param.Page, param.Size)

	if len(data) > 0 { // 递归数组
		return m.menuTree(&data)
	}

	return data
}

// 删除菜单
func (m *menuService) MenuDel(param *validation.BaseID) bool {

	return repository.MenuRep.Delete(param.ID)
}

// 递归处理菜单
func (m *menuService) menuTree(array *[]models.Menu) []models.Menu {

	data := *array

	return data
}
