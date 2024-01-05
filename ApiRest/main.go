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
	dnsServer := config.DSNServer()
	fmt.Printf("listen server: %s", dnsServer)

	log.Fatal(http.ListenAndServe(dnsServer, nil))
}
