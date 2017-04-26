package psql

import (
	"TruckMonitor-Backend/dao"
	"database/sql"
	"TruckMonitor-Backend/model"
)

type psqlMeasure struct {
	context PsqlContext
}

func MeasureDao(context PsqlContext) dao.MeasureDao {
	return &psqlMeasure{context}
}

func (dao *psqlMeasure) db() *sql.DB {
	return dao.context.GetDb()
}

func (dao * psqlMeasure) FindById(id int) (*model.Measure, error) {
	var data model.Measure
	row := dao.db().QueryRow("SELECT * FROM measure WHERE id=$1", id)
	err := row.Scan(&data.Id, &data.Name, &data.Abbreviation)
	if err != nil {
		return nil, err
	}
	return &data, nil
}