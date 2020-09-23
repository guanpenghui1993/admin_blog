package models

import (
	"fmt"
	"time"
)

type Base struct {
	ID        uint   `json:"id" gorm:"primary_key;AUTO_INCREMENT"`
	CreatedAt MyTime `json:"created_time" gorm:"type:datetime;not null;default: CURRENT_TIMESTAMP(0)"`
	UpdatedAt MyTime `json:"update_time" gorm:"type:datetime;not null;default: CURRENT_TIMESTAMP(0) ON UPDATE CURRENT_TIMESTAMP(0)"`
}

type MyTime time.Time

func (t MyTime) MarshalJSON() ([]byte, error) {

	formatted := fmt.Sprintf("\"%v\"", time.Time(t).Format("2006-01-02 15:04:05"))

	return []byte(formatted), nil
}
