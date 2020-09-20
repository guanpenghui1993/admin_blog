package utils

import (
	"time"

	"github.com/dgrijalva/jwt-go"
)

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
