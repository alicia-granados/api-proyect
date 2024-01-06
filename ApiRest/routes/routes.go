package routes

import (
	"ApiRest/handlers"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func Routes() http.Handler {

	mux := chi.NewRouter()

	//mux.Handle("/", http.StripPrefix("/", http.FileServer(http.Dir("./html/"))))

	//protected routes
	mux.Route("/api/champion", func(mux chi.Router) {
		mux.Get("/", handlers.AllChampions)
	})

	return mux
}
