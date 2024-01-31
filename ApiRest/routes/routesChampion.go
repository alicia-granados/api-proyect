package routes

import (
	"ApiRest/db"
	"ApiRest/handlers"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func championRoutes(router chi.Router, databaseRepo db.DatabaseRepo) {
	router.Get("/info/", func(w http.ResponseWriter, r *http.Request) {
		handlers.HandlerGetChampionsDetailList(databaseRepo, w, r)
	})
	router.Get("/info/{id:[0-9]+}", func(w http.ResponseWriter, r *http.Request) {
		handlers.HandlerGetChampionDetailsByID(databaseRepo, w, r)
	})

	router.Get("/", func(w http.ResponseWriter, r *http.Request) {
		handlers.HandlerGetChampions(databaseRepo, w, r)
	})
	router.Get("/{id:[0-9]+}", func(w http.ResponseWriter, r *http.Request) {
		handlers.HandlerGetChampionById(databaseRepo, w, r)
	})
	router.Post("/", func(w http.ResponseWriter, r *http.Request) {
		handlers.HandlerInsertChampion(databaseRepo, w, r)
	})

	router.Put("/{id:[0-9]+}", func(w http.ResponseWriter, r *http.Request) {
		handlers.HandlerUpdateChampionByID(databaseRepo, w, r)
	})

	router.Delete("/{id:[0-9]+}", func(w http.ResponseWriter, r *http.Request) {
		handlers.HandlerDeleteChampionByID(databaseRepo, w, r)
	})
}
