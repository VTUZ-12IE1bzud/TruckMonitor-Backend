package context

import (
	"TruckMonitor-Backend/dao/psql"
	"io/ioutil"
)

type DbContext interface {
	psql.PsqlContext
}

func NewDbContext(configuration Configuration) (DbContext, error) {
	config := configuration.PsqlConfiguration
	db, err := psql.NewConnect(config.Host, config.Port, config.User, config.Password,
		config.DbName, config.SslMode)
	if err != nil {
		return nil, err
	}

	// Initialization DB
	if config.IsCreate {
		scheme, err := ioutil.ReadFile("db.sql")
		if err != nil {
			return nil, err
		}
		if err = db.SchemeInit(string(scheme)); err != nil {
			return nil, err
		}
	}

	return db, nil
}
