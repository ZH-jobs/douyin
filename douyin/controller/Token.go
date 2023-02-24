package controller

import (
	"fmt"
	"time"
	"github.com/dgrijalva/jwt-go"
)

var jwtkey = []byte("xiaobaibaotuandui")

type Claims struct {
	UserId int
	jwt.StandardClaims
}

func SettingToken(ID int) string {
	expireTime := time.Now().Add(7 * 24 * time.Hour) //过期时间
	claims := &Claims{
		UserId: ID,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expireTime.Unix(), //过期时间
			IssuedAt:  time.Now().Unix(),
			Subject:   "user token", //签名主题
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtkey)
	if err != nil {
		fmt.Println(err)
	}
	return tokenString
}

//解析user ID 返回-1(解析出错,应该返回错误json)
func IsTrueToken(GetinToken string) int {
	if GetinToken == "" {
		return -1
	}
	token, claims, err := ParseToken(GetinToken)
	if err != nil || !token.Valid {
		return -1
	}
	return claims.UserId
}

//Token验证函数
func ParseToken(tokenString string) (*jwt.Token, *Claims, error) {
	Claims := &Claims{}
	token, err := jwt.ParseWithClaims(tokenString, Claims, func(token *jwt.Token) (i interface{}, err error) {
		return jwtkey, nil
	})
	return token, Claims, err
}
