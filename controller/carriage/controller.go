package carriage

import (
	"TruckMonitor-Backend/context"
)

type controller struct {
	dao context.DaoContext
}

func Controller(dao context.DaoContext) *controller {
	return &controller{
		dao: dao,
	}
}
