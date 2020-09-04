package models

import (
	"github.com/jinzhu/gorm"
)

type Admin struct {
	UID         uint   `gorm:"primary_key"json:"id"`
	Username    string `json:"username"`
	Password    string `json:"password"`
	UserPic     string `json:"user_pic"`
	Nickname    string `json:"nickname"`
	AdminStatus int8   `json:"status"`
}

func (a *Admin) Select() ([]*Admin, error) {

	defer Close()

	var adminList []*Admin

	err := link.Where("admin_status = ?", 1).Find(&adminList).Error

	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}

	return adminList, nil
}
