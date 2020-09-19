package models

import "time"

type Base struct {
	ID        uint      `gorm:"primary_key;AUTO_INCREMENT"`
	CreatedAt time.Time `gorm:"type:datetime;not null;default: CURRENT_TIMESTAMP(0)"`
	UpdatedAt time.Time `gorm:"type:datetime;not null;default: CURRENT_TIMESTAMP(0) ON UPDATE CURRENT_TIMESTAMP(0)"`
}
