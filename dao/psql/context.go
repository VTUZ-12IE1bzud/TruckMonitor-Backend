package psql

import (
	"log"
	"database/sql"
	_ "github.com/lib/pq"
	"fmt"
)

type (
	Connect interface {
		New() (*context, error)
	}

	connect struct {
		host     string
		port     string
		user     string
		password string
		dbName   string
		sslMode  string
	}
)

type (
	Context interface {
		Close()
		InitScheme(scheme string) error
	}

	context struct {
		Context
		*sql.DB
	}
)

func NewConnect(host string, port string, user string, password string, dbName string) Connect {
	return Connect(&connect{
		host:     host,
		port:     port,
		user:     user,
		password: password,
		dbName:   dbName,
		sslMode:  "disable",
	})
}

func (c *connect) New() (*context, error) {
	connectString := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		c.host, c.port, c.user, c.password, c.dbName, c.sslMode)
	db, err := sql.Open("postgres", connectString)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	if err = db.Ping(); err != nil {
		log.Println(err)
		return nil, err
	}
	return &context{Context(db), db}, nil
}

func (c *context) Close() {
	c.DB.Close()
}

func (c *context) InitScheme(scheme string) error {
	if _, err := c.DB.Exec(scheme); err != nil {
		log.Println(err)
		return err
	}
	return nil
}
