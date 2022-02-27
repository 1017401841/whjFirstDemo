package common

import (
	"github.com/dgrijalva/jwt-go"
	"time"
	"wanghaojun.whj/whjFirstPriject/model"
)

var jwtKey = []byte("a_secret_create")

type Claims struct {
	UserId uint
	jwt.StandardClaims
}

func ReleaseToken(user model.User) (string, error) {
	expirationTime := time.Now().Add(7 * 24 * time.Hour)
	claims := &Claims{
		UserId: user.ID,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
			IssuedAt:  time.Now().Unix(),
			Issuer:    "wanghaojun",
			Subject:   "user token",
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtKey)

	if err != nil {
		return "", err
	}
	return tokenString, nil
	//eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVc2VySWQiOjQsImV4cCI6MTY0NjQ4MDI2MywiaWF0IjoxNjQ1ODc1NDYzLCJpc3MiOiJ3YW5naGFvanVuIiwic3ViIjoidXNlciB0b2tlbiJ9.FT7iQDkmzvY7dVtbyv_wU2CDYungZcjq81jcPgWmN3M
	//token 包含三部分，第一部分是协议头header，储存的是加密协议；第二部分是claims中的信息；第三部分是前两部分加key在哈希的值。
}

func ParseToken(tokenString string) (*jwt.Token, *Claims, error) {
	claims := &Claims{}
	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (i interface{}, err error) {
		return jwtKey, nil
	})
	return token, claims, err
}
