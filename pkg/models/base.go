package models

import "time"

type Base struct {
	ID        uint      `json:"id" gorm:"primary_key;AUTO_INCREMENT"`
	CreatedAt time.Time `json:"created_time" gorm:"type:datetime;not null;default: CURRENT_TIMESTAMP(0)"`
	UpdatedAt time.Time `json:"update_time" gorm:"type:datetime;not null;default: CURRENT_TIMESTAMP(0) ON UPDATE CURRENT_TIMESTAMP(0)"`
}
