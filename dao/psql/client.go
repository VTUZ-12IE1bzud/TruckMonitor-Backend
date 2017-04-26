package psql

import (
	"TruckMonitor-Backend/dao"
	"database/sql"
	"TruckMonitor-Backend/model"
)

type psqlClient struct {
	context PsqlContext
}

func ClientDao(context PsqlContext) dao.ClientDao {
	return &psqlClient{context}
}

func (dao *psqlClient) db() *sql.DB {
	return dao.context.GetDb()
}

func (dao * psqlClient) FindById(id int) (*model.Client, error) {
	var data model.Client
	row := dao.db().QueryRow("SELECT * FROM client WHERE id=$1", id)
	err := row.Scan(&data.Id, &data.Name, &data.Itn, &data.Iec, &data.Address)
	if err != nil {
		return nil, err
	}
	return &data, nil
}