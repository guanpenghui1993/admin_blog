package models

type Role struct {
	ID      int
	Rolename string `gorm:"type:varchar(45);not null;unique;"`
	Roledesc string `gorm:"type:varchar(65);not null;"`
	Status   int8   `gorm:"type:tinyint(1);not null;index:idx_user_status;default:1"`
}
