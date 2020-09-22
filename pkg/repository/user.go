package repository

import (
	"watt/pkg/models"
	"watt/pkg/utils"

	"github.com/jinzhu/gorm"
)

type UserRepository struct {
}

var UserRep = newUserRepository()

func newUserRepository() *UserRepository {
	return new(UserRepository)
}

// 验证用户名密码是否匹配
func (u *UserRepository) UserInfoByName(username, password string) models.User {

	var user models.User

	err := models.Link.Select("id").Where("username = ? AND password = ? AND status = 1", username, utils.Md5(password)).First(&user).Error

	if err != nil && err != gorm.ErrRecordNotFound {
		utils.Error(err)
	}

	return user
}

// 通过用户ID获取用户信息
func (u *UserRepository) UserInfoById(userid int) models.User {

	var user models.User

	err := models.Link.Select("id,nickname,user_pic,roleid").Where("id = ? AND status = 1", userid).First(&user).Error

	if err != nil && err != gorm.ErrRecordNotFound {
		utils.Error(err)
	}

	return user
}
