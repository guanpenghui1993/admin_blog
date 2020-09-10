package helpers

import (
	"errors"
	"fmt"
	"io"
	"os"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
)

// http 响应统一json格式
func ResponseMap(code, total int, msg, data interface{}) map[string]interface{} {

	var resMap = make(map[string]interface{})

	resMap["code"] = code
	resMap["data"] = data
	resMap["total"] = total
	resMap["message"] = msg

	return resMap
}

// 当前时间格式化
func Datetime() string {
	return time.Now().Format("2006-01-02 15:04:05")
}

// 错误日志
func ErrLog(errmsg interface{}) {

	f, _ := os.OpenFile(PathConfig.ErrLogs, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)

	msg := fmt.Sprintf("[%s] -- error : %s\n", Datetime(), errmsg)

	io.WriteString(f, msg)

	defer f.Close()
}

// jwt 生成token
func Token(uid uint32) (string, error) {

	type tokenStruct struct {
		jwt.StandardClaims
		Uid uint32
	}

	claims := &tokenStruct{
		jwt.StandardClaims{
			IssuedAt:  time.Now().Unix(),
			ExpiresAt: time.Now().Add(time.Second * CommonConfig.TokenExpired).Unix(),
		},
		uid,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	stoken, err := token.SignedString([]byte(CommonConfig.TokenSecret))

	if err != nil {
		return "", err
	}

	return stoken, nil
}

// 解密token
func Parse(tokenString string) (int, error) {

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {

		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {

			return nil, errors.New("token 解析错误")
		}

		return []byte(CommonConfig.TokenSecret), nil
	})

	if err != nil {
		return 0, err
	}

	if !token.Valid {
		return 0, errors.New("无效token")
	}

	claim, ok := token.Claims.(jwt.MapClaims)

	if !ok {
		return 0, errors.New("token转换失败")
	}

	uid := int(claim["Uid"].(float64))

	return uid, nil
}
