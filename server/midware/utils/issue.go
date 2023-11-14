package utils

import (
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"time"
)

func IssueUserToken(name string) {
	myClaims := UserClaims{
		Name: name,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			NotBefore: jwt.NewNumericDate(time.Now()),
		},
	}

	hs := HS{Key: "12345678qwertyjasdasdasdasdasd"}
	sign, err := hs.Encode(myClaims)

	fmt.Println(name, sign, err)
}
