package main

import (
	"github.com/go-chi/chi"
)

func (app *application) routes() *chi.Mux {
	
	r := chi.NewRouter()
	r.Get("/", app.home)

	return r
}