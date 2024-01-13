package handlers

import (
	"ApiRest/db"
	"ApiRest/models"
	"database/sql"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
)

func GetTags(databaseRepo *db.RealDBRepo, rw http.ResponseWriter, r *http.Request) {
	tags, err := databaseRepo.GetTags()

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

	tag, err := databaseRepo.GetTagId(tagID)
	if err != nil {
		models.HandleError(rw, http.StatusNotFound, "tag not found", nil)
		return
	}
	models.SendData(rw, tag, "get tag by id", http.StatusOK)
}
