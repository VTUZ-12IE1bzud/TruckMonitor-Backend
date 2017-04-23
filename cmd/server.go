package main

import (
	"log"
	"os"
	"TruckMonitor-Backend/dao/psql"
	"strconv"
	"io/ioutil"
)

var (
	serviceHost      = os.Getenv("SERVICE_HOST")
	dbHost           = os.Getenv("DB_HOST")
	dbPort           = os.Getenv("DB_PORT")
	dbDatabase       = os.Getenv("DB_DATABASE")
	dbUser           = os.Getenv("DB_USER")
	dbPassword       = os.Getenv("DB_PASSWORD")
	dbInitialization = os.Getenv("DB_INITIALIZATION")
)

func main() {
	log.Print("[Start]")

	connect := psql.NewConnect(dbHost, dbPort, dbUser, dbPassword, dbDatabase)
	db, err := connect.New()
	if err != nil {
		log.Panic("[DB]", err)
	}
	defer db.Close()

	if isInitialization, _ := strconv.ParseBool(dbInitialization); isInitialization {
		scheme, err := ioutil.ReadFile("db.sql")
		if err != nil {
			log.Panic("[DB]", err)
		}
		if err = db.InitScheme(string(scheme)); err != nil {
			log.Panic("[DB]", err)
		}
	}

	// Route
	//_ := &token.Source{[]byte("37FjfjU7vka80OU3r520Yy2T7h0p7h7AM") }
}
