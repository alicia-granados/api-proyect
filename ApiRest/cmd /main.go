package main

import (
	"fmt"
	"log"
	"net/http"

	"ApiRest/config"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "¡Hello, world!")
	})

	dnsServer, err := config.DSNServer()
	fmt.Printf("listen server: %s", dnsServer)
	if err != nil {
		log.Fatal("Error configuring server:", err)
	}

	log.Fatal(http.ListenAndServe(dnsServer, nil))
}
