package token

import (
	"github.com/golang-jwt/jwt/v5"
	"time"
)

type RS struct {
	PublicKey  string
	PrivateKey string
}

func (rs *RS) Encode(claims jwt.Claims) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodRS256, claims)
	pKey, err := jwt.ParseRSAPrivateKeyFromPEM([]byte(rs.PrivateKey))
	if err != nil {
		return "", err
	}
	signature, err := token.SignedString(pKey)
	return signature, err
}

func (rs *RS) Decode(signature string, claims jwt.Claims) error {
	_, err := jwt.ParseWithClaims(signature, claims, func(token *jwt.Token) (interface{}, error) {
		return jwt.ParseRSAPublicKeyFromPEM([]byte(rs.PublicKey))
	})
	return err
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

func CheckRS(signature string) error {
	myClaims := UserClaims{}
	err := rs.Decode(signature, &myClaims)
	//fmt.Println(myClaims, err)
	return err
}

//signature,_ := token.IssueRS("helloword")
//fmt.Println("签名内容",signature)
//err := token.CheckRS(signature)
//fmt.Println("验签",err)
