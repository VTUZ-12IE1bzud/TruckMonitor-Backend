package model

import "time"

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
