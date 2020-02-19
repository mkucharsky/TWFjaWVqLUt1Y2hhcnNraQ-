package main

import (
	"github.com/go-chi/chi"
	"net/http"
)

func (app *application) routes() http.Handler {

	r := chi.NewRouter()
	r.Post("/api/fetcher/{id}", app.createURL)
	r.Delete("/api/fetcher/{id}", app.deleteURL)
	r.Get("/api/fetcher", app.listURLS)
	r.Get("/api/fetcher/{id}/history", app.showHistoryURL)

	return app.headers(r)
}
