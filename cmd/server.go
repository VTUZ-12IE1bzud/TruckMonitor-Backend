package main

import (
	"log"
	"os"
	"TruckMonitor-Backend/psql"
	"TruckMonitor-Backend/service"
	"fmt"
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

	// Data Base
	db, err := psql.NewConnect(fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		dbHost, dbPort, dbUser, dbPassword, dbDatabase))
	if err != nil {
		log.Panic("[DB]", err)
	}
	defer db.Close()
	// Data Base Initialization
	if isInitialization, _ := strconv.ParseBool(dbInitialization); isInitialization {
		log.Print("[DB] Scheme Initialization")
		var scheme []byte
		scheme, err = ioutil.ReadFile("db.sql")
		if err != nil {
			log.Panic("[DB]", err)
		}
		if err = db.Scheme(string(scheme)); err != nil {
			log.Panic("[DB]", err)
		}
	}

	// Route
	env := &service.Environment{db}
	appService := &service.AppService{serviceHost, env}
	appService.Run()
}
