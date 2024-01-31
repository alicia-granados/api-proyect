package routes

import (
	"ApiRest/db"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func Routes(databaseRepo db.DatabaseRepo) http.Handler {
	router := chi.NewRouter()

	router.Route("/api/champion", func(router chi.Router) {
		championRoutes(router, databaseRepo)
	})

	router.Route("/api/skins", func(router chi.Router) {
		skinRoutes(router, databaseRepo)
	})

	router.Route("/api/tags", func(router chi.Router) {
		tagRoutes(router, databaseRepo)
	})

	return router
}
