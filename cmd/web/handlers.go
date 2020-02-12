package main

import (
	"encoding/json"
	"github.com/go-chi/chi"
	"net/http"
	"strconv"
)

func (app *application) home(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("home"))
}

func (app *application) addURL(w http.ResponseWriter, r *http.Request) {

	_ = r.ParseForm()

	// if err != nil {
	// 	app.clientError(w, http.StatusBadRequest)
	// 	return
	// }
	form := r.PostForm

	url := form.Get("url")
	interval, err := strconv.Atoi(form.Get("interval"))

	var id int64
	if err != nil {

	} else {
		id, _ = app.urls.Insert(url, int64(interval))
	}

	json.NewEncoder(w).Encode(map[string]int64{"id": id})

}

func (app *application) deleteURL(w http.ResponseWriter, r *http.Request) {

	id, _ := strconv.Atoi(chi.URLParam(r, "id"))
	err := app.urls.Delete(int64(id))

	if err != nil {
		// co się dzieje gdy nie ma takiego id w bazie danych
	} else {
		json.NewEncoder(w).Encode(map[string]int64{"id": int64(id)})
	}
}

func (app *application) listURLS(w http.ResponseWriter, r *http.Request) {

	id, _ := strconv.Atoi(chi.URLParam(r, "id"))
	results, err := app.urls.Get(int64(id))

	if err != nil {
		// co się dzieje gdy error
	} else {
		json.NewEncoder(w).Encode(results)
	}
}

func (app *application) showHistoryURL(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("history"))
}
