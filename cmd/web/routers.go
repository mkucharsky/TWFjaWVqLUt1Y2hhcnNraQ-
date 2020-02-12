package main

import (
	"github.com/go-chi/chi"
)

func (app *application) routes() *chi.Mux {

	r := chi.NewRouter()
	r.Get("/", app.home)
	r.Post("/api/fetcher", app.addURL)
	r.Delete("/api/fetcher/{id}", app.deleteURL)
	r.Get("/api/fetcher", app.listURLS)
	r.Get("/api/fetcher/{id}/history", app.showHistoryURL)

	return r
}
