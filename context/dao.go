package context

import (
	"TruckMonitor-Backend/dao"
	"TruckMonitor-Backend/dao/psql"
)

type DaoContext interface {
	CarriageDao() dao.CarriageDao
	CheckPointDao() dao.CheckPointDao
	ClientDao() dao.ClientDao
	ContractDao() dao.ContractDao
	EmployeeDao() dao.EmployeeDao
	MeasureDao() dao.MeasureDao
	NomenclatureDao() dao.NomenclatureDao
	PackagingDao() dao.PackagingDao
	StoreDao() dao.StoreDao
	VehicleDao() dao.VehicleDao
}

type daoContext struct {
	carriageDao     dao.CarriageDao
	checkPointDao   dao.CheckPointDao
	clientDao       dao.ClientDao
	contractDao     dao.ContractDao
	employeeDao     dao.EmployeeDao
	measureDao      dao.MeasureDao
	nomenclatureDao dao.NomenclatureDao
	packagingDao    dao.PackagingDao
	storeDao        dao.StoreDao
	vehicleDao      dao.VehicleDao
}

func (dao *daoContext) CarriageDao() dao.CarriageDao {
	return dao.carriageDao
}

func (dao *daoContext) CheckPointDao() dao.CheckPointDao {
	return dao.checkPointDao
}

func (dao *daoContext) ClientDao() dao.ClientDao {
	return dao.clientDao
}

func (dao *daoContext) ContractDao() dao.ContractDao {
	return dao.contractDao
}

func (dao *daoContext) EmployeeDao() dao.EmployeeDao {
	return dao.employeeDao
}

func (dao *daoContext) MeasureDao() dao.MeasureDao {
	return dao.measureDao
}

func (dao *daoContext) NomenclatureDao() dao.NomenclatureDao {
	return dao.nomenclatureDao
}

func (dao *daoContext) PackagingDao() dao.PackagingDao {
	return dao.packagingDao
}

func (dao *daoContext) StoreDao() dao.StoreDao {
	return dao.storeDao
}

func (dao *daoContext) VehicleDao() dao.VehicleDao {
	return dao.vehicleDao
}

func NewDaoContext(psqlContext psql.PsqlContext) (DaoContext) {
	return &daoContext{
		carriageDao:     psql.CarriageDao(psqlContext),
		checkPointDao:   psql.CheckPointDao(psqlContext),
		clientDao:       psql.ClientDao(psqlContext),
		contractDao:     psql.ContractDao(psqlContext),
		employeeDao:     psql.EmployeeDao(psqlContext),
		measureDao:      psql.MeasureDao(psqlContext),
		nomenclatureDao: psql.NomenclatureDao(psqlContext),
		packagingDao:    psql.PackagingDao(psqlContext),
		storeDao:        psql.StoreDao(psqlContext),
		vehicleDao:      psql.VehicleDao(psqlContext),
	}
}
