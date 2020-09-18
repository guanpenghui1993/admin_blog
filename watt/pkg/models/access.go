package models

type Access struct {
	ID     int
	Roleid `gorm:"many2many:watt_role;"`
	Menuid `gorm:"many2many:watt_menu;"`
}
