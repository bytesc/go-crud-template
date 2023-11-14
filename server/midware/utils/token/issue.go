package token

import (
	"github.com/golang-jwt/jwt/v5"
	"time"
)

func IssueHS(name string) (string, error) {
	myClaims := UserClaims{
		Name: name,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			NotBefore: jwt.NewNumericDate(time.Now()),
		},
	}

	signature, err := hs.Encode(myClaims)
	//fmt.Println(name, signature, err)
	//bytes, _ := base64.StdEncoding.DecodeString("eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9")
	//fmt.Println(string(bytes))
	return signature, err
}

func IssueRS(name string) (string, error) {
	myClaims := UserClaims{
		Name: name,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			NotBefore: jwt.NewNumericDate(time.Now()),
		},
	}
	signature, err := rs.Encode(myClaims)
	//fmt.Println(signature, err)
	//bytes, _ := base64.StdEncoding.DecodeString(signature)
	//fmt.Println(string(bytes))
	return signature, err
}
