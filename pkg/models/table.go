package models

import "watt/pkg/utils"

// 自动迁移
func AutoMigrateTable() {

	Link.Set("gorm:table_options", "ENGINE=InnoDB").AutoMigrate(&User{}, &Role{}, &Menu{}, &Access{})

	utils.Strace("----------------------------初始化表结构结束---------------------------")
}

// 用户表
type User struct {
	Base
	Roleid   uint   `json:"roleid" gorm:"type:int(10);not null;default:0"`
	Username string `json:"username" gorm:"type:varchar(65);not null;unique;"`
	Password string `json:"-" gorm:"type:char(32);not null"`
	UserPic  string `json:"user_pic" gorm:"type:varchar(155);not null"`
	Nickname string `json:"nickname" gorm:"type:varchar(45);not null"`
	Status   int    `json:"status" gorm:"type:tinyint(1);not null;index:idx_user_status;default:1;comment:'1 正常 0 禁用 -1 删除'"`
}

// 角色表
type Role struct {
	Base
	Rolename string `json:"role_name" gorm:"type:varchar(45);not null;unique;comment:'角色名称'"`
	Roledesc string `json:"role_desc" gorm:"type:varchar(65);not null;comment:'角色描述'"`
	Status   int    `json:"role_status" gorm:"type:tinyint(1);not null;index:idx_status;default:1;comment:'1 正常 -1 删除'"`
}

// 菜单表
type Menu struct {
	Base
	Pid      uint   `json:"pid" gorm:"type:int(10);not null;default:0;comment:'父级id 默认0为顶级分类'"`
	Menuname string `json:"menu_name" gorm:"type:varchar(65);not null;unique;comment:'菜单名称'"`
	Router   string `json:"router" gorm:"type:varchar(125);not null;unique;comment:'菜单路由地址使用 @ 符号分割控制器跟action'"`
	Icon     string `json:"icon" gorm:"type:varchar(25);not null;comment:'菜单ICON图标'"`
	Sort     int    `json:"sort" gorm:"type:tinyint(3);not null;index:sort;default:0;comment:'排序字段'"`
	Status   int    `json:"status" gorm:"type:tinyint(1);not null;index:idx_status;default:1;comment:'1 正常 -1 删除'"`
}

// 权限表
type Access struct {
	Roleid uint `json:"role_id" gorm:"type:int(10);not null;comment:'角色ID'"`
	Menuid uint `json:"menu_id" gorm:"type:int(10);not null;comment:'菜单ID'"`
}
