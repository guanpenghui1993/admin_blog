package models

type Access struct {
	Roleid uint `gorm:"type:int(10);not null"`
	Menuid uint `gorm:"type:int(10);not null"`
}
