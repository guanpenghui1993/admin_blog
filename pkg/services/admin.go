package services

import (
	"crypto/md5"
	"encoding/hex"
	"watt/pkg/models"
)

type AdminService struct {
}

var AdminServiceObject = newAdminService()

func newAdminService() *AdminService {
	return new(AdminService)
}

// 登录
func (a *AdminService) Login(username, password string) (bool, error) {

	var admin models.Admin

	h := md5.New()

	h.Write([]byte(password))

	cipherStr := h.Sum(nil)

	err := models.Link.Select("uid").Where("username = ? AND password = ?", username, hex.EncodeToString(cipherStr)).First(&admin).Error

	if err != nil {
		return false, err
	}

	return true, nil
}
