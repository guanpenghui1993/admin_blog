package models

type Model struct {
	ID uint `gorm:"primary_key;AUTO_INCREMENT;column:id;";json:"id"`
}

// 用户表
type Admin struct {
	UID         uint   `gorm:"primary_key;AUTO_INCREMENT;column:uid;";json:"id"`
	Username    string `gorm:"type:varchar(65);not null;unique;";json:"username"`
	Password    string `gorm:"type:char(32);not null";json:"password"`
	UserPic     string `gorm:"type:varchar(155);not null";json:"user_pic"`
	Nickname    string `gorm:"type:varchar(45);not null";json:"nickname"`
	AdminStatus int8   `gorm:"type:tinyint(1);not null;index:idx_user_status;default:1";json:"status"`
}
