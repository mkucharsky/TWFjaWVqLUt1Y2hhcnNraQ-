package main

import (
	"net/http"
)

func (app *application) headers(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {


		next.ServeHTTP(w, r)

	})

}
