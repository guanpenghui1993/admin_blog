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
	Rolename string `json:"role_name" gorm:"type:varchar(45);not null;unique;"`
	Roledesc string `json:"role_desc" gorm:"type:varchar(65);not null;"`
	Status   int8   `json:"role_status" gorm:"type:tinyint(1);not null;index:idx_status;default:1;comment:'1 正常 -1 删除'"`
}

// 菜单表
type Menu struct {
	Base
	Pid      uint   `json:"pid" gorm:"type:int(10);not null;default:0"`
	Menuname string `json:"menu_name" gorm:"type:varchar(65);not null;unique;"`
	Router   string `json:"router" gorm:"type:varchar(125);not null;unique;"`
	Icon     string `json:"icon" gorm:"type:varchar(25);not null;"`
	Status   int8   `json:"status" gorm:"type:tinyint(1);not null;index:idx_status;default:1;comment:'1 正常 -1 删除'"`
}

// 权限表
type Access struct {
	Roleid uint `json:"role_id" gorm:"type:int(10);not null"`
	Menuid uint `json:"menu_id" gorm:"type:int(10);not null"`
}
