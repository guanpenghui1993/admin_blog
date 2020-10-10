package services

import "watt/pkg/repository"

type accessService struct {
}

var AccessService = newAccessService()

func newAccessService() *accessService {
	return new(accessService)
}

// 获取路由列表
func (a *accessService) RouterData(role uint) []string {

	var routerList []string

	data := repository.AccessRep.GetAccessAll(role)

	if len(data) > 0 {

		for _, val := range data {

			routerList = append(routerList, val.Router)
		}
	}

	return routerList
}

// 设置角色权限节点
