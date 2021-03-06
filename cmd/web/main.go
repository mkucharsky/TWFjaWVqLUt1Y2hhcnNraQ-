package main

import (
	"github.com/joho/godotenv"
	"log"
	"mkucharsky/wpapi/pkg/models/mysql"
	"mkucharsky/wpapi/pkg/workers"
	"net/http"
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

	infoLog := log.New(os.Stdout, "INFO\t", log.LstdFlags)
	errorLog := log.New(os.Stderr, "ERROR\t", log.Lshortfile|log.Lshortfile)

	if err := godotenv.Load("./../.env"); err != nil {
		errorLog.Fatal(err)
	}

	dbPort := os.Getenv("MYSQL_PORT")
	dbUser := os.Getenv("MYSQL_USER")
	dbPass := os.Getenv("MYSQL_PASS")
	dbName := os.Getenv("MYSQL_DB")
	dbContainerName := os.Getenv("MYSQL_CONTAINER_NAME")

	infoLog.Println("Connecting to database...")

	var dsn string = dbUser + ":" + dbPass + "@tcp("+ dbContainerName +":" + dbPort + ")/" + dbName + "?parseTime=true"

	db, err := mysql.OpenDB(dsn)

	if err != nil {
		errorLog.Println(dsn)
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
		Addr:    ":8080",
		Handler: app.routes(),
	}

	infoLog.Println("Jupi! Server running")
	server.ListenAndServe()

}
