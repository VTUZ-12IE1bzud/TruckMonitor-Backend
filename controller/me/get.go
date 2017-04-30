package me

import (
	"TruckMonitor-Backend/controller/authentication"
	"gopkg.in/gin-gonic/gin.v1"
	"log"
	"net/http"
	"time"
)

type (
	meResponse struct {
		Role       string    `json:"role"`
		Surname    string    `json:"surname"`
		Name       string    `json:"name"`
		Patronymic string    `json:"patronymic"`
		BirthDate  time.Time `json:"birthDate"`
		Email      string    `json:"email"`
		Photo      string    `json:"photo"`
		Phone      string    `json:"phone"`
	}
)

func (c *controller) GetMe(context *gin.Context) {
	employeeId := context.MustGet(authentication.PARAM_EMPLOYEE_ID).(int)
	employee, err := c.employeeDao.FindById(employeeId)
	if err != nil {
		context.AbortWithStatus(http.StatusBadGateway)
		log.Println(err)
		return
	}
	context.JSON(http.StatusOK, meResponse{
		Role:       employee.Role,
		Surname:    employee.Surname,
		Name:       employee.Name,
		Patronymic: employee.Patronymic,
		BirthDate:  employee.BirthDate,
		Email:      employee.Email,
		Photo:      employee.Photo,
		Phone:      employee.Phone,
	})
}
