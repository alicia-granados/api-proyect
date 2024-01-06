package main

import (
	"fmt"
	"log"
	"net/http"

	"ApiRest/config"
	"ApiRest/db"
	"ApiRest/routes"
)

func main() {

	databaseRepo := db.RealDBRepo{}

	databaseRepo.Connect()
	defer databaseRepo.Close()

	dnsServer, err := config.DSNServer()
	fmt.Printf("listen server: %s", dnsServer)
	if err != nil {
		log.Fatal("Error configuring server:", err)
	}

	err = http.ListenAndServe(dnsServer, routes.Routes())
	if err != nil {
		log.Fatal(err)
	}
}
