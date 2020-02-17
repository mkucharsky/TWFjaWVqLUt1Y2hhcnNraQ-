package main

import (
	"mkucharsky/wpapi/pkg/models/mysql"
	"mkucharsky/wpapi/pkg/workers"
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
	worker    *workers.Worker
}

func main() {

	var dsn string = "root:XmOBudtu@tcp(localhost:3306)/test?parseTime=true"
	infoLog := log.New(os.Stdout, "INFO\t", log.LstdFlags)
	errorLog := log.New(os.Stderr, "ERROR\t", log.Lshortfile|log.Lshortfile)

	infoLog.Println("Connecting to database...")
	db, err := mysql.OpenDB(dsn)

	if err != nil {
		errorLog.Fatal(err)
	}
	infoLog.Println("Database connected")

	defer db.Close()

	app := application{
		urls:      &mysql.URLObjectModel{DB: db},
		responses: &mysql.URLResponseModel{DB: db},
		infoLog:   errorLog,
		errorLog:  infoLog,
		worker:    workers.NewWorker(),
	}

	app.GetUrlsToWorker()

	infoLog.Println("Starting server...")
	server := &http.Server{
		Addr:    "localhost:8080",
		Handler: app.routes(),
	}

	infoLog.Println("Let's go! Server running")
	server.ListenAndServe()

}
