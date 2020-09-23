package repository

import (
	"errors"
	"watt/pkg/models"
	"watt/pkg/utils"
	"watt/pkg/validation"

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

	err := models.Link.Where("id = ? AND status = 1", userid).First(&user).Error

	if err != nil && err != gorm.ErrRecordNotFound {
		utils.Error(err)
	}

	return user
}

// 获取用列表
func (u *UserRepository) UserList(page, size int) []models.User {

	var user []models.User

	err := models.Link.Where("status = 1").Order("created_at DESC").Offset((page - 1) * size).Limit(size).Find(&user).Error

	if err != nil && err != gorm.ErrRecordNotFound {
		utils.Error(err)
	}

	return user
}

// 删除用户信息
func (u *UserRepository) DeleteUser(uid int) bool {

	var user models.User

	err := models.Link.Where("status = 1 AND id = ?", uid).First(&user).Error

	if err != nil {
		return false
	}

	user.Status = -1

	err = models.Link.Save(&user).Error

	if err != nil {
		return false
	}

	return true
}

// 添加用户
func (u *UserRepository) InsertUser(param *validation.InsertUserData) error {

	var user models.User

	err := models.Link.Select("id").Where("username = ? AND status > 0", param.Username).First(&user).Error

	if err != nil && err != gorm.ErrRecordNotFound {
		return errors.New("添加失败")
	}

	if user.ID > 0 {
		return errors.New("用户名重复")
	}

	user.Roleid = param.Roleid
	user.Username = param.Username
	user.Password = param.Password
	user.UserPic = param.UserPic
	user.Nickname = param.Nickname

	if err = models.Link.Create(&user).Error; err != nil {
		utils.Error(err)
		return errors.New("添加失败!")
	}

	return nil
}
