package handlers

import (
	"ApiRest/db"
	"ApiRest/models"
	"database/sql"
	"encoding/json"
	"net/http"
)

func AllChampions(databaseRepo *db.RealDBRepo, rw http.ResponseWriter, r *http.Request) {

	champions, err := databaseRepo.ListChampions()

	if err != nil {
		// Here you can add additional logic to determine the type of error and send a specific error message
		if err == sql.ErrNoRows {
			models.SendNotFound(rw)
		} else {
			models.SendUnprocessableEntity(rw)
		}
		return
	}
	models.SendData(rw, champions)
}

func CreateChampion(databaseRepo *db.RealDBRepo, rw http.ResponseWriter, r *http.Request) {

	champion := models.Champion{}
	decoder := json.NewDecoder(r.Body)

	// Decode the JSON of the request body
	if err := decoder.Decode(&champion); err != nil {
		models.SendUnprocessableEntity(rw)
		return
	}
	// Database insertion logic
	if err := databaseRepo.InsertChampion(champion.Name, champion.Title, champion.Lore); err != nil {
		models.SendUnprocessableEntity(rw)
		return
	}
	// Reply with a message
	models.SendData(rw, champion)

}
