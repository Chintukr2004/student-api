package utils

import (
	
	"errors"

	"github.com/golang-jwt/jwt/v5"
)

func ParseToken(tokenStr, secret string) (*Claims, error){
	token, err := jwt.ParseWithClaims(
		tokenStr,
		&Claims{},
		func(t *jwt.Token) (interface{}, error){
			return []byte(secret), nil
		},
	)
	if  err != nil || !token.Valid{
		return nil, errors.New("invalid token")
	}

	claims, ok := token.Claims.(*Claims)
	if !ok {
		return nil, errors.New("invalid claims")

	}
	return claims, nil

}