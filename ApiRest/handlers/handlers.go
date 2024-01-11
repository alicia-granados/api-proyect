package handlers

import (
	"ApiRest/db"
	"ApiRest/models"
	"database/sql"
	"net/http"
)

func AllChampions(databaseRepo *db.RealDBRepo, rw http.ResponseWriter, r *http.Request) {

	champions, err := databaseRepo.ListChampions()
	if err != nil {
		// Aquí puedes agregar lógica adicional para determinar el tipo de error y enviar un mensaje de error específico
		if err == sql.ErrNoRows {
			models.SendNotFound(rw)
		} else {
			models.SendUnprocessableEntity(rw)
		}
		return
	}
	models.SendData(rw, champions)
}
