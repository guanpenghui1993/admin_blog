package models

type User struct {
	ID       int
	Roleid   uint   `gorm:"type:int(10);not null;default:0"`
	Username string `gorm:"type:varchar(65);not null;unique;"`
	Password string `gorm:"type:char(32);not null"`
	UserPic  string `gorm:"type:varchar(155);not null"`
	Nickname string `gorm:"type:varchar(45);not null"`
	Status   int8   `gorm:"type:tinyint(1);not null;index:idx_user_status;default:1"`
}

// func (User) TableName() string {
// 	return ""
// }
