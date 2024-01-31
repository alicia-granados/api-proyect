package routes

import (
	"ApiRest/db"
	"ApiRest/handlers"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func skinRoutes(router chi.Router, databaseRepo db.DatabaseRepo) {
	router.Get("/", func(w http.ResponseWriter, r *http.Request) {
		handlers.HandlerGetSkinList(databaseRepo, w, r)
	})
	router.Get("/{id:[0-9]+}", func(w http.ResponseWriter, r *http.Request) {
		handlers.HandlerGetSkinById(databaseRepo, w, r)
	})
	router.Post("/", func(w http.ResponseWriter, r *http.Request) {
		handlers.HandlerInsertSkin(databaseRepo, w, r)
	})

	router.Put("/{id:[0-9]+}", func(w http.ResponseWriter, r *http.Request) {
		handlers.HandlerUpdateSkinByID(databaseRepo, w, r)
	})

	router.Delete("/{id:[0-9]+}", func(w http.ResponseWriter, r *http.Request) {
		handlers.HandlerDeleteSkinByID(databaseRepo, w, r)
	})
}
