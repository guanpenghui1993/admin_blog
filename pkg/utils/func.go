package utils

import (
	"time"
)

// 返回当前时间格式化
func Datetime() string {
	return time.Now().Format("2006-01-02 15:04:05")
}
