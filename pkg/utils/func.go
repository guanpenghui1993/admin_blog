package utils

import (
	"crypto/md5"
	"encoding/hex"
	"time"
)

// 返回当前时间格式化
func Datetime() string {
	return time.Now().Format("2006-01-02 15:04:05")
}

// md5 编码
func Md5(param string) string {

	h := md5.New()

	h.Write([]byte(param))

	cipherStr := h.Sum(nil)

	return hex.EncodeToString(cipherStr)
}
