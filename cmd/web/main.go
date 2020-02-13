package main

import (
	"mkucharsky/wpapi/pkg/models/mysql"
	"net/http"
	//"database/sql"
	"log"
	"os"
)

type application struct {
	urls      *mysql.URLObjectModel
	responses *mysql.URLResponseModel
	infoLog   *log.Logger
	errorLog  *log.Logger
}

func main() {

	var dsn string

	db, _ := mysql.OpenDB(dsn)

	app := application{
		urls:      &mysql.URLObjectModel{DB: db},
		responses: &mysql.URLResponseModel{DB: db},
		infoLog:   log.New(os.Stdout, "INFO\t", log.LstdFlags),
		errorLog:  log.New(os.Stderr, "INFO\t", log.LstdFlags|log.Lshortfile),
	}

	server := &http.Server{
		Addr:    "localhost:8080",
		Handler: app.routes(),
	}

	server.ListenAndServe()
}
