package handlers

import (
	"ApiRest/db"
	"ApiRest/models"
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
)

func AllChampions(databaseRepo *db.RealDBRepo, rw http.ResponseWriter, r *http.Request) {
	champions, err := databaseRepo.ListChampions()

	if err != nil {
		if err == sql.ErrNoRows {
			// No se encontraron filas, devolver 404 Not Found
			models.SendNotFound(rw, "no champions were found.")
		} else {
			// Otro tipo de error, devolver 500 Internal Server Error o manejar según el caso
			models.SendInternalServerError(rw, "an unexpected error occurred while retrieving champions.")
			fmt.Println("Error:", err)
		}
		return
	}

	models.SendData(rw, champions)
}

func GetChampionId(databaseRepo *db.RealDBRepo, rw http.ResponseWriter, r *http.Request) {

	id := chi.URLParam(r, "id")

	championID, err := strconv.Atoi(id)
	if err != nil {
		models.SendUnprocessableEntity(rw, "Invalid champion ID")
		return
	}

	champion, err := databaseRepo.GetChampionId(championID)
	if err != nil {
		models.SendNotFound(rw, "Champion not found")
		return
	}
	models.SendData(rw, champion)
}

func CreateChampion(databaseRepo *db.RealDBRepo, rw http.ResponseWriter, r *http.Request) {

	champion := models.Champion{}
	decoder := json.NewDecoder(r.Body)

	// Decode the JSON of the request body
	if err := decoder.Decode(&champion); err != nil {
		models.SendUnprocessableEntity(rw, "Error decoding JSON")
		return
	}
	// Database insertion logic
	if err := databaseRepo.InsertChampion(champion.Name, champion.Title, champion.Lore); err != nil {
		models.SendUnprocessableEntity(rw, "Error inserting champion into the database")
		return
	}
	// Reply with a message
	models.SendData(rw, champion)

}
