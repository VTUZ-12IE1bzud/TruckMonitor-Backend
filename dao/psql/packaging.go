package psql

import (
	"TruckMonitor-Backend/dao"
	"TruckMonitor-Backend/model"
	"database/sql"
)

type psqlPackaging struct {
	context PsqlContext
}

func PackagingDao(context PsqlContext) dao.PackagingDao {
	return &psqlPackaging{context}
}

func (dao *psqlPackaging) db() *sql.DB {
	return dao.context.GetDb()
}

func (dao *psqlPackaging) FindById(id int) (*model.Packaging, error) {
	var data model.Packaging
	row := dao.db().QueryRow("SELECT * FROM packaging WHERE id=$1", id)
	if err := row.Scan(&data.Id, &data.Name); err != nil {
		return nil, err
	}
	return &data, nil
}
