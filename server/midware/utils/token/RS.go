package token

import "github.com/golang-jwt/jwt/v5"

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
