package main

import (
	"encoding/json"
	"mkucharsky/wpapi/pkg/models"
	"net/http"
	"net/url"
	"strconv"

	"github.com/go-chi/chi"
)

func (app *application) addURL(w http.ResponseWriter, r *http.Request) {

	err := r.ParseForm()

	if err != nil {

		http.Error(w, "Niepoprawny json", http.StatusBadRequest)
		app.errorLog.Println(err)
		return
	}
	form := r.PostForm

	u := form.Get("url")

	_, err = url.ParseRequestURI(u)

	if err != nil {
		http.Error(w, "Niepoprawny json", http.StatusBadRequest)
		app.errorLog.Println(err)
		return
	}

	interval, err := strconv.Atoi(form.Get("interval"))

	var id int64
	if err != nil {
		http.Error(w, "Niepoprawny json", http.StatusBadRequest)
		app.errorLog.Println(err)
		return
	}

	id, err = app.urls.Insert(u, int64(interval))

	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(map[string]int64{"id": id})

}

func (app *application) deleteURL(w http.ResponseWriter, r *http.Request) {

	id, _ := strconv.Atoi(chi.URLParam(r, "id"))
	result, err := app.urls.Delete(int64(id))

	switch err {
	case nil:
		break
	case models.ErrNoRecord:
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		app.errorLog.Println(err)
		return
	default:
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		app.errorLog.Println(err)
		return
	}

	if result != int64(id) {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		app.errorLog.Println(err)
		return
	}

	json.NewEncoder(w).Encode(map[string]int64{"id": int64(id)})

}

func (app *application) listURLS(w http.ResponseWriter, r *http.Request) {

	results, err := app.urls.Get()

	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		app.errorLog.Println(err)
		return
	}

	json.NewEncoder(w).Encode(results)
}

func (app *application) showHistoryURL(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))

	if err != nil {
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		app.errorLog.Println(err)
	}

	_, err = app.urls.IfExists(int64(id))

	if err != nil && err != models.ErrNoRecord {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		app.errorLog.Println(err)
		return
	}

	if err == models.ErrNoRecord {
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		app.errorLog.Println(err)
		return
	}

	respResults, err := app.responses.Get(int64(id))

	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		app.errorLog.Println(err)
		return
	}

	json.NewEncoder(w).Encode(respResults)
}
