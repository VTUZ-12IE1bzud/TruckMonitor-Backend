package main

import (
	"log"
	"TruckMonitor-Backend/context"
	"TruckMonitor-Backend/server"
)

func main() {
	log.Print("[Start TruckMonitor]")
	config := context.Build()

	apiServer := server.Instance{Configuration: *config }
	err := apiServer.Start()
	log.Panicln(err)
}
