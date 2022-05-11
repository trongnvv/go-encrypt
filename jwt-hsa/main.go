package main

import (
	"errors"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"reflect"
)

type CustomCalim struct {
	jwt.StandardClaims
	ID uint64 `json:"id"`
}

var secret = "secret"

func encrypt(id uint64) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, CustomCalim{ID: id})
	return token.SignedString([]byte(secret))
}

func decrypt(tokenString string) (*CustomCalim, error) {
	token, err := jwt.ParseWithClaims(tokenString, &CustomCalim{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(secret), nil
	})
	if err != nil {
		return nil, err
	}
	claims, ok := token.Claims.(*CustomCalim)
	if ok && token.Valid {
		return claims, nil
	} else {
		return nil, errors.New("invalid")
	}
}

func main() {
	stringToken, err := encrypt(12345)
	if err != nil {
		panic(err)
	}
	fmt.Println("stringToken", stringToken)
	claims, err := decrypt(stringToken)
	if err != nil {
		panic(err)
	}
	fmt.Println(reflect.TypeOf(claims.ID))
}
