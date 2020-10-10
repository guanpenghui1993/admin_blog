package repository

import (
	"watt/pkg/models"
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
