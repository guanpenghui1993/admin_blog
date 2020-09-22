package models

type User struct {
	Base
	Roleid   uint   `json:"roleid" gorm:"type:int(10);not null;default:0"`
	Username string `json:"username" gorm:"type:varchar(65);not null;unique;"`
	Password string `json:"-" gorm:"type:char(32);not null"`
	UserPic  string `json:"user_pic" gorm:"type:varchar(155);not null"`
	Nickname string `json:"nickname" gorm:"type:varchar(45);not null"`
	Status   int8   `json:"status" gorm:"type:tinyint(1);not null;index:idx_user_status;default:1;comment:'1 正常 0 禁用 -1 删除'"`
}

// func (User) TableName() string {
// 	return ""
// }
