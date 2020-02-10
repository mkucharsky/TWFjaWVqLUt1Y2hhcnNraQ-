package main

import (
	"net/http"
)

func (app *application) home(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("home"))
}

func (app *application) addURL(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("add"))
}

func (app *application) deleteURL(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("delete"))
}

func (app *application) listURLS(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("list"))
}

func (app *application) showHistoryURL(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("history"))
}