package main

import (
	"net/http"
	"mkucharsky/wpapi/pkg/models/mysql"

	//"database/sql"
)

type application struct {
	urls *mysql.URLObjectModel
	responses *mysql.URLResponseModel
}

func main() {

	var dsn string

	db, _ := mysql.OpenDB(dsn)


	app := application{
		urls: &mysql.URLObjectModel{DB: db}, 
		responses: &mysql.URLResponseModel{DB: db},
	}

	server := &http.Server {
		Addr: "localhost:8080",
		Handler: app.routes(),
	}

	server.ListenAndServe()
}