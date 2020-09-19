package models

type Role struct {
	Base
	Rolename string `gorm:"type:varchar(45);not null;unique;"`
	Roledesc string `gorm:"type:varchar(65);not null;"`
	Status   int8   `gorm:"type:tinyint(1);not null;index:idx_status;default:1;comment:'1 正常 -1 删除'"`
}
