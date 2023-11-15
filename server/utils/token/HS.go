package token

import (
	"github.com/golang-jwt/jwt/v5"
	"time"
)

//type JwtValidator interface {
//	Encode(claims jwt.Claims) (string, error)
//	Decode(sign string, claims jwt.Claims) error
//}

type HS struct {
	Key string
}

func (hs *HS) Encode(claims jwt.Claims) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	sign, err := token.SignedString([]byte(hs.Key))
	return sign, err
}

func (hs *HS) Decode(sign string, claims jwt.Claims) error {
	_, err := jwt.ParseWithClaims(sign, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(hs.Key), nil
	})
	return err
}

func IssueHS(data interface{}, expTime time.Time) (string, error) {
	myClaims := UserClaims{
		Data: data,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expTime),
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

func CheckHS(signature string) error {
	myClaims := UserClaims{}
	err := hs.Decode(signature, &myClaims)
	//fmt.Println(myClaims, err)
	return err
}

//signature, _ := token.IssueHS("hello")
//fmt.Println("签名内容",signature)
//err := token.CheckHS(signature)
//fmt.Println("验签",err)
