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
func (m *menuService) MenuList() interface{} {

	data := repository.MenuRep.List()

	if len(data) > 0 { // 递归数组
		return m.menuTree(&data, 0)
	}

	return data
}

// 删除菜单
func (m *menuService) MenuDel(param *validation.BaseID) error {

	return repository.MenuRep.Delete(param.ID)
}

// 递归处理菜单
func (m *menuService) menuTree(array *[]models.Menu, pid uint) []validation.MenuList {

	var result []validation.MenuList

	for _, val := range *array {

		if val.Pid == pid {

			tmp := validation.MenuList{
				val.Menuname,
				val.Icon,
				val.Router,
				val.Pid,
				val.ID,
				m.menuTree(array, val.ID),
			}

			result = append(result, tmp)
		}
	}

	return result
}

// 编辑菜单
func (m *menuService) MenuEdite(id int, param *validation.InsertMenuData) error {

	if param.Pid > 0 {

		menu := repository.MenuRep.Info(param.Pid)

		if menu.ID <= 0 {
			return errors.New("菜单有误")
		}
	}

	param.Menuname = html.EscapeString(strings.TrimSpace(param.Menuname))

	param.Router = html.EscapeString(strings.TrimSpace(param.Router))

	param.Icon = html.EscapeString(strings.TrimSpace(param.Icon))

	return repository.MenuRep.UpdateMenu(id, param)
}
