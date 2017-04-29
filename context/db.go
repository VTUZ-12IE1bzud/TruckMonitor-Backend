package context

import (
	"TruckMonitor-Backend/dao/psql"
	"io/ioutil"
	"log"
)

type DbContext interface {
	psql.PsqlContext
}

func NewDbContext(configuration Configuration) DbContext {
	config := configuration.PsqlConfiguration
	db := psql.NewConnect(config.Host, config.Port, config.User, config.Password,
		config.DbName, config.SslMode)

	// Initialization DB
	if config.IsCreate {
		scheme, err := ioutil.ReadFile("db.sql")
		if err != nil {
			log.Panic(err)
		}
		if err = db.SchemeInit(string(scheme)); err != nil {
			log.Panic(err)
		}
	}

	return db
}
