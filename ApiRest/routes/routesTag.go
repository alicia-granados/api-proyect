package routes

import (
	"ApiRest/db"
	"ApiRest/handlers"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func tagRoutes(router chi.Router, databaseRepo db.DatabaseRepo) {
	router.Get("/", func(w http.ResponseWriter, r *http.Request) {
		handlers.HandlerGetTagsList(databaseRepo, w, r)
	})
	router.Get("/{id:[0-9]+}", func(w http.ResponseWriter, r *http.Request) {
		handlers.HandlerGetTagByID(databaseRepo, w, r)
	})
	router.Post("/", func(w http.ResponseWriter, r *http.Request) {
		handlers.HandlerInsertTag(databaseRepo, w, r)
	})

	router.Put("/{id:[0-9]+}", func(w http.ResponseWriter, r *http.Request) {
		handlers.HandlerUpdateTagByID(databaseRepo, w, r)
	})

	router.Delete("/{id:[0-9]+}", func(w http.ResponseWriter, r *http.Request) {
		handlers.HandlerDeleteTagByID(databaseRepo, w, r)
	})
}
