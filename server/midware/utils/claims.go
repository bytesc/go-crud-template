package utils

import "github.com/golang-jwt/jwt/v5"

type UserClaims struct {
	jwt.RegisteredClaims
	Name string
}

func (mc *UserClaims) Valid() error {
	//字段校验
	return nil
}
