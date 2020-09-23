package models

// 用户表
type User struct {
	Base
	Roleid   uint   `json:"roleid" gorm:"type:int(10);not null;default:0"`
	Username string `json:"username" gorm:"type:varchar(65);not null;unique;"`
	Password string `json:"-" gorm:"type:char(32);not null"`
	UserPic  string `json:"user_pic" gorm:"type:varchar(155);not null"`
	Nickname string `json:"nickname" gorm:"type:varchar(45);not null"`
	Status   int8   `json:"status" gorm:"type:tinyint(1);not null;index:idx_user_status;default:1;comment:'1 正常 0 禁用 -1 删除'"`
}

// 角色表
type Role struct {
	Base
	Rolename string `gorm:"type:varchar(45);not null;unique;"`
	Roledesc string `gorm:"type:varchar(65);not null;"`
	Status   int8   `gorm:"type:tinyint(1);not null;index:idx_status;default:1;comment:'1 正常 -1 删除'"`
}

// 菜单表
type Menu struct {
	Base
	Pid      uint   `gorm:"type:int(10);not null;default:0"`
	Menuname string `gorm:"type:varchar(65);not null;unique;"`
	Router   string `gorm:"type:varchar(125);not null;unique;"`
	Icon     string `gorm:"type:varchar(25);not null;"`
	Status   int8   `gorm:"type:tinyint(1);not null;index:idx_status;default:1;comment:'1 正常 -1 删除'"`
}

// 权限表
type Access struct {
	Roleid uint `gorm:"type:int(10);not null"`
	Menuid uint `gorm:"type:int(10);not null"`
}
