package services

import (
	"strconv"
	"strings"
	"watt/pkg/models"
	"watt/pkg/repository"
	"watt/pkg/validation"
)

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
func (a *accessService) SetRoleNode(node *validation.AccessData) error {

	var nodeList []models.Access

	menuData := strings.Split(node.Idstring, ",")

	for _, val := range menuData {

		vals, _ := strconv.Atoi(val)

		nodeList = append(nodeList, models.Access{node.Roleid, uint(vals)})
	}

	return repository.AccessRep.SetRoleNode(node.Roleid, &nodeList)
}
