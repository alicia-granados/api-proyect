package routes

import (
	"ApiRest/db"
	"ApiRest/handlers"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func Routes(databaseRepo db.DatabaseRepo) http.Handler {
	router := chi.NewRouter()

	router.Route("/api/champion", func(router chi.Router) {
		handleChampionRoutes(router, databaseRepo)
	})

	router.Route("/api/skins", func(router chi.Router) {
		handleSkinRoutes(router, databaseRepo)
	})

	router.Route("/api/tags", func(router chi.Router) {
		handleTagRoutes(router, databaseRepo)
	})

	return router
}

func handleChampionRoutes(router chi.Router, databaseRepo db.DatabaseRepo) {
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

func handleSkinRoutes(router chi.Router, databaseRepo db.DatabaseRepo) {
	router.Get("/", func(w http.ResponseWriter, r *http.Request) {
		handlers.GetSkins(databaseRepo, w, r)
	})
	router.Get("/{id:[0-9]+}", func(w http.ResponseWriter, r *http.Request) {
		handlers.GetSkinsId(databaseRepo, w, r)
	})
	router.Post("/", func(w http.ResponseWriter, r *http.Request) {
		handlers.CreateSkins(databaseRepo, w, r)
	})

	router.Put("/{id:[0-9]+}", func(w http.ResponseWriter, r *http.Request) {
		handlers.PutSkin(databaseRepo, w, r)
	})

	router.Delete("/{id:[0-9]+}", func(w http.ResponseWriter, r *http.Request) {
		handlers.DeleteSkin(databaseRepo, w, r)
	})
}

func handleTagRoutes(router chi.Router, databaseRepo db.DatabaseRepo) {
	router.Get("/", func(w http.ResponseWriter, r *http.Request) {
		handlers.GetTags(databaseRepo, w, r)
	})
	router.Get("/{id:[0-9]+}", func(w http.ResponseWriter, r *http.Request) {
		handlers.GetTagId(databaseRepo, w, r)
	})
	router.Post("/", func(w http.ResponseWriter, r *http.Request) {
		handlers.CreateTag(databaseRepo, w, r)
	})

	router.Put("/{id:[0-9]+}", func(w http.ResponseWriter, r *http.Request) {
		handlers.PutTag(databaseRepo, w, r)
	})

	router.Delete("/{id:[0-9]+}", func(w http.ResponseWriter, r *http.Request) {
		handlers.DeleteTag(databaseRepo, w, r)
	})
}
