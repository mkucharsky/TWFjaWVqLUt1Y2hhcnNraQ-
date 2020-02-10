package main

import (
	"net/http"
	
	//"database/sql"
)

type application struct {
	num int
}

func main() {

	app := application{
		num: 5,
	}

	server := &http.Server {
		Addr: "localhost:8080",
		Handler: app.routes(),
	}

	server.ListenAndServe()
}