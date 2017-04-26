package psql

import (
	"TruckMonitor-Backend/dao"
	"database/sql"
	"TruckMonitor-Backend/model"
)

type psqlEmployee struct {
	context PsqlContext
}

func EmployeeDao(context PsqlContext) dao.EmployeeDao {
	return &psqlEmployee{context}
}

func (dao *psqlEmployee) db() *sql.DB {
	return dao.context.GetDb()
}

func (dao *psqlEmployee) FindById(id int) (*model.Employee, error) {
	var data model.Employee
	row := dao.db().QueryRow("SELECT * FROM employee WHERE id=$1", id)
	err := row.Scan(&data.Id, &data.Role, &data.Surname, &data.Name, &data.Patronymic, &data.BirthDate,
		&data.Email, &data.Password, &data.Photo, &data.Phone)
	if err != nil {
		return nil, err
	}
	return &data, nil
}

func (dao *psqlEmployee) FindByEmailAndPassword(email string, password string) (*model.Employee, error) {
	var data model.Employee
	row := dao.db().QueryRow("SELECT * FROM employee WHERE email=$1 AND password=md5($2)", email, password)
	err := row.Scan(&data.Id, &data.Role, &data.Surname, &data.Name, &data.Patronymic, &data.BirthDate,
		&data.Email, &data.Password, &data.Photo, &data.Phone)
	if err != nil {
		return nil, err
	}
	return &data, nil
}
