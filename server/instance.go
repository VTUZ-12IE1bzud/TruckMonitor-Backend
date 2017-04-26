package server

import (
	"TruckMonitor-Backend/context"
	"TruckMonitor-Backend/controller"
)

type Instance struct {
	Configuration context.Configuration
}

func (instance Instance) Start() error {
	appContext, err := context.NewApplicationContext(instance.Configuration)
	if err != nil {
		return err
	}
	defer appContext.DbContext().Close()
	return controller.NewRouter(appContext).Run(instance.Configuration.ServerConfiguration.Port)
}
