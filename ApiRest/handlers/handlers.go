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

func HandlerGetChampionsDetailList(databaseRepo db.DatabaseRepo, rw http.ResponseWriter, r *http.Request) {
	champions, err := databaseRepo.GetChampionsDetailList()

	if err != nil {
		if err == sql.ErrNoRows {
			// No se encontraron filas, devolver 404 Not Found
			models.HandleError(rw, http.StatusNotFound, "no champions were found", nil)
		} else {
			// Otro tipo de error, devolver 500 Internal Server Error o manejar según el caso
			models.HandleError(rw, http.StatusInternalServerError, "an unexpected error occurred while retrieving champions ", err)
		}
		return
	}
	models.SendData(rw, champions, "info champions", http.StatusOK)
}

func HandlerGetChampionDetailsByID(databaseRepo db.DatabaseRepo, rw http.ResponseWriter, r *http.Request) {

	id := chi.URLParam(r, "id")

	championID, err := strconv.Atoi(id)
	if err != nil {
		models.HandleError(rw, http.StatusUnprocessableEntity, "Invalid champion ID", err)
		return
	}
	existsChampion := databaseRepo.ExistsID("Champion", championID)

	if !existsChampion {
		models.HandleError(rw, http.StatusNotFound, "Champion not found", nil)
		return
	}

	champion, err := databaseRepo.GetChampionDetailsByID(championID)
	if err != nil {
		models.HandleError(rw, http.StatusNotFound, "Champion not found", nil)
		return
	}
	models.SendData(rw, champion, "get list champion by id", http.StatusOK)
}

func HandlerInsertChampion(databaseRepo db.DatabaseRepo, rw http.ResponseWriter, r *http.Request) {

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

func HandlerUpdateChampionByID(databaseRepo db.DatabaseRepo, rw http.ResponseWriter, r *http.Request) {

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
	if err := databaseRepo.UpdateChampionByID(championID, champion); err != nil {
		models.HandleError(rw, http.StatusUnprocessableEntity, "Error updating champion into the database", err)
		return
	}
	models.SendData(rw, champion, "Updated champion", http.StatusOK)

}

func HandlerDeleteChampionByID(databaseRepo db.DatabaseRepo, rw http.ResponseWriter, r *http.Request) {

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

	if err := databaseRepo.DeleteChampionByID(championID); err != nil {
		models.HandleError(rw, http.StatusUnprocessableEntity, "Error deleting champion into the database", err)
		return
	}
	models.SendData(rw, championID, "deleted champion", http.StatusOK)

}

func HandlerGetChampions(databaseRepo db.DatabaseRepo, rw http.ResponseWriter, r *http.Request) {
	champions, err := databaseRepo.GetChampions()

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
	models.SendData(rw, champions, "get champions", http.StatusOK)
}

func HandlerGetChampionById(databaseRepo db.DatabaseRepo, rw http.ResponseWriter, r *http.Request) {

	id := chi.URLParam(r, "id")

	championID, err := strconv.Atoi(id)
	if err != nil {
		models.HandleError(rw, http.StatusUnprocessableEntity, "Invalid champion ID", err)
		return
	}
	existsChampion := databaseRepo.ExistsID("Champion", championID)

	if !existsChampion {
		models.HandleError(rw, http.StatusNotFound, "Champion not found", nil)
		return
	}

	champion, err := databaseRepo.GetChampionByID(championID)
	if err != nil {
		models.HandleError(rw, http.StatusNotFound, "Champion not found", nil)
		return
	}
	models.SendData(rw, champion, "get champion by id", http.StatusOK)
}
