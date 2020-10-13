package utils

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
)

// 返回当前时间格式化
func Datetime() string {
	return time.Now().Format("2006-01-02 15:04:05")
}

// 返回当前Unix时间
func Time() int64 {
	return time.Now().Unix()
}

// Unix 转 format
func Dateformat(timeint int64) string {
	return time.Unix(timeint, 0).Format("2006-01-02 15:04:05")
}

// md5 编码
func Md5(param string) string {

	h := md5.New()

	h.Write([]byte(param))

	cipherStr := h.Sum(nil)

	return hex.EncodeToString(cipherStr)
}

// 创建目录
func Mkdir(dir string) {

	err := os.MkdirAll(dir, os.ModePerm)

	if err != nil {
		Exit(err)
	}
}

// 生成token
func Token(uid uint) (string, error) {

	type Claims struct {
		Userid uint
		jwt.StandardClaims
	}

	claims := Claims{
		uid,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(Setting.Jwt.Expired * time.Second).Unix(),
			Issuer:    Setting.Jwt.Issuer,
		},
	}

	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	token, err := tokenClaims.SignedString([]byte(Setting.Jwt.Secret))

	return token, err
}

// 解析token
func Parse(token string) (int, error) {

	claim, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {

		return []byte(Setting.Jwt.Secret), nil
	})

	if err != nil {

		return 0, err
	}

	mapinfo := claim.Claims.(jwt.MapClaims)

	userid := mapinfo["Userid"].(float64)

	return int(userid), nil
}

// 打印日志并终止程序
func Exit(message interface{}) {

	log.SetFlags(log.Llongfile | log.Ldate | log.Ltime)

	if Setting.Common.Debug {
		log.SetPrefix(ERROR_COLOR_PREFIX)
		log.Fatal(colorString(message, 0))
	}

	log.SetPrefix(ERROR_PREFIX)

	logFile, err := os.OpenFile(Setting.Common.Log+"/"+ERROR_LOG, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)

	if err != nil {
		log.Fatal(colorString(err, 0))
	}

	log.SetOutput(logFile)

	log.Fatal("Fail to read file -- ", message)
}

// 打印日志不终止程序
func Strace(message interface{}) {

	if Setting.Common.Debug {
		log.SetPrefix(STRACE_COLOR_PREFIX)
		log.Printf(colorString(message, 1))
		return
	}

	log.SetPrefix(STRACE_PREFIX)

	logFile, err := os.OpenFile(Setting.Common.Log+"/"+STRACE_LOG, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)

	if err != nil {
		Exit(err)
	}

	log.SetOutput(logFile)

	log.Printf("%v \n", message)
}

// 记录错误日志不终止程序
func Error(message interface{}) {

	if Setting.Common.Debug {
		log.SetPrefix(STRACE_COLOR_PREFIX)
		log.Printf(colorString(message, 0))
		return
	}

	log.SetPrefix(STRACE_PREFIX)

	logFile, err := os.OpenFile(Setting.Common.Log+"/"+ERROR_LOG, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)

	if err != nil {
		Exit(err)
	}

	log.SetOutput(logFile)

	log.Printf("%v \n", message)
}

// 返回颜色shell字符串
func colorString(msg interface{}, color uint) string {

	if color == 0 {
		return "\033[41;37m " + fmt.Sprintf("%v", msg) + " \033[0m"
	}

	return "\033[42;37m " + fmt.Sprintf("%v", msg) + " \033[0m"
}
