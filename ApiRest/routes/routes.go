package routes

import (
	"ApiRest/db"
	"ApiRest/handlers"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func Routes(databaseRepo *db.RealDBRepo) http.Handler {

	router := chi.NewRouter()

	//mux.Handle("/", http.StripPrefix("/", http.FileServer(http.Dir("./html/"))))

	//protected routes
	router.Route("/api/champion", func(router chi.Router) {
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
	})

	router.Route("/api/skins", func(router chi.Router) {
		router.Get("/", func(w http.ResponseWriter, r *http.Request) {
			handlers.GetSkins(databaseRepo, w, r)
		})
		router.Get("/{id:[0-9]+}", func(w http.ResponseWriter, r *http.Request) {
			handlers.GetSkinsId(databaseRepo, w, r)
		})

	})

	return router
}
