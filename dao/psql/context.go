package psql

import (
	"log"
	"database/sql"
	_ "github.com/lib/pq"
	"fmt"
)

type (
	PsqlContext interface {
		GetDb() *sql.DB
		Close()
		SchemeInit(scheme string) error
	}

	psqlContext struct {
		db *sql.DB
	}
)

func (c *psqlContext) GetDb() *sql.DB {
	return c.db
}

func (c *psqlContext) Close() {
	c.db.Close()
}

func (c *psqlContext) SchemeInit(scheme string) error {
	if _, err := c.db.Exec(scheme); err != nil {
		log.Println(err)
		return err
	}
	return nil
}

func NewConnect(host string, port string, user string, password string, dbName string, sslMode string) (PsqlContext, error) {
	connectString := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		host, port, user, password, dbName, sslMode)
	db, err := sql.Open("postgres", connectString)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	if err = db.Ping(); err != nil {
		log.Println(err)
		return nil, err
	}
	return &psqlContext{db }, nil
}
