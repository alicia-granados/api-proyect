package handlers

import (
	"ApiRest/db"
	"ApiRest/models"
	"database/sql"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
)

func AllInfoChampions(databaseRepo *db.RealDBRepo, rw http.ResponseWriter, r *http.Request) {
	champions, err := databaseRepo.AllInfoChampions()

	if err != nil {
		if err == sql.ErrNoRows {
			// No se encontraron filas, devolver 404 Not Found
			models.HandleError(rw, http.StatusNotFound, "no champions were found", nil)
		} else {
			// Otro tipo de error, devolver 500 Internal Server Error o manejar según el caso
			models.HandleError(rw, http.StatusInternalServerError, "an unexpected error occurred while retrieving champions", err)
		}
		return
	}
	models.SendData(rw, champions, "info champions", http.StatusOK)
}

func GetInfoChampionId(databaseRepo *db.RealDBRepo, rw http.ResponseWriter, r *http.Request) {

	id := chi.URLParam(r, "id")

	championID, err := strconv.Atoi(id)
	if err != nil {
		models.HandleError(rw, http.StatusUnprocessableEntity, "Invalid champion ID", err)
		return
	}

	champion, err := databaseRepo.GetInfoChampionId(championID)
	if err != nil {
		models.HandleError(rw, http.StatusNotFound, "Champion not found", nil)
		return
	}
	models.SendData(rw, champion, "get list champion by id", http.StatusOK)
}

func CreateChampion(databaseRepo *db.RealDBRepo, rw http.ResponseWriter, r *http.Request) {

	champion := models.Champion{}
	decoder := json.NewDecoder(r.Body)

	// Decode the JSON of the request body
	if err := decoder.Decode(&champion); err != nil {
		models.HandleError(rw, http.StatusUnprocessableEntity, "Error decoding JSON", err)
		return
	}
	// Database insertion logic
	if err := databaseRepo.InsertChampion(champion.Name, champion.Title, champion.Lore); err != nil {
		models.HandleError(rw, http.StatusUnprocessableEntity, "Error inserting champion into the database", err)
		return
	}
	// Reply with a message
	models.SendData(rw, champion, "Champion created", http.StatusOK)

}

func PutChampion(databaseRepo *db.RealDBRepo, rw http.ResponseWriter, r *http.Request) {

	id := chi.URLParam(r, "id")
	championID, err := strconv.Atoi(id)
	if err != nil {
		models.HandleError(rw, http.StatusNotFound, "Invalid champion ID", err)
		return
	}

	existsChampion := databaseRepo.ExistsID("Champion", championID)

	if !existsChampion {
		models.HandleError(rw, http.StatusNotFound, "Champion not found", nil)
		return
	}

	champion := models.Champion{}
	decoder := json.NewDecoder(r.Body)

	// Decode the JSON of the request body
	if err := decoder.Decode(&champion); err != nil {
		models.HandleError(rw, http.StatusUnprocessableEntity, "Error decoding JSON", err)
		return
	}
	// Database updated logic
	if err := databaseRepo.UpdateChampion(championID, champion); err != nil {
		models.HandleError(rw, http.StatusUnprocessableEntity, "Error updating champion into the database", err)
		return
	}
	models.SendData(rw, champion, "Updated champion", http.StatusOK)

}

func DeleteChampion(databaseRepo *db.RealDBRepo, rw http.ResponseWriter, r *http.Request) {

	id := chi.URLParam(r, "id")
	championID, err := strconv.Atoi(id)
	if err != nil {
		models.HandleError(rw, http.StatusNotFound, "Invalid champion ID", err)

		return
	}
	existsChampion := databaseRepo.ExistsID("Champion", championID)

	if !existsChampion {
		models.HandleError(rw, http.StatusNotFound, "Champion not found", nil)
		return
	}

	if err := databaseRepo.DeleteChampion(championID); err != nil {
		models.HandleError(rw, http.StatusUnprocessableEntity, "Error deleting champion into the database", err)
		return
	}
	models.SendData(rw, championID, "deleted champion", http.StatusOK)

}
