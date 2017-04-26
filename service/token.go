package service

import (
	"TruckMonitor-Backend/dao"
	"TruckMonitor-Backend/model"
	"errors"
	"github.com/dgrijalva/jwt-go"
)

type TokenService interface {
	Create(user *model.Employee) (string, error)
	Resolve(sessionToken string) (*model.SessionToken, error)
}

type tokenService struct {
	key         []byte
	employeeDao dao.EmployeeDao
}

func NewTokenService(key string, employeeDao dao.EmployeeDao) TokenService {
	return &tokenService{
		key:         []byte(key),
		employeeDao: employeeDao,
	}
}

func (s *tokenService) Create(user *model.Employee) (string, error) {
	token := jwt.NewWithClaims(
		jwt.SigningMethodHS256,
		model.SessionToken{
			EmployeeId:    user.Id,
			EmployeeEmail: user.Email,
			StandardClaims: jwt.StandardClaims{
				Issuer: "TruckMonitor",
			},
		})
	return token.SignedString(s.key)
}

func (s *tokenService) Resolve(sessionToken string) (*model.SessionToken, error) {
	userToken, err := jwt.ParseWithClaims(sessionToken, &model.SessionToken{}, func(token *jwt.Token) (interface{}, error) {
		return s.key, nil
	})
	if err != nil {
		return nil, err
	} else {
		if claims, ok := userToken.Claims.(*model.SessionToken); ok && userToken.Valid {
			return claims, nil
		} else {
			return nil, errors.New("Invalid token")
		}
	}
}
