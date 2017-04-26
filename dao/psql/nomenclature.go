package psql

import (
	"TruckMonitor-Backend/dao"
	"database/sql"
	"TruckMonitor-Backend/model"
)

type psqlNomenclature struct {
	context PsqlContext
}

func NomenclatureDao(context PsqlContext) dao.NomenclatureDao {
	return &psqlNomenclature{context}
}

func (dao *psqlNomenclature) db() *sql.DB {
	return dao.context.GetDb()
}

func (dao *psqlNomenclature) FindById(id int) (*model.Nomenclature, error) {
	var data model.Nomenclature
	row := dao.db().QueryRow("SELECT * FROM nomenclature WHERE id=$1", id)
	err := row.Scan(&data.Id, &data.MeasureId, &data.Name)
	if err != nil {
		return nil, err
	}
	return &data, nil
}
