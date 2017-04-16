package token

import (
	"github.com/dgrijalva/jwt-go"
	"errors"
)

type AccountToken interface {
	Create(accountId int) (string, error)
	Parse(tokenString string) (int, error)
}

type Source struct {
	Key []byte
}

type claims struct {
	AccountId int    `json:"accountId"`
	jwt.StandardClaims
}

func (src *Source) Create(accountId int) (string, error) {
	newToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims{
		AccountId: accountId, StandardClaims: jwt.StandardClaims{Issuer: "TruckMonitor"},
	})
	return newToken.SignedString(src.Key)
}

func (src *Source) Parse(tokenString string) (int, error) {
	userToken, err := jwt.ParseWithClaims(tokenString, &claims{}, func(token *jwt.Token) (interface{}, error) {
		return src.Key, nil
	})
	if err != nil {
		return -1, err
	} else {
		if claims, ok := userToken.Claims.(*claims); ok && userToken.Valid {
			return claims.AccountId, nil
		} else {
			return -1, errors.New("Invalid token")
		}
	}
}