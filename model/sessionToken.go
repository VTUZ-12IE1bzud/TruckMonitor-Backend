package model

import "github.com/dgrijalva/jwt-go"

// Пользовательский токен
type SessionToken struct {
	EmployeeId    int    `json:"employeeId"`
	EmployeeEmail string `json:"employeeEmail"`
	jwt.StandardClaims
}
