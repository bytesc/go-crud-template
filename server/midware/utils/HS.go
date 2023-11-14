package utils

import "github.com/golang-jwt/jwt/v5"

type JwtValidator interface {
	Encode(claims jwt.Claims) (string, error)
	Decode(sign string, claims jwt.Claims) error
}

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
