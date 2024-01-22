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

func GetTags(databaseRepo *db.RealDBRepo, rw http.ResponseWriter, r *http.Request) {
	tags, err := databaseRepo.GetTagsList()

	if err != nil {
		if err == sql.ErrNoRows {
			// No se encontraron filas, devolver 404 Not Found
			models.HandleError(rw, http.StatusNotFound, "no tags were found", nil)
		} else {
			// Otro tipo de error, devolver 500 Internal Server Error o manejar según el caso
			models.HandleError(rw, http.StatusInternalServerError, "an unexpected error occurred while retrieving tags", err)
		}
		return
	}
	models.SendData(rw, tags, "get tags", http.StatusOK)
}

func GetTagId(databaseRepo *db.RealDBRepo, rw http.ResponseWriter, r *http.Request) {

	id := chi.URLParam(r, "id")

	tagID, err := strconv.Atoi(id)
	if err != nil {
		models.HandleError(rw, http.StatusUnprocessableEntity, "Invalid tag ID", err)
		return
	}

	existsTag := databaseRepo.ExistsID("Tags", tagID)

	if !existsTag {
		models.HandleError(rw, http.StatusNotFound, "tag not found", nil)
		return
	}

	tag, err := databaseRepo.GetTagByID(tagID)
	if err != nil {
		models.HandleError(rw, http.StatusNotFound, "tag not found", nil)
		return
	}
	models.SendData(rw, tag, "get tag by id", http.StatusOK)
}

func CreateTag(databaseRepo *db.RealDBRepo, rw http.ResponseWriter, r *http.Request) {

	tag := models.Tags{}
	decoder := json.NewDecoder(r.Body)

	// Decode the JSON of the request body
	if err := decoder.Decode(&tag); err != nil {
		models.HandleError(rw, http.StatusUnprocessableEntity, "Error decoding JSON", err)
		return
	}

	//	verify if idchampion exists
	existsChampion := databaseRepo.ExistsID("Champion", tag.IdChampion)

	if !existsChampion {
		models.HandleError(rw, http.StatusNotFound, "Champion not found", nil)
		return
	}

	// Database insertion logic
	if err := databaseRepo.InsertTag(tag.IdChampion, tag.Name); err != nil {
		models.HandleError(rw, http.StatusUnprocessableEntity, "Error inserting Tags into the database", err)
		return
	}
	// Reply with a message
	models.SendData(rw, tag, "Tag created", http.StatusOK)

}

func PutTag(databaseRepo *db.RealDBRepo, rw http.ResponseWriter, r *http.Request) {

	id := chi.URLParam(r, "id")
	tagID, err := strconv.Atoi(id)
	if err != nil {
		models.HandleError(rw, http.StatusNotFound, "Invalid Tad ID", err)
		return
	}

	tag := models.Tags{}
	decoder := json.NewDecoder(r.Body)

	// Decode the JSON of the request body
	if err := decoder.Decode(&tag); err != nil {
		models.HandleError(rw, http.StatusUnprocessableEntity, "Error decoding JSON", err)
		return
	}
	existsChampion := databaseRepo.ExistsID("Champion", tag.IdChampion)
	if !existsChampion {
		models.HandleError(rw, http.StatusNotFound, "Champion not found", nil)
		return
	}

	existstag := databaseRepo.ExistsID("Tags", tagID)

	if !existstag {
		models.HandleError(rw, http.StatusNotFound, "tags not found", nil)
		return
	}

	// Database updated logic
	if err := databaseRepo.UpdateTagByID(tagID, tag); err != nil {
		models.HandleError(rw, http.StatusUnprocessableEntity, "Error updating tag into the database", err)
		return
	}
	models.SendData(rw, tag, "Updated tag", http.StatusOK)

}

func DeleteTag(databaseRepo *db.RealDBRepo, rw http.ResponseWriter, r *http.Request) {

	id := chi.URLParam(r, "id")
	tagID, err := strconv.Atoi(id)
	if err != nil {
		models.HandleError(rw, http.StatusNotFound, "Invalid tag ID", err)

		return
	}
	existsTag := databaseRepo.ExistsID("Tags", tagID)

	if !existsTag {
		models.HandleError(rw, http.StatusNotFound, "tag not found", nil)
		return
	}

	if err := databaseRepo.DeleteTagByID(tagID); err != nil {
		models.HandleError(rw, http.StatusUnprocessableEntity, "Error deleting skin into the database", err)
		return
	}
	models.SendData(rw, tagID, "deleted skin ", http.StatusOK)

}
