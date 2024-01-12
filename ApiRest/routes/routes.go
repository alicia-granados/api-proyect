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
		router.Get("/", func(w http.ResponseWriter, r *http.Request) {
			handlers.AllChampions(databaseRepo, w, r)
		})
		router.Get("/{id:[0-9]+}", func(w http.ResponseWriter, r *http.Request) {
			handlers.GetChampionId(databaseRepo, w, r)
		})
		router.Post("/", func(w http.ResponseWriter, r *http.Request) {
			handlers.CreateChampion(databaseRepo, w, r)
		})
	})

	return router
}
