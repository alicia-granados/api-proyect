package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "¡Hello, world!")
	})

	// Especificar el puerto
	PORT := ":8080"
	fmt.Printf("Server running on port %s\n", PORT)

	// Iniciar el servidor
	log.Fatal(http.ListenAndServe(PORT, nil))
}
