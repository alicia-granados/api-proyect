package handlers

import (
	"fmt"
	"net/http"
)

func AllChampions(rw http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(rw, "lista de todos los campeones")
}
