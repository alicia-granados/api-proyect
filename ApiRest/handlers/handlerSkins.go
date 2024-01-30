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

func HandlerGetSkinList(databaseRepo db.DatabaseRepo, rw http.ResponseWriter, r *http.Request) {
	skins, err := databaseRepo.GetSkinList()

	if err != nil {
		if err == sql.ErrNoRows {
			// No se encontraron filas, devolver 404 Not Found
			models.HandleError(rw, http.StatusNotFound, "no skins were found", nil)
		} else {
			// Otro tipo de error, devolver 500 Internal Server Error o manejar según el caso
			models.HandleError(rw, http.StatusInternalServerError, "an unexpected error occurred while retrieving skins", err)
		}
		return
	}
	models.SendData(rw, skins, "get skins", http.StatusOK)
}

func HandlerGetSkinById(databaseRepo db.DatabaseRepo, rw http.ResponseWriter, r *http.Request) {

	id := chi.URLParam(r, "id")

	SkinID, err := strconv.Atoi(id)
	if err != nil {
		models.HandleError(rw, http.StatusUnprocessableEntity, "Invalid skin ID", err)
		return
	}

	existsSkins := databaseRepo.ExistsID("Skins", SkinID)

	if !existsSkins {
		models.HandleError(rw, http.StatusNotFound, "Skin not found", nil)
		return
	}

	skin, err := databaseRepo.GetSkinByID(SkinID)
	if err != nil {
		models.HandleError(rw, http.StatusNotFound, "skin not found", nil)
		return
	}
	models.SendData(rw, skin, "get skin by id", http.StatusOK)
}

func HandlerInsertSkin(databaseRepo db.DatabaseRepo, rw http.ResponseWriter, r *http.Request) {

	skin := models.Skins{}
	decoder := json.NewDecoder(r.Body)

	// Decode the JSON of the request body
	if err := decoder.Decode(&skin); err != nil {
		models.HandleError(rw, http.StatusUnprocessableEntity, "Error decoding JSON", err)
		return
	}

	//	verify if idchampion exists
	existsChampion := databaseRepo.ExistsID("Champion", skin.IdChampion)

	if !existsChampion {
		models.HandleError(rw, http.StatusNotFound, "Champion not found", nil)
		return
	}

	// Database insertion logic
	if err := databaseRepo.InsertSkin(skin.IdNum, skin.Num, skin.IdChampion, skin.Name); err != nil {
		models.HandleError(rw, http.StatusUnprocessableEntity, "Error inserting skins into the database", err)
		return
	}
	// Reply with a message
	models.SendData(rw, skin, "Skin created", http.StatusOK)

}

func HandlerUpdateSkinByID(databaseRepo db.DatabaseRepo, rw http.ResponseWriter, r *http.Request) {

	id := chi.URLParam(r, "id")
	skinID, err := strconv.Atoi(id)
	if err != nil {
		models.HandleError(rw, http.StatusNotFound, "Invalid champion ID", err)
		return
	}

	skin := models.Skins{}
	decoder := json.NewDecoder(r.Body)

	// Decode the JSON of the request body
	if err := decoder.Decode(&skin); err != nil {
		models.HandleError(rw, http.StatusUnprocessableEntity, "Error decoding JSON", err)
		return
	}
	existsChampion := databaseRepo.ExistsID("Champion", skin.IdChampion)
	if !existsChampion {
		models.HandleError(rw, http.StatusNotFound, "Champion not found", nil)
		return
	}

	existsSkin := databaseRepo.ExistsID("Skins", skinID)

	if !existsSkin {
		models.HandleError(rw, http.StatusNotFound, "Skins not found", nil)
		return
	}

	// Database updated logic
	if err := databaseRepo.UpdateSkinByID(skinID, skin); err != nil {
		models.HandleError(rw, http.StatusUnprocessableEntity, "Error updating skin into the database", err)
		return
	}
	models.SendData(rw, skin, "Updated skin", http.StatusOK)

}

func HandlerDeleteSkinByID(databaseRepo db.DatabaseRepo, rw http.ResponseWriter, r *http.Request) {

	id := chi.URLParam(r, "id")
	skinID, err := strconv.Atoi(id)
	if err != nil {
		models.HandleError(rw, http.StatusNotFound, "Invalid champion ID", err)

		return
	}
	existsSkin := databaseRepo.ExistsID("Skins", skinID)

	if !existsSkin {
		models.HandleError(rw, http.StatusNotFound, "skin not found", nil)
		return
	}

	if err := databaseRepo.DeleteSkinByID(skinID); err != nil {
		models.HandleError(rw, http.StatusUnprocessableEntity, "Error deleting skin into the database", err)
		return
	}
	models.SendData(rw, skinID, "deleted skin ", http.StatusOK)

}
