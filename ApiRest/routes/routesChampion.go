package routes

import (
	"ApiRest/db"
	"ApiRest/handlers"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func championRoutes(router chi.Router, databaseRepo db.DatabaseRepo) {
	router.Get("/info/", func(w http.ResponseWriter, r *http.Request) {
		handlers.AllInfoChampions(databaseRepo, w, r)
	})
	router.Get("/info/{id:[0-9]+}", func(w http.ResponseWriter, r *http.Request) {
		handlers.GetInfoChampionId(databaseRepo, w, r)
	})

	router.Get("/", func(w http.ResponseWriter, r *http.Request) {
		handlers.GetChampions(databaseRepo, w, r)
	})
	router.Get("/{id:[0-9]+}", func(w http.ResponseWriter, r *http.Request) {
		handlers.GetChampionId(databaseRepo, w, r)
	})
	router.Post("/", func(w http.ResponseWriter, r *http.Request) {
		handlers.CreateChampion(databaseRepo, w, r)
	})

	router.Put("/{id:[0-9]+}", func(w http.ResponseWriter, r *http.Request) {
		handlers.PutChampion(databaseRepo, w, r)
	})

	router.Delete("/{id:[0-9]+}", func(w http.ResponseWriter, r *http.Request) {
		handlers.DeleteChampion(databaseRepo, w, r)
	})
}
