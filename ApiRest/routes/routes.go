package routes

import (
	"ApiRest/db"
	"ApiRest/handlers"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func Routes(databaseRepo *db.RealDBRepo) http.Handler {

	mux := chi.NewRouter()

	//mux.Handle("/", http.StripPrefix("/", http.FileServer(http.Dir("./html/"))))

	//protected routes
	mux.Route("/api/champion", func(mux chi.Router) {
		mux.Get("/", func(w http.ResponseWriter, r *http.Request) {
			handlers.AllChampions(databaseRepo, w, r)
		})
		mux.Post("/", func(w http.ResponseWriter, r *http.Request) {
			handlers.CreateChampion(databaseRepo, w, r)
		})
	})

	return mux
}
