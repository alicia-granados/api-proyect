package handlers

import (
	"ApiRest/db"
	"ApiRest/models"
	"database/sql"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
)

func GetSkins(databaseRepo *db.RealDBRepo, rw http.ResponseWriter, r *http.Request) {
	skins, err := databaseRepo.GetSkins()

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

func GetSkinsId(databaseRepo *db.RealDBRepo, rw http.ResponseWriter, r *http.Request) {

	id := chi.URLParam(r, "id")

	SkinID, err := strconv.Atoi(id)
	if err != nil {
		models.HandleError(rw, http.StatusUnprocessableEntity, "Invalid skin ID", err)
		return
	}

	existsChampion := databaseRepo.ExistsID("Skins", SkinID)

	if !existsChampion {
		models.HandleError(rw, http.StatusNotFound, "Champion not found", nil)
		return
	}

	skin, err := databaseRepo.GetSkinId(SkinID)
	if err != nil {
		models.HandleError(rw, http.StatusNotFound, "skin not found", nil)
		return
	}
	models.SendData(rw, skin, "get skin by id", http.StatusOK)
}
