package services

import (
	"api-admin/models"
)

type AdminService struct {
}

var AdminServiceObject = newAdminService()

func newAdminService() *AdminService {
	return new(AdminService)
}

// 测试登录
func (a *AdminService) Login(username, password string) bool {

	var admin models.Admin

	err := models.Link.Where("username = ? AND password = ?", username, password).First(&admin).Error

	if err != nil {
		return false
	}

	return true
}
