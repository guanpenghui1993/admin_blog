package models

type Menu struct {
	ID       int
	Pid      uint   `gorm:"type:int(10);not null;default:0"`
	Menuname string `gorm:"type:varchar(65);not null;unique;"`
	Router   string `gorm:"type:varchar(125);not null;unique;"`
	Icon     string `gorm:"type:varchar(25);not null;"`
	Status   int8   `gorm:"type:tinyint(1);not null;index:idx_user_status;default:1"`
}
