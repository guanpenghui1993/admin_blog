package services

import "watt/pkg/validation"

type menuService struct {
}

var MenuService = newMenuService()

func newMenuService() *menuService {
	return new(menuService)
}

// 添加菜单
func (m *menuService) Insert(param *validation.InsertMenuData) error {

}
