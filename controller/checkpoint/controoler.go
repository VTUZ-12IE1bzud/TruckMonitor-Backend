package checkpoint

import (
	"TruckMonitor-Backend/dao"
)

type controller struct {
	carriageDao dao.CarriageDao
}

func Controller(carriageDao dao.CarriageDao) *controller {
	return &controller{
		carriageDao: carriageDao,
	}
}
