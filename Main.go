package main

import (
	"database/sql"
	"log"
	"net/http"

	"./controllers/"
	"./helpers/"
	"github.com/gorilla/mux"
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	db, err := sql.Open("sqlite3", "./test.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", helpers.HttpLogger(controllers.Index()))
	router.HandleFunc("/users", helpers.HttpLogger(controllers.GetUserList(db)))
	router.HandleFunc("/user/{Id}", helpers.HttpLogger(controllers.GetUser(db)))
	router.NotFoundHandler = http.HandlerFunc(helpers.NotFoundLogger)

	log.Fatal(http.ListenAndServe(":9090", router))
}
