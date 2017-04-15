package token

import (
	"github.com/dgrijalva/jwt-go"
	"errors"
)

type TokenParser interface {
	Create(accountId int) (string, error)
	Parse(tokenString string) (int, error)
}

type AccountToken struct {
	Key []byte
}

type claims struct {
	AccountId int    `json:"accountId"`
	jwt.StandardClaims
}

func (token *AccountToken) Create(accountId int) (string, error) {
	newToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims{
		AccountId: accountId, StandardClaims: jwt.StandardClaims{Issuer: "TruckMonitor"},
	})
	return newToken.SignedString(token.Key)
}

func (a *AccountToken) Parse(tokenString string) (int, error) {
	userToken, err := jwt.ParseWithClaims(tokenString, &claims{}, func(token *jwt.Token) (interface{}, error) {
		return a.Key, nil
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