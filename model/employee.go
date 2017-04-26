package model

import "time"

// Роль сотрудника
const (
	ADMIN      string = "admin"
	ACCOUNTANT string = "accountant"
	MANAGER    string = "manager"
	DRIVER     string = "driver"
)

// Сотрудник
type Employee struct {
	Id         int
	Role       string
	Surname    string
	Name       string
	Patronymic string
	BirthDate  time.Time
	Email      string
	Password   string
	Photo      string
	Phone      string
}
